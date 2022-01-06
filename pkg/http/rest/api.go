package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sss-eda/lemi025/pkg/driving"
)

// API TODO
type API struct {
	driver  *driving.Service
	devices map[string]*driving.Device
}

// ReadConfig TODO
func (api *API) ReadConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	request := ReadConfigRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Load device
	// api.devices.Load(request.DeviceID)
	device := api.devices[request.DeviceID]

	api.driver.ReadConfig(device)
}
