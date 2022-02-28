package application

import (
	"context"

	"github.com/sss-eda/lemi025/pkg/datapoints"
)

// API TODO
type API struct {
	Repository datapoints.Repository
}

// AddDataPoint TODO
func (api *API) AddDataPoint(
	ctx context.Context,
	id datapoints.ID,
	datapoint *datapoints.DataPoint,
) error {
	api.Repository.Save(ctx, id, datapoint)
}
