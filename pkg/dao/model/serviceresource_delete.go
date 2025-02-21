// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
)

// ServiceResourceDelete is the builder for deleting a ServiceResource entity.
type ServiceResourceDelete struct {
	config
	hooks    []Hook
	mutation *ServiceResourceMutation
}

// Where appends a list predicates to the ServiceResourceDelete builder.
func (srd *ServiceResourceDelete) Where(ps ...predicate.ServiceResource) *ServiceResourceDelete {
	srd.mutation.Where(ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *ServiceResourceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srd.sqlExec, srd.mutation, srd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *ServiceResourceDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *ServiceResourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(serviceresource.Table, sqlgraph.NewFieldSpec(serviceresource.FieldID, field.TypeString))
	_spec.Node.Schema = srd.schemaConfig.ServiceResource
	ctx = internal.NewSchemaConfigContext(ctx, srd.schemaConfig)
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srd.mutation.done = true
	return affected, err
}

// ServiceResourceDeleteOne is the builder for deleting a single ServiceResource entity.
type ServiceResourceDeleteOne struct {
	srd *ServiceResourceDelete
}

// Where appends a list predicates to the ServiceResourceDelete builder.
func (srdo *ServiceResourceDeleteOne) Where(ps ...predicate.ServiceResource) *ServiceResourceDeleteOne {
	srdo.srd.mutation.Where(ps...)
	return srdo
}

// Exec executes the deletion query.
func (srdo *ServiceResourceDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{serviceresource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *ServiceResourceDeleteOne) ExecX(ctx context.Context) {
	if err := srdo.Exec(ctx); err != nil {
		panic(err)
	}
}
