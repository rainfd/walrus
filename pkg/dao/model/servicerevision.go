// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// ServiceRevision is the model entity for the ServiceRevision schema.
type ServiceRevision struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// ID of the environment to which the revision belongs.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the service to which the revision belongs.
	ServiceID object.ID `json:"service_id,omitempty"`
	// Name of the template.
	TemplateName string `json:"template_name,omitempty"`
	// Version of the template.
	TemplateVersion string `json:"template_version,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `json:"variables,omitempty"`
	// Input plan of the revision.
	InputPlan string `json:"-"`
	// Output of the revision.
	Output string `json:"-"`
	// Type of deployer.
	DeployerType string `json:"deployer_type,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `json:"previous_required_providers,omitempty"`
	// Tags of the revision.
	Tags []string `json:"tags,omitempty"`
	// Record of the revision.
	Record string `json:"record,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceRevisionQuery when eager-loading is set.
	Edges        ServiceRevisionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ServiceRevisionEdges holds the relations/edges for other nodes in the graph.
type ServiceRevisionEdges struct {
	// Project to which the revision belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the revision deploys.
	Environment *Environment `json:"environment,omitempty"`
	// Service to which the revision belongs.
	Service *Service `json:"service,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceRevisionEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceRevisionEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// ServiceOrErr returns the Service value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceRevisionEdges) ServiceOrErr() (*Service, error) {
	if e.loadedTypes[2] {
		if e.Service == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: service.Label}
		}
		return e.Service, nil
	}
	return nil, &NotLoadedError{edge: "service"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceRevision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case servicerevision.FieldStatus, servicerevision.FieldPreviousRequiredProviders, servicerevision.FieldTags:
			values[i] = new([]byte)
		case servicerevision.FieldVariables:
			values[i] = new(crypto.Map[string, string])
		case servicerevision.FieldID, servicerevision.FieldProjectID, servicerevision.FieldEnvironmentID, servicerevision.FieldServiceID:
			values[i] = new(object.ID)
		case servicerevision.FieldAttributes:
			values[i] = new(property.Values)
		case servicerevision.FieldDuration:
			values[i] = new(sql.NullInt64)
		case servicerevision.FieldTemplateName, servicerevision.FieldTemplateVersion, servicerevision.FieldInputPlan, servicerevision.FieldOutput, servicerevision.FieldDeployerType, servicerevision.FieldRecord:
			values[i] = new(sql.NullString)
		case servicerevision.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceRevision fields.
func (sr *ServiceRevision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case servicerevision.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sr.ID = *value
			}
		case servicerevision.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sr.CreateTime = new(time.Time)
				*sr.CreateTime = value.Time
			}
		case servicerevision.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sr.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case servicerevision.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				sr.ProjectID = *value
			}
		case servicerevision.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				sr.EnvironmentID = *value
			}
		case servicerevision.FieldServiceID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field service_id", values[i])
			} else if value != nil {
				sr.ServiceID = *value
			}
		case servicerevision.FieldTemplateName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_name", values[i])
			} else if value.Valid {
				sr.TemplateName = value.String
			}
		case servicerevision.FieldTemplateVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_version", values[i])
			} else if value.Valid {
				sr.TemplateVersion = value.String
			}
		case servicerevision.FieldAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value != nil {
				sr.Attributes = *value
			}
		case servicerevision.FieldVariables:
			if value, ok := values[i].(*crypto.Map[string, string]); !ok {
				return fmt.Errorf("unexpected type %T for field variables", values[i])
			} else if value != nil {
				sr.Variables = *value
			}
		case servicerevision.FieldInputPlan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input_plan", values[i])
			} else if value.Valid {
				sr.InputPlan = value.String
			}
		case servicerevision.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				sr.Output = value.String
			}
		case servicerevision.FieldDeployerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployer_type", values[i])
			} else if value.Valid {
				sr.DeployerType = value.String
			}
		case servicerevision.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				sr.Duration = int(value.Int64)
			}
		case servicerevision.FieldPreviousRequiredProviders:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field previous_required_providers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sr.PreviousRequiredProviders); err != nil {
					return fmt.Errorf("unmarshal field previous_required_providers: %w", err)
				}
			}
		case servicerevision.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sr.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case servicerevision.FieldRecord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record", values[i])
			} else if value.Valid {
				sr.Record = value.String
			}
		default:
			sr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServiceRevision.
// This includes values selected through modifiers, order, etc.
func (sr *ServiceRevision) Value(name string) (ent.Value, error) {
	return sr.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the ServiceRevision entity.
func (sr *ServiceRevision) QueryProject() *ProjectQuery {
	return NewServiceRevisionClient(sr.config).QueryProject(sr)
}

// QueryEnvironment queries the "environment" edge of the ServiceRevision entity.
func (sr *ServiceRevision) QueryEnvironment() *EnvironmentQuery {
	return NewServiceRevisionClient(sr.config).QueryEnvironment(sr)
}

// QueryService queries the "service" edge of the ServiceRevision entity.
func (sr *ServiceRevision) QueryService() *ServiceQuery {
	return NewServiceRevisionClient(sr.config).QueryService(sr)
}

// Update returns a builder for updating this ServiceRevision.
// Note that you need to call ServiceRevision.Unwrap() before calling this method if this ServiceRevision
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *ServiceRevision) Update() *ServiceRevisionUpdateOne {
	return NewServiceRevisionClient(sr.config).UpdateOne(sr)
}

// Unwrap unwraps the ServiceRevision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *ServiceRevision) Unwrap() *ServiceRevision {
	_tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("model: ServiceRevision is not a transactional entity")
	}
	sr.config.driver = _tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *ServiceRevision) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceRevision(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sr.ID))
	if v := sr.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sr.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.EnvironmentID))
	builder.WriteString(", ")
	builder.WriteString("service_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.ServiceID))
	builder.WriteString(", ")
	builder.WriteString("template_name=")
	builder.WriteString(sr.TemplateName)
	builder.WriteString(", ")
	builder.WriteString("template_version=")
	builder.WriteString(sr.TemplateVersion)
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(fmt.Sprintf("%v", sr.Attributes))
	builder.WriteString(", ")
	builder.WriteString("variables=")
	builder.WriteString(fmt.Sprintf("%v", sr.Variables))
	builder.WriteString(", ")
	builder.WriteString("input_plan=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("output=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("deployer_type=")
	builder.WriteString(sr.DeployerType)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", sr.Duration))
	builder.WriteString(", ")
	builder.WriteString("previous_required_providers=")
	builder.WriteString(fmt.Sprintf("%v", sr.PreviousRequiredProviders))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", sr.Tags))
	builder.WriteString(", ")
	builder.WriteString("record=")
	builder.WriteString(sr.Record)
	builder.WriteByte(')')
	return builder.String()
}

// ServiceRevisions is a parsable slice of ServiceRevision.
type ServiceRevisions []*ServiceRevision
