package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/application"
)

// ReadConfig TODO - We may depend on the application layer,
// because we are now in the infrastructure layer!
func ReadConfig(
	w io.Writer,
) func(application.ReadConfigRequest) (application.ReadConfigResponse, error) {
	return func(
		request application.ReadConfigRequest,
	) (application.ReadConfigResponse, error) {
		_, err := w.Write([]byte{0x3D, 0x30})
		if err != nil {
			return application.ReadConfigResponse{}, err
		}

		return application.ReadConfigResponse{}, nil
	}
}
