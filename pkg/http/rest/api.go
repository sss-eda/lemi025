package rest

import (
	"net/http"

	"github.com/sss-eda/lemi025"
)

// API TODO
type API struct {
	Connection *lemi025.Connection
}

func (api *API) getData(
	w http.ResponseWriter,
	r *http.Request,
) {
	api.Subscribe(lemi025.DataFrameAcquiredEvent)
}

func (api *API) readTime(
	w http.ResponseWriter,
	r *http.Request,
) {
	api.Connection.ReadTime(&lemi025.ReadTimeCommand{})
}
