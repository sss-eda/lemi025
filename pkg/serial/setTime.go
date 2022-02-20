package serial

import (
	"fmt"
	"io"
	"log"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025"
)

// SetTime TODO
func SetTime(
	writer io.Writer,
) func(w lemi025.SetTimeResponseWriter, r *lemi025.SetTimeRequest) {
	return func(w lemi025.SetTimeResponseWriter, r *lemi025.SetTimeRequest) {
		response := lemi025.SetTimeResponse{}
		year, err := bcd.Encode(uint8(r.Year - 2000))
		if err != nil {
			response.Error = fmt.Errorf("failed to decode year bcd value: %v", err)
			w.Write(response)
			return
		}
		month, err := bcd.Encode(r.Month)
		if err != nil {
			response.Error = fmt.Errorf("failed to decode month bcd value: %v", err)
			w.Write(response)
			return
		}
		day, err := bcd.Encode(r.Day)
		if err != nil {
			response.Error = fmt.Errorf("failed to decode day bcd value: %v", err)
			w.Write(response)
			return
		}
		hour, err := bcd.Encode(r.Hour)
		if err != nil {
			response.Error = fmt.Errorf("failed to decode hour bcd value: %v", err)
			w.Write(response)
			return
		}
		minute, err := bcd.Encode(r.Minute)
		if err != nil {
			response.Error = fmt.Errorf("failed to decode minute bcd value: %v", err)
			w.Write(response)
			return
		}
		second, err := bcd.Encode(r.Second)
		if err != nil {
			response.Error = fmt.Errorf("failed to decode second bcd value: %v", err)
			w.Write(response)
			return
		}
		_, err = writer.Write([]byte{0x3D, 0x32, year, month, day, hour, minute, second})
		if err != nil {
			response.Error = fmt.Errorf("failed to write to serial port: %v", err)
		}

		err = w.Write(response)
		if err != nil {
			log.Println(err)
		}
	}
}
