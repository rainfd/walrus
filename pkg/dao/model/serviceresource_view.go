// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ServiceResourceCreateInput holds the creation input of the ServiceResource entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to create ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Shape of the resource, it can be class or instance shape.
	Shape string `path:"-" query:"-" json:"shape"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Name of the generated resource, it is the real identifier of the resource, which provides by deployer.
	Name string `path:"-" query:"-" json:"name"`
	// Type of the generated resource, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `path:"-" query:"-" json:"type"`
	// Mode that manages the generated resource, it is the management way of the deployer to the resource, which provides by deployer.
	Mode string `path:"-" query:"-" json:"mode"`
	// Status of the resource.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components specifies full inserting the new ServiceResource entities of the ServiceResource entity.
	Components []*ServiceResourceCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances specifies full inserting the new ServiceResource entities of the ServiceResource entity.
	Instances []*ServiceResourceCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies specifies full inserting the new ServiceResourceRelationship entities of the ServiceResource entity.
	Dependencies []*ServiceResourceRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the ServiceResource entity for creating,
// after validating.
func (srci *ServiceResourceCreateInput) Model() *ServiceResource {
	if srci == nil {
		return nil
	}

	_sr := &ServiceResource{
		Shape:        srci.Shape,
		DeployerType: srci.DeployerType,
		Name:         srci.Name,
		Type:         srci.Type,
		Mode:         srci.Mode,
		Status:       srci.Status,
	}

	if srci.Project != nil {
		_sr.ProjectID = srci.Project.ID
	}
	if srci.Environment != nil {
		_sr.EnvironmentID = srci.Environment.ID
	}
	if srci.Service != nil {
		_sr.ServiceID = srci.Service.ID
	}

	if srci.Components != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Components = make([]*ServiceResource, 0, len(srci.Components))
	}
	for j := range srci.Components {
		if srci.Components[j] == nil {
			continue
		}
		_sr.Edges.Components = append(_sr.Edges.Components,
			srci.Components[j].Model())
	}
	if srci.Instances != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Instances = make([]*ServiceResource, 0, len(srci.Instances))
	}
	for j := range srci.Instances {
		if srci.Instances[j] == nil {
			continue
		}
		_sr.Edges.Instances = append(_sr.Edges.Instances,
			srci.Instances[j].Model())
	}
	if srci.Dependencies != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Dependencies = make([]*ServiceResourceRelationship, 0, len(srci.Dependencies))
	}
	for j := range srci.Dependencies {
		if srci.Dependencies[j] == nil {
			continue
		}
		_sr.Edges.Dependencies = append(_sr.Edges.Dependencies,
			srci.Dependencies[j].Model())
	}
	return _sr
}

// Validate checks the ServiceResourceCreateInput entity.
func (srci *ServiceResourceCreateInput) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceCreateInput entity with the given context and client set.
func (srci *ServiceResourceCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srci.Project != nil {
		if err := srci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Environment route.
	if srci.Environment != nil {
		if err := srci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Service route.
	if srci.Service != nil {
		if err := srci.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	for i := range srci.Components {
		if srci.Components[i] == nil {
			continue
		}

		if err := srci.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Components[i] = nil
			}
		}
	}

	for i := range srci.Instances {
		if srci.Instances[i] == nil {
			continue
		}

		if err := srci.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Instances[i] = nil
			}
		}
	}

	for i := range srci.Dependencies {
		if srci.Dependencies[i] == nil {
			continue
		}

		if err := srci.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ServiceResourceCreateInputs holds the creation input item of the ServiceResource entities.
type ServiceResourceCreateInputsItem struct {
	// Shape of the resource, it can be class or instance shape.
	Shape string `path:"-" query:"-" json:"shape"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Name of the generated resource, it is the real identifier of the resource, which provides by deployer.
	Name string `path:"-" query:"-" json:"name"`
	// Type of the generated resource, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `path:"-" query:"-" json:"type"`
	// Mode that manages the generated resource, it is the management way of the deployer to the resource, which provides by deployer.
	Mode string `path:"-" query:"-" json:"mode"`
	// Status of the resource.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components specifies full inserting the new ServiceResource entities.
	Components []*ServiceResourceCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances specifies full inserting the new ServiceResource entities.
	Instances []*ServiceResourceCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies specifies full inserting the new ServiceResourceRelationship entities.
	Dependencies []*ServiceResourceRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// ValidateWith checks the ServiceResourceCreateInputsItem entity with the given context and client set.
func (srci *ServiceResourceCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range srci.Components {
		if srci.Components[i] == nil {
			continue
		}

		if err := srci.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Components[i] = nil
			}
		}
	}

	for i := range srci.Instances {
		if srci.Instances[i] == nil {
			continue
		}

		if err := srci.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Instances[i] = nil
			}
		}
	}

	for i := range srci.Dependencies {
		if srci.Dependencies[i] == nil {
			continue
		}

		if err := srci.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ServiceResourceCreateInputs holds the creation input of the ServiceResource entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to create ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResource entities for creating,
// after validating.
func (srci *ServiceResourceCreateInputs) Model() []*ServiceResource {
	if srci == nil || len(srci.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceResource, len(srci.Items))

	for i := range srci.Items {
		_sr := &ServiceResource{
			Shape:        srci.Items[i].Shape,
			DeployerType: srci.Items[i].DeployerType,
			Name:         srci.Items[i].Name,
			Type:         srci.Items[i].Type,
			Mode:         srci.Items[i].Mode,
			Status:       srci.Items[i].Status,
		}

		if srci.Project != nil {
			_sr.ProjectID = srci.Project.ID
		}
		if srci.Environment != nil {
			_sr.EnvironmentID = srci.Environment.ID
		}
		if srci.Service != nil {
			_sr.ServiceID = srci.Service.ID
		}

		if srci.Items[i].Components != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Components = make([]*ServiceResource, 0, len(srci.Items[i].Components))
		}
		for j := range srci.Items[i].Components {
			if srci.Items[i].Components[j] == nil {
				continue
			}
			_sr.Edges.Components = append(_sr.Edges.Components,
				srci.Items[i].Components[j].Model())
		}
		if srci.Items[i].Instances != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Instances = make([]*ServiceResource, 0, len(srci.Items[i].Instances))
		}
		for j := range srci.Items[i].Instances {
			if srci.Items[i].Instances[j] == nil {
				continue
			}
			_sr.Edges.Instances = append(_sr.Edges.Instances,
				srci.Items[i].Instances[j].Model())
		}
		if srci.Items[i].Dependencies != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Dependencies = make([]*ServiceResourceRelationship, 0, len(srci.Items[i].Dependencies))
		}
		for j := range srci.Items[i].Dependencies {
			if srci.Items[i].Dependencies[j] == nil {
				continue
			}
			_sr.Edges.Dependencies = append(_sr.Edges.Dependencies,
				srci.Items[i].Dependencies[j].Model())
		}

		_srs[i] = _sr
	}

	return _srs
}

// Validate checks the ServiceResourceCreateInputs entity .
func (srci *ServiceResourceCreateInputs) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceCreateInputs entity with the given context and client set.
func (srci *ServiceResourceCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if len(srci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srci.Project != nil {
		if err := srci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if srci.Environment != nil {
		if err := srci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Environment = nil
			}
		}
	}
	// Validate when creating under the Service route.
	if srci.Service != nil {
		if err := srci.Service.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Service = nil
			}
		}
	}

	for i := range srci.Items {
		if srci.Items[i] == nil {
			continue
		}

		if err := srci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceResourceDeleteInput holds the deletion input of the ServiceResource entity,
// please tags with `path:",inline"` if embedding.
type ServiceResourceDeleteInput struct {
	ServiceResourceQueryInput `path:",inline"`
}

// ServiceResourceDeleteInputs holds the deletion input item of the ServiceResource entities.
type ServiceResourceDeleteInputsItem struct {
	// ID of the ServiceResource entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// ServiceResourceDeleteInputs holds the deletion input of the ServiceResource entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to delete ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to delete ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResource entities for deleting,
// after validating.
func (srdi *ServiceResourceDeleteInputs) Model() []*ServiceResource {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceResource, len(srdi.Items))
	for i := range srdi.Items {
		_srs[i] = &ServiceResource{
			ID: srdi.Items[i].ID,
		}
	}
	return _srs
}

// IDs returns the ID list of the ServiceResource entities for deleting,
// after validating.
func (srdi *ServiceResourceDeleteInputs) IDs() []object.ID {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srdi.Items))
	for i := range srdi.Items {
		ids[i] = srdi.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceResourceDeleteInputs entity.
func (srdi *ServiceResourceDeleteInputs) Validate() error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	return srdi.ValidateWith(srdi.inputConfig.Context, srdi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceDeleteInputs entity with the given context and client set.
func (srdi *ServiceResourceDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	if len(srdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResources().Query()

	// Validate when deleting under the Project route.
	if srdi.Project != nil {
		if err := srdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				serviceresource.ProjectID(srdi.Project.ID))
		}
	}

	// Validate when deleting under the Environment route.
	if srdi.Environment != nil {
		if err := srdi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.EnvironmentID(srdi.Environment.ID))
		}
	}

	// Validate when deleting under the Service route.
	if srdi.Service != nil {
		if err := srdi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.ServiceID(srdi.Service.ID))
		}
	}

	ids := make([]object.ID, 0, len(srdi.Items))

	for i := range srdi.Items {
		if srdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if srdi.Items[i].ID != "" {
			ids = append(ids, srdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(serviceresource.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ServiceResourceQueryInput holds the query input of the ServiceResource entity,
// please tags with `path:",inline"` if embedding.
type ServiceResourceQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`
	// Environment indicates to query ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"environment"`
	// Service indicates to query ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"service"`

	// Refer holds the route path reference of the ServiceResource entity.
	Refer *object.Refer `path:"serviceresource,default=" query:"-" json:"-"`
	// ID of the ServiceResource entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the ServiceResource entity for querying,
// after validating.
func (srqi *ServiceResourceQueryInput) Model() *ServiceResource {
	if srqi == nil {
		return nil
	}

	return &ServiceResource{
		ID: srqi.ID,
	}
}

// Validate checks the ServiceResourceQueryInput entity.
func (srqi *ServiceResourceQueryInput) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceQueryInput entity with the given context and client set.
func (srqi *ServiceResourceQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if srqi.Refer != nil && *srqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", serviceresource.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResources().Query()

	// Validate when querying under the Project route.
	if srqi.Project != nil {
		if err := srqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				serviceresource.ProjectID(srqi.Project.ID))
		}
	}

	// Validate when querying under the Environment route.
	if srqi.Environment != nil {
		if err := srqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.EnvironmentID(srqi.Environment.ID))
		}
	}

	// Validate when querying under the Service route.
	if srqi.Service != nil {
		if err := srqi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.ServiceID(srqi.Service.ID))
		}
	}

	if srqi.Refer != nil {
		if srqi.Refer.IsID() {
			q.Where(
				serviceresource.ID(srqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of serviceresource")
		}
	} else if srqi.ID != "" {
		q.Where(
			serviceresource.ID(srqi.ID))
	} else {
		return errors.New("invalid identify of serviceresource")
	}

	q.Select(
		serviceresource.FieldID,
	)

	var e *ServiceResource
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*ServiceResource)
		}
	}

	srqi.ID = e.ID
	return nil
}

// ServiceResourceQueryInputs holds the query input of the ServiceResource entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ServiceResourceQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to query ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to query ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the ServiceResourceQueryInputs entity.
func (srqi *ServiceResourceQueryInputs) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceQueryInputs entity with the given context and client set.
func (srqi *ServiceResourceQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if srqi.Project != nil {
		if err := srqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Environment route.
	if srqi.Environment != nil {
		if err := srqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Service route.
	if srqi.Service != nil {
		if err := srqi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceResourceUpdateInput holds the modification input of the ServiceResource entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceUpdateInput struct {
	ServiceResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Status of the resource.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components indicates replacing the stale ServiceResource entities.
	Components []*ServiceResourceCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances indicates replacing the stale ServiceResource entities.
	Instances []*ServiceResourceCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies indicates replacing the stale ServiceResourceRelationship entities.
	Dependencies []*ServiceResourceRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the ServiceResource entity for modifying,
// after validating.
func (srui *ServiceResourceUpdateInput) Model() *ServiceResource {
	if srui == nil {
		return nil
	}

	_sr := &ServiceResource{
		ID:     srui.ID,
		Status: srui.Status,
	}

	if srui.Components != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Components = make([]*ServiceResource, 0, len(srui.Components))
	}
	for j := range srui.Components {
		if srui.Components[j] == nil {
			continue
		}
		_sr.Edges.Components = append(_sr.Edges.Components,
			srui.Components[j].Model())
	}
	if srui.Instances != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Instances = make([]*ServiceResource, 0, len(srui.Instances))
	}
	for j := range srui.Instances {
		if srui.Instances[j] == nil {
			continue
		}
		_sr.Edges.Instances = append(_sr.Edges.Instances,
			srui.Instances[j].Model())
	}
	if srui.Dependencies != nil {
		// Empty slice is used for clearing the edge.
		_sr.Edges.Dependencies = make([]*ServiceResourceRelationship, 0, len(srui.Dependencies))
	}
	for j := range srui.Dependencies {
		if srui.Dependencies[j] == nil {
			continue
		}
		_sr.Edges.Dependencies = append(_sr.Edges.Dependencies,
			srui.Dependencies[j].Model())
	}
	return _sr
}

// Validate checks the ServiceResourceUpdateInput entity.
func (srui *ServiceResourceUpdateInput) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceUpdateInput entity with the given context and client set.
func (srui *ServiceResourceUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := srui.ServiceResourceQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	for i := range srui.Components {
		if srui.Components[i] == nil {
			continue
		}

		if err := srui.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Components[i] = nil
			}
		}
	}

	for i := range srui.Instances {
		if srui.Instances[i] == nil {
			continue
		}

		if err := srui.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Instances[i] = nil
			}
		}
	}

	for i := range srui.Dependencies {
		if srui.Dependencies[i] == nil {
			continue
		}

		if err := srui.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ServiceResourceUpdateInputs holds the modification input item of the ServiceResource entities.
type ServiceResourceUpdateInputsItem struct {
	// ID of the ServiceResource entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Status of the resource.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components indicates replacing the stale ServiceResource entities.
	Components []*ServiceResourceCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances indicates replacing the stale ServiceResource entities.
	Instances []*ServiceResourceCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies indicates replacing the stale ServiceResourceRelationship entities.
	Dependencies []*ServiceResourceRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// ValidateWith checks the ServiceResourceUpdateInputsItem entity with the given context and client set.
func (srui *ServiceResourceUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range srui.Components {
		if srui.Components[i] == nil {
			continue
		}

		if err := srui.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Components[i] = nil
			}
		}
	}

	for i := range srui.Instances {
		if srui.Instances[i] == nil {
			continue
		}

		if err := srui.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Instances[i] = nil
			}
		}
	}

	for i := range srui.Dependencies {
		if srui.Dependencies[i] == nil {
			continue
		}

		if err := srui.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ServiceResourceUpdateInputs holds the modification input of the ServiceResource entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceResourceUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update ServiceResource entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to update ServiceResource entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to update ServiceResource entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceResourceUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceResource entities for modifying,
// after validating.
func (srui *ServiceResourceUpdateInputs) Model() []*ServiceResource {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceResource, len(srui.Items))

	for i := range srui.Items {
		_sr := &ServiceResource{
			ID:     srui.Items[i].ID,
			Status: srui.Items[i].Status,
		}

		if srui.Items[i].Components != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Components = make([]*ServiceResource, 0, len(srui.Items[i].Components))
		}
		for j := range srui.Items[i].Components {
			if srui.Items[i].Components[j] == nil {
				continue
			}
			_sr.Edges.Components = append(_sr.Edges.Components,
				srui.Items[i].Components[j].Model())
		}
		if srui.Items[i].Instances != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Instances = make([]*ServiceResource, 0, len(srui.Items[i].Instances))
		}
		for j := range srui.Items[i].Instances {
			if srui.Items[i].Instances[j] == nil {
				continue
			}
			_sr.Edges.Instances = append(_sr.Edges.Instances,
				srui.Items[i].Instances[j].Model())
		}
		if srui.Items[i].Dependencies != nil {
			// Empty slice is used for clearing the edge.
			_sr.Edges.Dependencies = make([]*ServiceResourceRelationship, 0, len(srui.Items[i].Dependencies))
		}
		for j := range srui.Items[i].Dependencies {
			if srui.Items[i].Dependencies[j] == nil {
				continue
			}
			_sr.Edges.Dependencies = append(_sr.Edges.Dependencies,
				srui.Items[i].Dependencies[j].Model())
		}

		_srs[i] = _sr
	}

	return _srs
}

// IDs returns the ID list of the ServiceResource entities for modifying,
// after validating.
func (srui *ServiceResourceUpdateInputs) IDs() []object.ID {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srui.Items))
	for i := range srui.Items {
		ids[i] = srui.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceResourceUpdateInputs entity.
func (srui *ServiceResourceUpdateInputs) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceResourceUpdateInputs entity with the given context and client set.
func (srui *ServiceResourceUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if len(srui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceResources().Query()

	// Validate when updating under the Project route.
	if srui.Project != nil {
		if err := srui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				serviceresource.ProjectID(srui.Project.ID))
		}
	}

	// Validate when updating under the Environment route.
	if srui.Environment != nil {
		if err := srui.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.EnvironmentID(srui.Environment.ID))
		}
	}

	// Validate when updating under the Service route.
	if srui.Service != nil {
		if err := srui.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				serviceresource.ServiceID(srui.Service.ID))
		}
	}

	ids := make([]object.ID, 0, len(srui.Items))

	for i := range srui.Items {
		if srui.Items[i] == nil {
			return errors.New("nil item")
		}

		if srui.Items[i].ID != "" {
			ids = append(ids, srui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(serviceresource.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range srui.Items {
		if err := srui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceResourceOutput holds the output of the ServiceResource entity.
type ServiceResourceOutput struct {
	ID           object.ID                           `json:"id,omitempty"`
	CreateTime   *time.Time                          `json:"createTime,omitempty"`
	UpdateTime   *time.Time                          `json:"updateTime,omitempty"`
	Mode         string                              `json:"mode,omitempty"`
	Type         string                              `json:"type,omitempty"`
	Name         string                              `json:"name,omitempty"`
	DeployerType string                              `json:"deployerType,omitempty"`
	Shape        string                              `json:"shape,omitempty"`
	Status       types.ServiceResourceStatus         `json:"status,omitempty"`
	Keys         *types.ServiceResourceOperationKeys `json:"keys,omitempty"`

	Project      *ProjectOutput                       `json:"project,omitempty"`
	Environment  *EnvironmentOutput                   `json:"environment,omitempty"`
	Service      *ServiceOutput                       `json:"service,omitempty"`
	Connector    *ConnectorOutput                     `json:"connector,omitempty"`
	Composition  *ServiceResourceOutput               `json:"composition,omitempty"`
	Components   []*ServiceResourceOutput             `json:"components,omitempty"`
	Class        *ServiceResourceOutput               `json:"class,omitempty"`
	Instances    []*ServiceResourceOutput             `json:"instances,omitempty"`
	Dependencies []*ServiceResourceRelationshipOutput `json:"dependencies,omitempty"`
}

// View returns the output of ServiceResource entity.
func (_sr *ServiceResource) View() *ServiceResourceOutput {
	return ExposeServiceResource(_sr)
}

// View returns the output of ServiceResource entities.
func (_srs ServiceResources) View() []*ServiceResourceOutput {
	return ExposeServiceResources(_srs)
}

// ExposeServiceResource converts the ServiceResource to ServiceResourceOutput.
func ExposeServiceResource(_sr *ServiceResource) *ServiceResourceOutput {
	if _sr == nil {
		return nil
	}

	sro := &ServiceResourceOutput{
		ID:           _sr.ID,
		CreateTime:   _sr.CreateTime,
		UpdateTime:   _sr.UpdateTime,
		Mode:         _sr.Mode,
		Type:         _sr.Type,
		Name:         _sr.Name,
		DeployerType: _sr.DeployerType,
		Shape:        _sr.Shape,
		Status:       _sr.Status,
		Keys:         _sr.Keys,
	}

	if _sr.Edges.Project != nil {
		sro.Project = ExposeProject(_sr.Edges.Project)
	} else if _sr.ProjectID != "" {
		sro.Project = &ProjectOutput{
			ID: _sr.ProjectID,
		}
	}
	if _sr.Edges.Environment != nil {
		sro.Environment = ExposeEnvironment(_sr.Edges.Environment)
	} else if _sr.EnvironmentID != "" {
		sro.Environment = &EnvironmentOutput{
			ID: _sr.EnvironmentID,
		}
	}
	if _sr.Edges.Service != nil {
		sro.Service = ExposeService(_sr.Edges.Service)
	} else if _sr.ServiceID != "" {
		sro.Service = &ServiceOutput{
			ID: _sr.ServiceID,
		}
	}
	if _sr.Edges.Connector != nil {
		sro.Connector = ExposeConnector(_sr.Edges.Connector)
	} else if _sr.ConnectorID != "" {
		sro.Connector = &ConnectorOutput{
			ID: _sr.ConnectorID,
		}
	}
	if _sr.Edges.Composition != nil {
		sro.Composition = ExposeServiceResource(_sr.Edges.Composition)
	} else if _sr.CompositionID != "" {
		sro.Composition = &ServiceResourceOutput{
			ID: _sr.CompositionID,
		}
	}
	if _sr.Edges.Components != nil {
		sro.Components = ExposeServiceResources(_sr.Edges.Components)
	}
	if _sr.Edges.Class != nil {
		sro.Class = ExposeServiceResource(_sr.Edges.Class)
	} else if _sr.ClassID != "" {
		sro.Class = &ServiceResourceOutput{
			ID: _sr.ClassID,
		}
	}
	if _sr.Edges.Instances != nil {
		sro.Instances = ExposeServiceResources(_sr.Edges.Instances)
	}
	if _sr.Edges.Dependencies != nil {
		sro.Dependencies = ExposeServiceResourceRelationships(_sr.Edges.Dependencies)
	}
	return sro
}

// ExposeServiceResources converts the ServiceResource slice to ServiceResourceOutput pointer slice.
func ExposeServiceResources(_srs []*ServiceResource) []*ServiceResourceOutput {
	if len(_srs) == 0 {
		return nil
	}

	sros := make([]*ServiceResourceOutput, len(_srs))
	for i := range _srs {
		sros[i] = ExposeServiceResource(_srs[i])
	}
	return sros
}
