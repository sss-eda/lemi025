package serial

import (
	"fmt"
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfig TODO
func ReadConfig(
	writer io.Writer,
) func(lemi025.ReadConfigRequest) lemi025.ReadConfigResponse {
	return func(request lemi025.ReadConfigRequest) lemi025.ReadConfigResponse {
		response := lemi025.ReadConfigResponse{}
		_, err := writer.Write([]byte{0x3D, 0x30})
		if err != nil {
			response.Error = fmt.Errorf("failed to send command over serial: %v", err)
		}
		return response
	}
}
