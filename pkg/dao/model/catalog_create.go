// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/catalog"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// CatalogCreate is the builder for creating a Catalog entity.
type CatalogCreate struct {
	config
	mutation   *CatalogMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Catalog
	fromUpsert bool
}

// SetName sets the "name" field.
func (cc *CatalogCreate) SetName(s string) *CatalogCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CatalogCreate) SetDescription(s string) *CatalogCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CatalogCreate) SetNillableDescription(s *string) *CatalogCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetLabels sets the "labels" field.
func (cc *CatalogCreate) SetLabels(m map[string]string) *CatalogCreate {
	cc.mutation.SetLabels(m)
	return cc
}

// SetAnnotations sets the "annotations" field.
func (cc *CatalogCreate) SetAnnotations(m map[string]string) *CatalogCreate {
	cc.mutation.SetAnnotations(m)
	return cc
}

// SetCreateTime sets the "create_time" field.
func (cc *CatalogCreate) SetCreateTime(t time.Time) *CatalogCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CatalogCreate) SetNillableCreateTime(t *time.Time) *CatalogCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CatalogCreate) SetUpdateTime(t time.Time) *CatalogCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CatalogCreate) SetNillableUpdateTime(t *time.Time) *CatalogCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CatalogCreate) SetStatus(s status.Status) *CatalogCreate {
	cc.mutation.SetStatus(s)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CatalogCreate) SetNillableStatus(s *status.Status) *CatalogCreate {
	if s != nil {
		cc.SetStatus(*s)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *CatalogCreate) SetType(s string) *CatalogCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetSource sets the "source" field.
func (cc *CatalogCreate) SetSource(s string) *CatalogCreate {
	cc.mutation.SetSource(s)
	return cc
}

// SetSync sets the "sync" field.
func (cc *CatalogCreate) SetSync(ts *types.CatalogSync) *CatalogCreate {
	cc.mutation.SetSync(ts)
	return cc
}

// SetID sets the "id" field.
func (cc *CatalogCreate) SetID(o object.ID) *CatalogCreate {
	cc.mutation.SetID(o)
	return cc
}

// AddTemplateIDs adds the "templates" edge to the Template entity by IDs.
func (cc *CatalogCreate) AddTemplateIDs(ids ...object.ID) *CatalogCreate {
	cc.mutation.AddTemplateIDs(ids...)
	return cc
}

// AddTemplates adds the "templates" edges to the Template entity.
func (cc *CatalogCreate) AddTemplates(t ...*Template) *CatalogCreate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTemplateIDs(ids...)
}

// Mutation returns the CatalogMutation object of the builder.
func (cc *CatalogCreate) Mutation() *CatalogMutation {
	return cc.mutation
}

// Save creates the Catalog in the database.
func (cc *CatalogCreate) Save(ctx context.Context) (*Catalog, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CatalogCreate) SaveX(ctx context.Context) *Catalog {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CatalogCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CatalogCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CatalogCreate) defaults() error {
	if _, ok := cc.mutation.Labels(); !ok {
		v := catalog.DefaultLabels
		cc.mutation.SetLabels(v)
	}
	if _, ok := cc.mutation.Annotations(); !ok {
		v := catalog.DefaultAnnotations
		cc.mutation.SetAnnotations(v)
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		if catalog.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized catalog.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := catalog.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		if catalog.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized catalog.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := catalog.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CatalogCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Catalog.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := catalog.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Catalog.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Catalog.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Catalog.update_time"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "Catalog.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := catalog.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`model: validator failed for field "Catalog.type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`model: missing required field "Catalog.source"`)}
	}
	if v, ok := cc.mutation.Source(); ok {
		if err := catalog.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "Catalog.source": %w`, err)}
		}
	}
	return nil
}

func (cc *CatalogCreate) sqlSave(ctx context.Context) (*Catalog, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CatalogCreate) createSpec() (*Catalog, *sqlgraph.CreateSpec) {
	var (
		_node = &Catalog{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(catalog.Table, sqlgraph.NewFieldSpec(catalog.FieldID, field.TypeString))
	)
	_spec.Schema = cc.schemaConfig.Catalog
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(catalog.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(catalog.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Labels(); ok {
		_spec.SetField(catalog.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	if value, ok := cc.mutation.Annotations(); ok {
		_spec.SetField(catalog.FieldAnnotations, field.TypeJSON, value)
		_node.Annotations = value
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(catalog.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(catalog.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(catalog.FieldStatus, field.TypeJSON, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(catalog.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.Source(); ok {
		_spec.SetField(catalog.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := cc.mutation.Sync(); ok {
		_spec.SetField(catalog.FieldSync, field.TypeJSON, value)
		_node.Sync = value
	}
	if nodes := cc.mutation.TemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   catalog.TemplatesTable,
			Columns: []string{catalog.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(template.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.Template
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (cc *CatalogCreate) Set(obj *Catalog) *CatalogCreate {
	// Required.
	cc.SetName(obj.Name)
	cc.SetType(obj.Type)
	cc.SetSource(obj.Source)

	// Optional.
	if obj.Description != "" {
		cc.SetDescription(obj.Description)
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		cc.SetLabels(obj.Labels)
	}
	if !reflect.ValueOf(obj.Annotations).IsZero() {
		cc.SetAnnotations(obj.Annotations)
	}
	if obj.CreateTime != nil {
		cc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		cc.SetUpdateTime(*obj.UpdateTime)
	}
	if !reflect.ValueOf(obj.Status).IsZero() {
		cc.SetStatus(obj.Status)
	}
	if !reflect.ValueOf(obj.Sync).IsZero() {
		cc.SetSync(obj.Sync)
	}

	// Record the given object.
	cc.object = obj

	return cc
}

// getClientSet returns the ClientSet for the given builder.
func (cc *CatalogCreate) getClientSet() (mc ClientSet) {
	if _, ok := cc.config.driver.(*txDriver); ok {
		tx := &Tx{config: cc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: cc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Catalog entity,
// which is always good for cascading create operations.
func (cc *CatalogCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) (*Catalog, error) {
	obj, err := cc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := cc.getClientSet()
	if cc.fromUpsert {
		q := mc.Catalogs().Query().
			Where(
				catalog.Name(obj.Name),
			)
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of Catalog entity: %w", err)
		}
	}

	if x := cc.object; x != nil {
		if _, set := cc.mutation.Field(catalog.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := cc.mutation.Field(catalog.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := cc.mutation.Field(catalog.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := cc.mutation.Field(catalog.FieldType); set {
			obj.Type = x.Type
		}
		if _, set := cc.mutation.Field(catalog.FieldSource); set {
			obj.Source = x.Source
		}
		if _, set := cc.mutation.Field(catalog.FieldSync); set {
			obj.Sync = x.Sync
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (cc *CatalogCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) *Catalog {
	obj, err := cc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (cc *CatalogCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) error {
	_, err := cc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (cc *CatalogCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) {
	if err := cc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the CatalogCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (ccb *CatalogCreateBulk) Set(objs ...*Catalog) *CatalogCreateBulk {
	if len(objs) != 0 {
		client := NewCatalogClient(ccb.config)

		ccb.builders = make([]*CatalogCreate, len(objs))
		for i := range objs {
			ccb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		ccb.objects = objs
	}

	return ccb
}

// getClientSet returns the ClientSet for the given builder.
func (ccb *CatalogCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := ccb.config.driver.(*txDriver); ok {
		tx := &Tx{config: ccb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: ccb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Catalog entities,
// which is always good for cascading create operations.
func (ccb *CatalogCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) ([]*Catalog, error) {
	objs, err := ccb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := ccb.getClientSet()
	if ccb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.Catalogs().Query().
				Where(
					catalog.Name(obj.Name),
				)
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of Catalog entity: %w", err)
			}
		}
	}

	if x := ccb.objects; x != nil {
		for i := range x {
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldDescription); set {
				objs[i].Description = x[i].Description
			}
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldStatus); set {
				objs[i].Status = x[i].Status
			}
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldType); set {
				objs[i].Type = x[i].Type
			}
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldSource); set {
				objs[i].Source = x[i].Source
			}
			if _, set := ccb.builders[i].mutation.Field(catalog.FieldSync); set {
				objs[i].Sync = x[i].Sync
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (ccb *CatalogCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) []*Catalog {
	objs, err := ccb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (ccb *CatalogCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) error {
	_, err := ccb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (ccb *CatalogCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Catalog) error) {
	if err := ccb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *CatalogUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Catalog) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for CatalogUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *CatalogUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Catalog) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *CatalogUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Catalog) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the CatalogUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for CatalogUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *CatalogUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Catalog) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Catalog.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CatalogUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cc *CatalogCreate) OnConflict(opts ...sql.ConflictOption) *CatalogUpsertOne {
	cc.conflict = opts
	return &CatalogUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Catalog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CatalogCreate) OnConflictColumns(columns ...string) *CatalogUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CatalogUpsertOne{
		create: cc,
	}
}

type (
	// CatalogUpsertOne is the builder for "upsert"-ing
	//  one Catalog node.
	CatalogUpsertOne struct {
		create *CatalogCreate
	}

	// CatalogUpsert is the "OnConflict" setter.
	CatalogUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *CatalogUpsert) SetDescription(v string) *CatalogUpsert {
	u.Set(catalog.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateDescription() *CatalogUpsert {
	u.SetExcluded(catalog.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *CatalogUpsert) ClearDescription() *CatalogUpsert {
	u.SetNull(catalog.FieldDescription)
	return u
}

// SetLabels sets the "labels" field.
func (u *CatalogUpsert) SetLabels(v map[string]string) *CatalogUpsert {
	u.Set(catalog.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateLabels() *CatalogUpsert {
	u.SetExcluded(catalog.FieldLabels)
	return u
}

// ClearLabels clears the value of the "labels" field.
func (u *CatalogUpsert) ClearLabels() *CatalogUpsert {
	u.SetNull(catalog.FieldLabels)
	return u
}

// SetAnnotations sets the "annotations" field.
func (u *CatalogUpsert) SetAnnotations(v map[string]string) *CatalogUpsert {
	u.Set(catalog.FieldAnnotations, v)
	return u
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateAnnotations() *CatalogUpsert {
	u.SetExcluded(catalog.FieldAnnotations)
	return u
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *CatalogUpsert) ClearAnnotations() *CatalogUpsert {
	u.SetNull(catalog.FieldAnnotations)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CatalogUpsert) SetUpdateTime(v time.Time) *CatalogUpsert {
	u.Set(catalog.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateUpdateTime() *CatalogUpsert {
	u.SetExcluded(catalog.FieldUpdateTime)
	return u
}

// SetStatus sets the "status" field.
func (u *CatalogUpsert) SetStatus(v status.Status) *CatalogUpsert {
	u.Set(catalog.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateStatus() *CatalogUpsert {
	u.SetExcluded(catalog.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *CatalogUpsert) ClearStatus() *CatalogUpsert {
	u.SetNull(catalog.FieldStatus)
	return u
}

// SetSync sets the "sync" field.
func (u *CatalogUpsert) SetSync(v *types.CatalogSync) *CatalogUpsert {
	u.Set(catalog.FieldSync, v)
	return u
}

// UpdateSync sets the "sync" field to the value that was provided on create.
func (u *CatalogUpsert) UpdateSync() *CatalogUpsert {
	u.SetExcluded(catalog.FieldSync)
	return u
}

// ClearSync clears the value of the "sync" field.
func (u *CatalogUpsert) ClearSync() *CatalogUpsert {
	u.SetNull(catalog.FieldSync)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Catalog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(catalog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CatalogUpsertOne) UpdateNewValues() *CatalogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(catalog.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(catalog.FieldName)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(catalog.FieldCreateTime)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(catalog.FieldType)
		}
		if _, exists := u.create.mutation.Source(); exists {
			s.SetIgnore(catalog.FieldSource)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Catalog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CatalogUpsertOne) Ignore() *CatalogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CatalogUpsertOne) DoNothing() *CatalogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CatalogCreate.OnConflict
// documentation for more info.
func (u *CatalogUpsertOne) Update(set func(*CatalogUpsert)) *CatalogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CatalogUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *CatalogUpsertOne) SetDescription(v string) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateDescription() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *CatalogUpsertOne) ClearDescription() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *CatalogUpsertOne) SetLabels(v map[string]string) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateLabels() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *CatalogUpsertOne) ClearLabels() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearLabels()
	})
}

// SetAnnotations sets the "annotations" field.
func (u *CatalogUpsertOne) SetAnnotations(v map[string]string) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetAnnotations(v)
	})
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateAnnotations() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateAnnotations()
	})
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *CatalogUpsertOne) ClearAnnotations() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearAnnotations()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CatalogUpsertOne) SetUpdateTime(v time.Time) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateUpdateTime() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *CatalogUpsertOne) SetStatus(v status.Status) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateStatus() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *CatalogUpsertOne) ClearStatus() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearStatus()
	})
}

// SetSync sets the "sync" field.
func (u *CatalogUpsertOne) SetSync(v *types.CatalogSync) *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.SetSync(v)
	})
}

// UpdateSync sets the "sync" field to the value that was provided on create.
func (u *CatalogUpsertOne) UpdateSync() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateSync()
	})
}

// ClearSync clears the value of the "sync" field.
func (u *CatalogUpsertOne) ClearSync() *CatalogUpsertOne {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearSync()
	})
}

// Exec executes the query.
func (u *CatalogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for CatalogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CatalogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CatalogUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: CatalogUpsertOne.ID is not supported by MySQL driver. Use CatalogUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CatalogUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CatalogCreateBulk is the builder for creating many Catalog entities in bulk.
type CatalogCreateBulk struct {
	config
	builders   []*CatalogCreate
	conflict   []sql.ConflictOption
	objects    []*Catalog
	fromUpsert bool
}

// Save creates the Catalog entities in the database.
func (ccb *CatalogCreateBulk) Save(ctx context.Context) ([]*Catalog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Catalog, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CatalogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CatalogCreateBulk) SaveX(ctx context.Context) []*Catalog {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CatalogCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CatalogCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Catalog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CatalogUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ccb *CatalogCreateBulk) OnConflict(opts ...sql.ConflictOption) *CatalogUpsertBulk {
	ccb.conflict = opts
	return &CatalogUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Catalog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CatalogCreateBulk) OnConflictColumns(columns ...string) *CatalogUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CatalogUpsertBulk{
		create: ccb,
	}
}

// CatalogUpsertBulk is the builder for "upsert"-ing
// a bulk of Catalog nodes.
type CatalogUpsertBulk struct {
	create *CatalogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Catalog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(catalog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CatalogUpsertBulk) UpdateNewValues() *CatalogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(catalog.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(catalog.FieldName)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(catalog.FieldCreateTime)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(catalog.FieldType)
			}
			if _, exists := b.mutation.Source(); exists {
				s.SetIgnore(catalog.FieldSource)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Catalog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CatalogUpsertBulk) Ignore() *CatalogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CatalogUpsertBulk) DoNothing() *CatalogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CatalogCreateBulk.OnConflict
// documentation for more info.
func (u *CatalogUpsertBulk) Update(set func(*CatalogUpsert)) *CatalogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CatalogUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *CatalogUpsertBulk) SetDescription(v string) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateDescription() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *CatalogUpsertBulk) ClearDescription() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *CatalogUpsertBulk) SetLabels(v map[string]string) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateLabels() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *CatalogUpsertBulk) ClearLabels() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearLabels()
	})
}

// SetAnnotations sets the "annotations" field.
func (u *CatalogUpsertBulk) SetAnnotations(v map[string]string) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetAnnotations(v)
	})
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateAnnotations() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateAnnotations()
	})
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *CatalogUpsertBulk) ClearAnnotations() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearAnnotations()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CatalogUpsertBulk) SetUpdateTime(v time.Time) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateUpdateTime() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *CatalogUpsertBulk) SetStatus(v status.Status) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateStatus() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *CatalogUpsertBulk) ClearStatus() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearStatus()
	})
}

// SetSync sets the "sync" field.
func (u *CatalogUpsertBulk) SetSync(v *types.CatalogSync) *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.SetSync(v)
	})
}

// UpdateSync sets the "sync" field to the value that was provided on create.
func (u *CatalogUpsertBulk) UpdateSync() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.UpdateSync()
	})
}

// ClearSync clears the value of the "sync" field.
func (u *CatalogUpsertBulk) ClearSync() *CatalogUpsertBulk {
	return u.Update(func(s *CatalogUpsert) {
		s.ClearSync()
	})
}

// Exec executes the query.
func (u *CatalogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the CatalogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for CatalogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CatalogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
