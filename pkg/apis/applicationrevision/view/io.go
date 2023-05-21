package view

import (
	"context"
	"errors"

	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/applicationinstance"
	"github.com/seal-io/seal/pkg/dao/model/applicationrevision"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/oid"
	"github.com/seal-io/seal/pkg/dao/types/property"
	"github.com/seal-io/seal/pkg/dao/types/status"
	"github.com/seal-io/seal/pkg/deployer/terraform"
	"github.com/seal-io/seal/pkg/topic/datamessage"
	"github.com/seal-io/seal/utils/json"
)

// Basic APIs.

type GetRequest struct {
	*model.ApplicationRevisionQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *GetRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	return nil
}

type GetResponse = *model.ApplicationRevisionOutput

type StreamRequest struct {
	ID        oid.ID `uri:"id"`
	ProjectID oid.ID `query:"projectID"`
}

func (r *StreamRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	modelClient := input.(model.ClientSet)

	exist, err := modelClient.ApplicationRevisions().Query().
		Where(applicationrevision.ID(r.ID)).
		Exist(ctx)
	if err != nil || !exist {
		return runtime.Errorw(err, "invalid id: not found")
	}

	return nil
}

type StreamResponse struct {
	Type       datamessage.EventType              `json:"type"`
	IDs        []oid.ID                           `json:"ids,omitempty"`
	Collection []*model.ApplicationRevisionOutput `json:"collection,omitempty"`
}

// Batch APIs.

type CollectionDeleteRequest []*model.ApplicationRevisionQueryInput

func (r CollectionDeleteRequest) ValidateWith(ctx context.Context, input any) error {
	if len(r) == 0 {
		return errors.New("invalid ids: blank")
	}

	var (
		ids         = make([]oid.ID, 0, len(r))
		modelClient = input.(model.ClientSet)
	)

	for _, i := range r {
		if !i.ID.Valid(0) {
			return errors.New("invalid ids: blank")
		}

		ids = append(ids, i.ID)
	}

	revisions, err := modelClient.ApplicationRevisions().Query().
		Select(applicationrevision.FieldID, applicationrevision.FieldInstanceID).
		Where(applicationrevision.IDIn(ids...)).
		All(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get application revisions")
	}

	if len(revisions) != len(r) {
		return errors.New("invalid ids: some revisions are not found")
	}

	instanceID := revisions[0].InstanceID
	for _, revision := range revisions {
		if revision.InstanceID != instanceID {
			return errors.New("invalid ids: revision ids are not in the same instance")
		}
	}

	latestRevision, err := modelClient.ApplicationRevisions().Query().
		Select(applicationrevision.FieldID).
		Where(applicationrevision.InstanceID(instanceID)).
		Order(model.Desc(applicationrevision.FieldCreateTime)).
		First(ctx)
	if err != nil {
		return runtime.Errorw(err, "failed to get latest revision")
	}

	for _, revision := range revisions {
		// Prevent deleting the latest revision.
		if revision.ID == latestRevision.ID {
			return errors.New("invalid ids: can not delete latest revision")
		}
	}

	return nil
}

type CollectionGetRequest struct {
	runtime.RequestPagination                               `query:",inline"`
	runtime.RequestExtracting                               `query:",inline"`
	runtime.RequestSorting[applicationrevision.OrderOption] `query:",inline"`

	ProjectID  oid.ID `query:"projectID"`
	InstanceID oid.ID `query:"instanceID,omitempty"`
}

func (r *CollectionGetRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if r.InstanceID != "" {
		if !r.InstanceID.IsNaive() {
			return errors.New("invalid instance id")
		}

		modelClient := input.(model.ClientSet)

		_, err := modelClient.ApplicationInstances().Query().
			Where(applicationinstance.ID(r.InstanceID)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get application instance")
		}
	}

	return nil
}

type CollectionGetResponse = []*model.ApplicationRevisionOutput

type CollectionStreamRequest struct {
	runtime.RequestExtracting `query:",inline"`

	ProjectID  oid.ID `query:"projectID"`
	InstanceID oid.ID `query:"instanceID,omitempty"`
}

func (r *CollectionStreamRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if r.InstanceID != "" {
		if !r.InstanceID.IsNaive() {
			return errors.New("invalid instance id")
		}

		modelClient := input.(model.ClientSet)

		_, err := modelClient.ApplicationInstances().Query().
			Where(applicationinstance.ID(r.InstanceID)).
			OnlyID(ctx)
		if err != nil {
			return runtime.Errorw(err, "failed to get application instance")
		}
	}

	return nil
}

// Extensional APIs.

type GetTerraformStatesRequest = GetRequest

type GetTerraformStatesResponse = json.RawMessage

type UpdateTerraformStatesRequest struct {
	GetRequest `query:",inline" uri:",inline"`

	json.RawMessage `uri:"-" json:",inline"`
}

type StreamLogRequest struct {
	ID        oid.ID `uri:"id"`
	ProjectID oid.ID `query:"projectID"`
	JobType   string `query:"jobType,omitempty"`
}

func (r *StreamLogRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	if r.JobType == "" {
		r.JobType = terraform.JobTypeApply
	}

	if r.JobType != terraform.JobTypeApply && r.JobType != terraform.JobTypeDestroy {
		return errors.New("invalid job type")
	}

	return nil
}

type RollbackInstanceRequest struct {
	_ struct{} `route:"POST=/rollback-instances"`

	*model.ApplicationRevisionQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *RollbackInstanceRequest) ValidateWith(ctx context.Context, input any) error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	// Check latest revision if running.
	modelClient := input.(model.ClientSet)

	entity, err := modelClient.ApplicationRevisions().Get(ctx, r.ID)
	if err != nil {
		return runtime.Errorw(err, "failed to get application revision")
	}

	latestRevision, err := modelClient.ApplicationRevisions().Query().
		Select(applicationrevision.FieldStatus).
		Where(applicationrevision.InstanceID(entity.InstanceID)).
		Order(model.Desc(applicationrevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return runtime.Errorw(err, "failed to get latest revision")
	}

	if latestRevision.Status == status.ApplicationRevisionStatusRunning {
		return errors.New("latest revision is running")
	}

	return nil
}

type RollbackApplicationRequest struct {
	_ struct{} `route:"POST=/rollback-applications"`

	*model.ApplicationRevisionQueryInput `uri:",inline"`

	ProjectID oid.ID `query:"projectID"`
}

func (r *RollbackApplicationRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	return nil
}

type DiffLatestRequest struct {
	ID        oid.ID `uri:"id"`
	ProjectID oid.ID `query:"projectID"`
}

func (r *DiffLatestRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	return nil
}

type RevisionDiffPreviousRequest struct {
	ID        oid.ID `uri:"id"`
	ProjectID oid.ID `query:"projectID"`
}

func (r *RevisionDiffPreviousRequest) Validate() error {
	if !r.ProjectID.Valid(0) {
		return errors.New("invalid project id: blank")
	}

	if !r.ID.Valid(0) {
		return errors.New("invalid id: blank")
	}

	return nil
}

type RevisionDiff struct {
	InputVariables property.Values           `json:"inputVariables"`
	Variables      property.Schemas          `json:"variables"`
	Modules        []types.ApplicationModule `json:"modules"`
}

type RevisionDiffResponse struct {
	Old RevisionDiff `json:"old"`
	New RevisionDiff `json:"new"`
}
