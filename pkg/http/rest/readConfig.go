package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ReadConfigCommand TODO
type ReadConfigCommand struct {
	URL          string
	InstrumentID string
}

// ReadConfigRequest TODO
type ReadConfigRequest struct {
	InstrumentID string `json:"instrumentID"`
}

// ReadConfigResponse TODO
type ReadConfigResponse struct {
	InstrumentID string `json:"instrumentID"`
}

// Execute TODO
func (command *ReadConfigCommand) Execute() error {
	request, err := json.Marshal(ReadConfigRequest{
		InstrumentID: command.InstrumentID,
	})
	if err != nil {
		return err
	}

	// This needs to be looked at... Can't have command type in the URL
	// http://api.sansa.dev/v1/instrument/bfnl251/command/read-config
	resp, err := http.Post(
		command.URL+"/instrument/"+command.InstrumentID+"/commands/read-config",
		"application/json",
		bytes.NewReader(request),
	)
	if err != nil {
		return err
	}

	response := ReadConfigResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	return nil
}
