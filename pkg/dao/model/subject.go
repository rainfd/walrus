// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/subject"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID oid.ID `json:"id,omitempty" sql:"id"`
	// CreateTime holds the value of the "createTime" field.
	CreateTime *time.Time `json:"createTime,omitempty" sql:"createTime"`
	// UpdateTime holds the value of the "updateTime" field.
	UpdateTime *time.Time `json:"updateTime,omitempty" sql:"updateTime"`
	// The kind of the subject.
	Kind string `json:"kind,omitempty" sql:"kind"`
	// The domain of the subject.
	Domain string `json:"domain,omitempty" sql:"domain"`
	// The name of the subject.
	Name string `json:"name,omitempty" sql:"name"`
	// The detail of the subject.
	Description string `json:"description,omitempty" sql:"description"`
	// Indicate whether the subject is builtin, decide when creating.
	Builtin bool `json:"builtin,omitempty" sql:"builtin"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges        SubjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// Tokens that belong to the subject.
	Tokens []*Token `json:"tokens,omitempty" sql:"tokens"`
	// Roles holds the value of the roles edge.
	Roles []*SubjectRoleRelationship `json:"roles,omitempty" sql:"roles"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[0] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) RolesOrErr() ([]*SubjectRoleRelationship, error) {
	if e.loadedTypes[1] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			values[i] = new(oid.ID)
		case subject.FieldBuiltin:
			values[i] = new(sql.NullBool)
		case subject.FieldKind, subject.FieldDomain, subject.FieldName, subject.FieldDescription:
			values[i] = new(sql.NullString)
		case subject.FieldCreateTime, subject.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			if value, ok := values[i].(*oid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subject.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				s.CreateTime = new(time.Time)
				*s.CreateTime = value.Time
			}
		case subject.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updateTime", values[i])
			} else if value.Valid {
				s.UpdateTime = new(time.Time)
				*s.UpdateTime = value.Time
			}
		case subject.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				s.Kind = value.String
			}
		case subject.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case subject.FieldBuiltin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field builtin", values[i])
			} else if value.Valid {
				s.Builtin = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subject.
// This includes values selected through modifiers, order, etc.
func (s *Subject) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTokens queries the "tokens" edge of the Subject entity.
func (s *Subject) QueryTokens() *TokenQuery {
	return NewSubjectClient(s.config).QueryTokens(s)
}

// QueryRoles queries the "roles" edge of the Subject entity.
func (s *Subject) QueryRoles() *SubjectRoleRelationshipQuery {
	return NewSubjectClient(s.config).QueryRoles(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return NewSubjectClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("model: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	if v := s.CreateTime; v != nil {
		builder.WriteString("createTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.UpdateTime; v != nil {
		builder.WriteString("updateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(s.Kind)
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(s.Domain)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("builtin=")
	builder.WriteString(fmt.Sprintf("%v", s.Builtin))
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject
