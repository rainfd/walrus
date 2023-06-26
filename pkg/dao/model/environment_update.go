// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/service"
	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// EnvironmentUpdate is the builder for updating Environment entities.
type EnvironmentUpdate struct {
	config
	hooks     []Hook
	mutation  *EnvironmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EnvironmentUpdate builder.
func (eu *EnvironmentUpdate) Where(ps ...predicate.Environment) *EnvironmentUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EnvironmentUpdate) SetName(s string) *EnvironmentUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EnvironmentUpdate) SetDescription(s string) *EnvironmentUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EnvironmentUpdate) SetNillableDescription(s *string) *EnvironmentUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// ClearDescription clears the value of the "description" field.
func (eu *EnvironmentUpdate) ClearDescription() *EnvironmentUpdate {
	eu.mutation.ClearDescription()
	return eu
}

// SetLabels sets the "labels" field.
func (eu *EnvironmentUpdate) SetLabels(m map[string]string) *EnvironmentUpdate {
	eu.mutation.SetLabels(m)
	return eu
}

// ClearLabels clears the value of the "labels" field.
func (eu *EnvironmentUpdate) ClearLabels() *EnvironmentUpdate {
	eu.mutation.ClearLabels()
	return eu
}

// SetAnnotations sets the "annotations" field.
func (eu *EnvironmentUpdate) SetAnnotations(m map[string]string) *EnvironmentUpdate {
	eu.mutation.SetAnnotations(m)
	return eu
}

// ClearAnnotations clears the value of the "annotations" field.
func (eu *EnvironmentUpdate) ClearAnnotations() *EnvironmentUpdate {
	eu.mutation.ClearAnnotations()
	return eu
}

// SetUpdateTime sets the "updateTime" field.
func (eu *EnvironmentUpdate) SetUpdateTime(t time.Time) *EnvironmentUpdate {
	eu.mutation.SetUpdateTime(t)
	return eu
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (eu *EnvironmentUpdate) AddServiceIDs(ids ...oid.ID) *EnvironmentUpdate {
	eu.mutation.AddServiceIDs(ids...)
	return eu
}

// AddServices adds the "services" edges to the Service entity.
func (eu *EnvironmentUpdate) AddServices(s ...*Service) *EnvironmentUpdate {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.AddServiceIDs(ids...)
}

// AddServiceRevisionIDs adds the "serviceRevisions" edge to the ServiceRevision entity by IDs.
func (eu *EnvironmentUpdate) AddServiceRevisionIDs(ids ...oid.ID) *EnvironmentUpdate {
	eu.mutation.AddServiceRevisionIDs(ids...)
	return eu
}

// AddServiceRevisions adds the "serviceRevisions" edges to the ServiceRevision entity.
func (eu *EnvironmentUpdate) AddServiceRevisions(s ...*ServiceRevision) *EnvironmentUpdate {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.AddServiceRevisionIDs(ids...)
}

// Mutation returns the EnvironmentMutation object of the builder.
func (eu *EnvironmentUpdate) Mutation() *EnvironmentMutation {
	return eu.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (eu *EnvironmentUpdate) ClearServices() *EnvironmentUpdate {
	eu.mutation.ClearServices()
	return eu
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (eu *EnvironmentUpdate) RemoveServiceIDs(ids ...oid.ID) *EnvironmentUpdate {
	eu.mutation.RemoveServiceIDs(ids...)
	return eu
}

// RemoveServices removes "services" edges to Service entities.
func (eu *EnvironmentUpdate) RemoveServices(s ...*Service) *EnvironmentUpdate {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.RemoveServiceIDs(ids...)
}

// ClearServiceRevisions clears all "serviceRevisions" edges to the ServiceRevision entity.
func (eu *EnvironmentUpdate) ClearServiceRevisions() *EnvironmentUpdate {
	eu.mutation.ClearServiceRevisions()
	return eu
}

// RemoveServiceRevisionIDs removes the "serviceRevisions" edge to ServiceRevision entities by IDs.
func (eu *EnvironmentUpdate) RemoveServiceRevisionIDs(ids ...oid.ID) *EnvironmentUpdate {
	eu.mutation.RemoveServiceRevisionIDs(ids...)
	return eu
}

// RemoveServiceRevisions removes "serviceRevisions" edges to ServiceRevision entities.
func (eu *EnvironmentUpdate) RemoveServiceRevisions(s ...*ServiceRevision) *EnvironmentUpdate {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.RemoveServiceRevisionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EnvironmentUpdate) Save(ctx context.Context) (int, error) {
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, EnvironmentMutation](ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EnvironmentUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EnvironmentUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EnvironmentUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EnvironmentUpdate) defaults() error {
	if _, ok := eu.mutation.UpdateTime(); !ok {
		if environment.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized environment.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := environment.UpdateDefaultUpdateTime()
		eu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eu *EnvironmentUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := environment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Environment.name": %w`, err)}
		}
	}
	if _, ok := eu.mutation.ProjectID(); eu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Environment.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eu *EnvironmentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvironmentUpdate {
	eu.modifiers = append(eu.modifiers, modifiers...)
	return eu
}

func (eu *EnvironmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(environment.Table, environment.Columns, sqlgraph.NewFieldSpec(environment.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(environment.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(environment.FieldDescription, field.TypeString, value)
	}
	if eu.mutation.DescriptionCleared() {
		_spec.ClearField(environment.FieldDescription, field.TypeString)
	}
	if value, ok := eu.mutation.Labels(); ok {
		_spec.SetField(environment.FieldLabels, field.TypeJSON, value)
	}
	if eu.mutation.LabelsCleared() {
		_spec.ClearField(environment.FieldLabels, field.TypeJSON)
	}
	if value, ok := eu.mutation.Annotations(); ok {
		_spec.SetField(environment.FieldAnnotations, field.TypeJSON, value)
	}
	if eu.mutation.AnnotationsCleared() {
		_spec.ClearField(environment.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := eu.mutation.UpdateTime(); ok {
		_spec.SetField(environment.FieldUpdateTime, field.TypeTime, value)
	}
	if eu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.Service
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedServicesIDs(); len(nodes) > 0 && !eu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ServiceRevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.ServiceRevision
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedServiceRevisionsIDs(); len(nodes) > 0 && !eu.mutation.ServiceRevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.ServiceRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ServiceRevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.ServiceRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = eu.schemaConfig.Environment
	ctx = internal.NewSchemaConfigContext(ctx, eu.schemaConfig)
	_spec.AddModifiers(eu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{environment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EnvironmentUpdateOne is the builder for updating a single Environment entity.
type EnvironmentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EnvironmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (euo *EnvironmentUpdateOne) SetName(s string) *EnvironmentUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EnvironmentUpdateOne) SetDescription(s string) *EnvironmentUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EnvironmentUpdateOne) SetNillableDescription(s *string) *EnvironmentUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// ClearDescription clears the value of the "description" field.
func (euo *EnvironmentUpdateOne) ClearDescription() *EnvironmentUpdateOne {
	euo.mutation.ClearDescription()
	return euo
}

// SetLabels sets the "labels" field.
func (euo *EnvironmentUpdateOne) SetLabels(m map[string]string) *EnvironmentUpdateOne {
	euo.mutation.SetLabels(m)
	return euo
}

// ClearLabels clears the value of the "labels" field.
func (euo *EnvironmentUpdateOne) ClearLabels() *EnvironmentUpdateOne {
	euo.mutation.ClearLabels()
	return euo
}

// SetAnnotations sets the "annotations" field.
func (euo *EnvironmentUpdateOne) SetAnnotations(m map[string]string) *EnvironmentUpdateOne {
	euo.mutation.SetAnnotations(m)
	return euo
}

// ClearAnnotations clears the value of the "annotations" field.
func (euo *EnvironmentUpdateOne) ClearAnnotations() *EnvironmentUpdateOne {
	euo.mutation.ClearAnnotations()
	return euo
}

// SetUpdateTime sets the "updateTime" field.
func (euo *EnvironmentUpdateOne) SetUpdateTime(t time.Time) *EnvironmentUpdateOne {
	euo.mutation.SetUpdateTime(t)
	return euo
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (euo *EnvironmentUpdateOne) AddServiceIDs(ids ...oid.ID) *EnvironmentUpdateOne {
	euo.mutation.AddServiceIDs(ids...)
	return euo
}

// AddServices adds the "services" edges to the Service entity.
func (euo *EnvironmentUpdateOne) AddServices(s ...*Service) *EnvironmentUpdateOne {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.AddServiceIDs(ids...)
}

// AddServiceRevisionIDs adds the "serviceRevisions" edge to the ServiceRevision entity by IDs.
func (euo *EnvironmentUpdateOne) AddServiceRevisionIDs(ids ...oid.ID) *EnvironmentUpdateOne {
	euo.mutation.AddServiceRevisionIDs(ids...)
	return euo
}

// AddServiceRevisions adds the "serviceRevisions" edges to the ServiceRevision entity.
func (euo *EnvironmentUpdateOne) AddServiceRevisions(s ...*ServiceRevision) *EnvironmentUpdateOne {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.AddServiceRevisionIDs(ids...)
}

// Mutation returns the EnvironmentMutation object of the builder.
func (euo *EnvironmentUpdateOne) Mutation() *EnvironmentMutation {
	return euo.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (euo *EnvironmentUpdateOne) ClearServices() *EnvironmentUpdateOne {
	euo.mutation.ClearServices()
	return euo
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (euo *EnvironmentUpdateOne) RemoveServiceIDs(ids ...oid.ID) *EnvironmentUpdateOne {
	euo.mutation.RemoveServiceIDs(ids...)
	return euo
}

// RemoveServices removes "services" edges to Service entities.
func (euo *EnvironmentUpdateOne) RemoveServices(s ...*Service) *EnvironmentUpdateOne {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.RemoveServiceIDs(ids...)
}

// ClearServiceRevisions clears all "serviceRevisions" edges to the ServiceRevision entity.
func (euo *EnvironmentUpdateOne) ClearServiceRevisions() *EnvironmentUpdateOne {
	euo.mutation.ClearServiceRevisions()
	return euo
}

// RemoveServiceRevisionIDs removes the "serviceRevisions" edge to ServiceRevision entities by IDs.
func (euo *EnvironmentUpdateOne) RemoveServiceRevisionIDs(ids ...oid.ID) *EnvironmentUpdateOne {
	euo.mutation.RemoveServiceRevisionIDs(ids...)
	return euo
}

// RemoveServiceRevisions removes "serviceRevisions" edges to ServiceRevision entities.
func (euo *EnvironmentUpdateOne) RemoveServiceRevisions(s ...*ServiceRevision) *EnvironmentUpdateOne {
	ids := make([]oid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.RemoveServiceRevisionIDs(ids...)
}

// Where appends a list predicates to the EnvironmentUpdate builder.
func (euo *EnvironmentUpdateOne) Where(ps ...predicate.Environment) *EnvironmentUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EnvironmentUpdateOne) Select(field string, fields ...string) *EnvironmentUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Environment entity.
func (euo *EnvironmentUpdateOne) Save(ctx context.Context) (*Environment, error) {
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Environment, EnvironmentMutation](ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EnvironmentUpdateOne) SaveX(ctx context.Context) *Environment {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EnvironmentUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EnvironmentUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EnvironmentUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdateTime(); !ok {
		if environment.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized environment.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := environment.UpdateDefaultUpdateTime()
		euo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (euo *EnvironmentUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := environment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Environment.name": %w`, err)}
		}
	}
	if _, ok := euo.mutation.ProjectID(); euo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Environment.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (euo *EnvironmentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvironmentUpdateOne {
	euo.modifiers = append(euo.modifiers, modifiers...)
	return euo
}

func (euo *EnvironmentUpdateOne) sqlSave(ctx context.Context) (_node *Environment, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(environment.Table, environment.Columns, sqlgraph.NewFieldSpec(environment.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Environment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, environment.FieldID)
		for _, f := range fields {
			if !environment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != environment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(environment.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(environment.FieldDescription, field.TypeString, value)
	}
	if euo.mutation.DescriptionCleared() {
		_spec.ClearField(environment.FieldDescription, field.TypeString)
	}
	if value, ok := euo.mutation.Labels(); ok {
		_spec.SetField(environment.FieldLabels, field.TypeJSON, value)
	}
	if euo.mutation.LabelsCleared() {
		_spec.ClearField(environment.FieldLabels, field.TypeJSON)
	}
	if value, ok := euo.mutation.Annotations(); ok {
		_spec.SetField(environment.FieldAnnotations, field.TypeJSON, value)
	}
	if euo.mutation.AnnotationsCleared() {
		_spec.ClearField(environment.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := euo.mutation.UpdateTime(); ok {
		_spec.SetField(environment.FieldUpdateTime, field.TypeTime, value)
	}
	if euo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.Service
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedServicesIDs(); len(nodes) > 0 && !euo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServicesTable,
			Columns: []string{environment.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ServiceRevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.ServiceRevision
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedServiceRevisionsIDs(); len(nodes) > 0 && !euo.mutation.ServiceRevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.ServiceRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ServiceRevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   environment.ServiceRevisionsTable,
			Columns: []string{environment.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.ServiceRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = euo.schemaConfig.Environment
	ctx = internal.NewSchemaConfigContext(ctx, euo.schemaConfig)
	_spec.AddModifiers(euo.modifiers...)
	_node = &Environment{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{environment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
