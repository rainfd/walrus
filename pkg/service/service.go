package service

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	deptypes "github.com/seal-io/walrus/pkg/deployer/types"
	"github.com/seal-io/walrus/utils/errorx"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/strs"
)

const annotationSubjectIDName = "walrus.seal.io/subject-id"

// Options for deploy or destroy.
type Options struct {
	TlsCertified bool
	Tags         []string
}

func Create(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (*model.ServiceOutput, error) {
	err := mc.WithTx(ctx, func(tx *model.Tx) (err error) {
		// TODO(thxCode): generated by entc.
		status.ServiceStatusDeployed.Unknown(entity, "")
		entity.Status.SetSummary(status.WalkService(&entity.Status))

		entity, err = tx.Services().Create().
			Set(entity).
			SaveE(ctx, dao.ServiceDependenciesEdgeSave)

		return err
	})
	if err != nil {
		return nil, err
	}

	ready, err := CheckDependencyStatus(ctx, mc, entity)
	if err != nil {
		return nil, err
	}

	// Service dependency ready can be applied promptly.
	if ready {
		// Deploy service.
		err = Apply(ctx, mc, dp, entity, Options{
			TlsCertified: opts.TlsCertified,
			Tags:         opts.Tags,
		})
		if err != nil {
			return nil, err
		}
	}

	return model.ExposeService(entity), nil
}

func UpdateStatus(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Service,
) error {
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	err := mc.Services().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
	if err != nil && !model.IsNotFound(err) {
		return err
	}

	return nil
}

func Apply(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (err error) {
	logger := log.WithName("service")

	if !status.ServiceStatusDeployed.IsUnknown(entity) {
		return errorx.Errorf("service status is not deploying, service: %s", entity.ID)
	}

	applyOpts := deptypes.ApplyOptions{
		SkipTLSVerify: !opts.TlsCertified,
		Tags:          opts.Tags,
	}

	err = dp.Apply(ctx, entity, applyOpts)
	if err != nil {
		err = fmt.Errorf("failed to apply service: %w", err)
		logger.Error(err)

		// Update a failure status.
		status.ServiceStatusDeployed.False(entity, err.Error())

		err = UpdateStatus(ctx, mc, entity)
		if err != nil {
			logger.Errorf("error updating status of service %s: %v", entity.ID, err)
		}
	}

	return nil
}

func Destroy(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (err error) {
	logger := log.WithName("service")

	updateFailedStatus := func(err error) {
		status.ServiceStatusDeleted.False(entity, err.Error())

		err = UpdateStatus(ctx, mc, entity)
		if err != nil {
			logger.Errorf("error updating status of service %s: %v", entity.ID, err)
		}
	}

	// Check dependants.
	dependants, err := dao.GetServiceDependantNames(ctx, mc, entity)
	if err != nil {
		updateFailedStatus(err)
		return err
	}

	if len(dependants) > 0 {
		msg := fmt.Sprintf("Waiting for dependants to be deleted: %s", strs.Join(", ", dependants...))
		if !status.ServiceStatusProgressing.IsUnknown(entity) ||
			status.ServiceStatusDeleted.GetMessage(entity) != msg {
			// Mark status to deleting with dependency message.
			status.ServiceStatusDeleted.Reset(entity, msg)
			status.ServiceStatusProgressing.Unknown(entity, "")

			if err = UpdateStatus(ctx, mc, entity); err != nil {
				return fmt.Errorf("failed to update service status: %w", err)
			}
		}

		return nil
	} else {
		// Mark status to deleting.
		status.ServiceStatusDeleted.Reset(entity, "")
		status.ServiceStatusProgressing.True(entity, "Resolved dependencies")

		if err = UpdateStatus(ctx, mc, entity); err != nil {
			return fmt.Errorf("failed to update service status: %w", err)
		}
	}

	destroyOpts := deptypes.DestroyOptions{
		SkipTLSVerify: !opts.TlsCertified,
	}

	err = dp.Destroy(ctx, entity, destroyOpts)
	if err != nil {
		log.Errorf("fail to destroy service: %w", err)

		updateFailedStatus(err)
	}

	return nil
}

func GetSubjectID(entity *model.Service) (object.ID, error) {
	if entity == nil {
		return "", fmt.Errorf("service is nil")
	}

	subjectIDStr := entity.Annotations[annotationSubjectIDName]

	return object.ID(subjectIDStr), nil
}

func SetSubjectID(ctx context.Context, services ...*model.Service) error {
	sj, err := session.GetSubject(ctx)
	if err != nil {
		return err
	}

	for i := range services {
		if services[i].Annotations == nil {
			services[i].Annotations = make(map[string]string)
		}
		services[i].Annotations[annotationSubjectIDName] = string(sj.ID)
	}

	return nil
}

// SetServiceStatusScheduled sets the status of the service to scheduled.
func SetServiceStatusScheduled(ctx context.Context, mc model.ClientSet, entity *model.Service) error {
	if entity == nil {
		return fmt.Errorf("service is nil")
	}

	dependencyNames := dao.ServiceRelationshipGetDependencyNames(entity)

	msg := ""
	if len(dependencyNames) > 0 {
		msg = fmt.Sprintf("Waiting for dependent services to be ready: %s", strs.Join(", ", dependencyNames...))
	}

	status.ServiceStatusProgressing.Reset(entity, msg)
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	return mc.Services().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
}

// CreateScheduledServices creates scheduled services.
func CreateScheduledServices(ctx context.Context, mc model.ClientSet, entities model.Services) (model.Services, error) {
	results := make(model.Services, 0, len(entities))

	sortedServices, err := TopologicalSortServices(entities)
	if err != nil {
		return nil, err
	}

	for i := range sortedServices {
		entity := sortedServices[i]

		err = mc.WithTx(ctx, func(tx *model.Tx) error {
			// TODO(thxCode): generated by entc.
			status.ServiceStatusDeployed.Unknown(entity, "")
			entity.Status.SetSummary(status.WalkService(&entity.Status))

			entity, err = tx.Services().Create().
				Set(entity).
				SaveE(ctx, dao.ServiceDependenciesEdgeSave)
			if err != nil {
				return err
			}

			return SetServiceStatusScheduled(ctx, tx, entity)
		})
		if err != nil {
			return nil, err
		}

		results = append(results, entity)
	}

	return results, nil
}

// IsStatusReady returns true if the service is ready.
func IsStatusReady(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "Preparing", "NotReady", "Ready":
		return true
	}

	return false
}

// IsStatusFalse returns true if the service is in error status.
func IsStatusFalse(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "DeployFailed", "DeleteFailed":
		return true
	case "Progressing":
		return entity.Status.Error
	}

	return false
}

// IsStatusDeleted returns true if the service is deleted.
func IsStatusDeleted(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "Deleted", "Deleting":
		return true
	}

	return false
}

const (
	summaryStatusDeploying   = "Deploying"
	summaryStatusProgressing = "Progressing"
)

// CheckDependencyStatus check service dependencies status is ready to apply.
func CheckDependencyStatus(ctx context.Context, mc model.ClientSet, entity *model.Service) (bool, error) {
	// Check dependants.
	dependencies, err := mc.ServiceRelationships().Query().
		Where(
			servicerelationship.ServiceID(entity.ID),
			servicerelationship.DependencyIDNEQ(entity.ID),
		).
		QueryDependency().
		Select(service.FieldID).
		Where(
			service.Or(
				func(s *sql.Selector) {
					s.Where(sqljson.ValueEQ(
						service.FieldStatus,
						summaryStatusDeploying,
						sqljson.Path("summaryStatus"),
					))
				},
				service.And(
					func(s *sql.Selector) {
						s.Where(sqljson.ValueEQ(
							service.FieldStatus,
							summaryStatusProgressing,
							sqljson.Path("summaryStatus"),
						))
					},
					func(s *sql.Selector) {
						s.Where(sqljson.ValueEQ(
							service.FieldStatus,
							true,
							sqljson.Path("transitioning"),
						))
					},
				),
			),
		).
		All(ctx)
	if err != nil {
		return false, err
	}

	if len(dependencies) > 0 {
		// If dependency services is in deploying status.
		err = SetServiceStatusScheduled(ctx, mc, entity)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return true, nil
}
