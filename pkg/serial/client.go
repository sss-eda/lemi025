package serial

import (
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/pkg/driving"
)

// Client TODO
type Client struct {
	io.Writer
}

// ReadConfig TODO
func (client *Client) ReadConfig(
	input driving.ReadConfig,
) error {
	_, err := client.Write([]byte{0x3D, 0x30})

	return err
}

// ReadTime TODO
func (client *Client) ReadTime(
	input driving.ReadConfig,
) error {
	_, err := client.Write([]byte{0x3D, 0x31})

	return err
}

// SetTime TODO
func (client *Client) SetTime(
	input driving.SetTime,
) error {
	buffer := make([]byte, 8)

	buffer[0] = 0x3D
	buffer[1] = 0x32
	buffer[2] = bcd.Encode(byte(input.Year - 2000))
	buffer[3] = bcd.Encode(byte(input.Day))
	buffer[4] = bcd.Encode(byte(input.Month))
	buffer[5] = bcd.Encode(byte(input.Hour))
	buffer[6] = bcd.Encode(byte(input.Minute))
	buffer[7] = bcd.Encode(byte(input.Second))

	_, err := client.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
