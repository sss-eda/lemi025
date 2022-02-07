package serial

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/sss-eda/lemi025/internal/domain"
)

// SplitFunc TODO
func SplitFunc(
	instrument *domain.Instrument,
) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// TODO: This might be 8 bytes? Just check
		switch data[0] {
		case 0x3F:
			switch data[1] {
			case 0x30:
				if len(data) < 7 {
					advance = 0
					token = nil
					err = nil
				} else {
					// err = instrument.Config.OnRead(domain.ConfigReadEvent{
					// 	StationNumber: uint8(data[6]),
					// })

					advance = 7
					token = nil
					err = nil
				}
			case 0x31:
				if len(data) < 8 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 8
					token = data[:8]
					err = nil
				}
			case 0x32:
				if len(data) < 8 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 8
					token = data[:8]
					err = nil
				}
			case 0x33:
				if len(data) < 4 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 4
					token = data[:4]
					err = nil
				}
			case 0x34:
				if len(data) < 6 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 6
					token = data[:6]
					err = nil
				}
			case 0x35:
				if len(data) < 2 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 2
					token = data[:2]
					err = nil
				}
			case 0x36:
				if len(data) < 84 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 84
					token = data[:84]
					err = nil
				}
			case 0x37:
				if len(data) < 16 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 16
					token = data[:16]
					err = nil
				}
			case 0x39:
				advance = 2
				token = data[:2]
				err = nil
			case 0x3A:
				// TODO: Something is wrong with the docs here. Make sure?
				if len(data) < 5 {
					advance = 0
					token = nil
					err = nil
				} else {
					advance = 5
					token = data[:5]
					err = nil
				}
			default:
				advance = 1
				token = nil
				err = fmt.Errorf("invalid byte in stream: %v", token)
			}
		case 0x4C:
			if len(data) < 153 {
				advance = 0
				token = nil
				err = nil
			} else {
				if bytes.Equal(data[:4], []byte{0x4c, 0x30, 0x32, 0x35}) {
					if len(data) < 153 {
						advance = 0
						token = nil
						err = nil
					} else {
						instrument.OnDatumAcquired(data[:153])
						advance = 153
						token = data[:153]
						err = nil
					}
					return
				} else {
					advance = 1
					token = nil
					err = nil
				}
				advance = 5
				token = data[:5]
				err = nil
			}

		}

		return
	}
}
