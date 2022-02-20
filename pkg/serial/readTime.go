package serial

import (
	"fmt"
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadTime TODO
func ReadTime(
	writer io.Writer,
) func(lemi025.ReadTimeRequest) lemi025.ReadTimeResponse {
	return func(request lemi025.ReadTimeRequest) lemi025.ReadTimeResponse {
		response := lemi025.ReadTimeResponse{}
		_, err := writer.Write([]byte{0x3D, 0x31})
		if err != nil {
			response.Error = fmt.Errorf("failed to send command over serial: %v", err)
		}
		return response
	}
}
