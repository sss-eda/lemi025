package serial

import (
	"io"

	"github.com/sss-eda/lemi025/controlling/application/readconfig"
)

// ReadConfig TODO
func ReadConfig(
	writer io.Writer,
) readconfig.HandlerFunc {
	return func(respond readconfig.ResponseWriter, request *readconfig.Request) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x30 // Token "0" indicating command type: read config
		buffer[2] = 0xFF // XX - Don't care
		buffer[3] = 0xFF // XX - Don't care

		_, err := writer.Write(buffer)
		if err != nil {
			return err
		}

		return respond(&readconfig.Response{})
	}
}
