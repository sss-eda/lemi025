package openapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client TODO
type Client struct {
	url string
}

// NewClient TODO
func NewClient(
	url string,
) (*Client, error) {
	return &Client{
		url: url,
	}, nil
}

// ReadConfig TODO
func (client *Client) ReadConfig(
	id string,
) error {
	request := ReadConfigRequest{}
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	responseData, err := http.Post(
		client.url+"/lemi025/"+id+"/read-config",
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}
	defer responseData.Body.Close()

	dec := json.NewDecoder(responseData.Body)
	response := ReadConfigResponse{}
	err = dec.Decode(&response)
	if err != nil {
		return err
	}

	return nil
}

// ReadConfigRequest TODO
type ReadConfigRequest struct{}

// ReadConfigResponse TODO
type ReadConfigResponse struct{}
