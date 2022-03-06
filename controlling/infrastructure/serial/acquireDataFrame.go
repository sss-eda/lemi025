package serial

import (
	"bytes"
	"fmt"

	"github.com/sss-eda/lemi025/controlling"
)

// OnDataFrameAcquired TODO
func OnDataFrameAcquired(
	handle func(*controlling.DataFrameAcquiredEvent) error,
) func([]byte) error {
	return func(data []byte) error {
		if len(data) != 153 {
			return fmt.Errorf("invalid data length (expected slice of length 153, got: %v)", len(data))
		}
		if !bytes.Equal(data[:4], []byte{0x4C, 0x30, 0x32, 0x35}) {
			return fmt.Errorf("invalid data (expected slice to start with 0x4C303235, got: %v)", data[:4])
		}
		return handle(&controlling.DataFrameAcquiredEvent{})
	}
}
