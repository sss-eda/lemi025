package application

import (
	"fmt"

	"github.com/sss-eda/lemi025/pkg/datapoints"
)

// DataPointAlreadyExistsError TODO
type DataPointAlreadyExistsError struct {
	ID datapoints.ID
}

// Error TODO
func (err DataPointAlreadyExistsError) Error() string {
	return fmt.Sprintf("datapoint with id=%v already exists", err.ID.String())
}
