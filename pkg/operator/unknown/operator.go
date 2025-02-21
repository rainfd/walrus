package unknown

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	optypes "github.com/seal-io/walrus/pkg/operator/types"
)

const OperatorType = "Unknown"

// New returns types.Operator with the given options.
func New(ctx context.Context, opts optypes.CreateOptions) (optypes.Operator, error) {
	if opts.Connector.Category != types.ConnectorCategoryCustom {
		return nil, errors.New("not custom connector")
	}

	return Operator{}, nil
}

type Operator struct{}

func (Operator) Type() optypes.Type {
	return OperatorType
}

func (Operator) IsConnected(ctx context.Context) error {
	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return ""
}

func (Operator) GetKeys(
	ctx context.Context,
	resource *model.ServiceResource,
) (*types.ServiceResourceOperationKeys, error) {
	return nil, nil
}

func (Operator) GetStatus(ctx context.Context, resource *model.ServiceResource) (*status.Status, error) {
	return &status.Status{
		Summary: status.Summary{
			SummaryStatus: "Ready",
		},
	}, nil
}

func (Operator) GetEndpoints(
	ctx context.Context,
	resource *model.ServiceResource,
) ([]types.ServiceResourceEndpoint, error) {
	return nil, nil
}

func (Operator) GetComponents(
	ctx context.Context,
	resource *model.ServiceResource,
) ([]*model.ServiceResource, error) {
	return nil, nil
}

func (Operator) Log(ctx context.Context, s string, options optypes.LogOptions) error {
	return errors.New("cannot log")
}

func (Operator) Exec(ctx context.Context, s string, options optypes.ExecOptions) error {
	return errors.New("cannot execute")
}

func (Operator) Label(ctx context.Context, resource *model.ServiceResource, m map[string]string) error {
	return nil
}
