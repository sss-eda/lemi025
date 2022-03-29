package json

import (
	"encoding/json"

	"github.com/sss-eda/lemi025"
)

// Serialize TODO
func Serialize[Payload lemi025.CommandPayloads | lemi025.EventPayloads](
	payload *Payload,
) ([]byte, error) {
	return json.Marshal(payload)
}

// Deserialize TODO
func Deserialize[Payload lemi025.CommandPayloads | lemi025.EventPayloads](
	payload []byte,
) (*Payload, error) {
	var result Payload
	err := json.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
