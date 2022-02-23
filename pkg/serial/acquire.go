package serial

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025"
)

type state int

const (
	idle state = iota
	data
	resp
	readConfig
	readTime
	setTime
	setCoefficients1
	readCoefficients1
	setCoefficients2
	readCoefficients2
	readGPSData
	stopSystem
	startSystem
	checkFLASH
	setDACx
	setDACy
	setDACz
)

func Acquire(
	reader io.Reader,
) func(lemi025.AcquireReponseWriter, *lemi025.AcquireRequest) {
	return func(w lemi025.AcquireReponseWriter, r *lemi025.AcquireRequest) {
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanBytes)

		var next state = idle
		var cursor uint8 = 0
		var buffer []byte = []byte{}
		for scanner.Scan() {
			token := scanner.Bytes()[0]
			switch next {
			case idle:
				switch token {
				case 0x3F:
					next = resp
				case 0x4C:
					next = data
				default:
					continue
				}
				cursor++
			case data:
				cursor++
			case resp:
				switch token {
				case 0x30:
					next = readConfig
				case 0x31:
					next = readTime
				case 0x32:
					next = setTime
				case 0x33:
					next = setCoefficients1
				case 0x34:
					next = readCoefficients1
				case 0x35:
					next = setCoefficients2
				case 0x36:
					next = readCoefficients2
				case 0x37:
					next = readGPSData
				case 0x38:
					next = stopSystem
				case 0x39:
					next = startSystem
				case 0x3A:
					next = checkFLASH
				case 0x3D:
					next = setDACx
				case 0x3E:
					next = setDACy
				case 0x3F:
					next = setDACz
				default:
					continue
				}
				cursor++
			case readConfig:
				switch len(buffer) {
				case 0:
					if token != 0x30 {
						err := w.Write(lemi025.AcquireResponse{
							Error: fmt.Errorf("invalid token for config: %v", token),
						})
						if err != nil {
							log.Println(err)
						}
						next = idle
					}
				case 1:
					if token != 0x32 {
						err := w.Write(lemi025.AcquireResponse{
							Error: fmt.Errorf("invalid token for config: %v", token),
						})
						if err != nil {
							log.Println(err)
						}
						next = idle
					}
				case 2:
					if token != 0x35 {
						err := w.Write(lemi025.AcquireResponse{
							Error: fmt.Errorf("invalid token for config: %v", token),
						})
						if err != nil {
							log.Println(err)
						}
						next = idle
					}
				case 3:
					if token != 0x20 {
						err := w.Write(lemi025.AcquireResponse{
							Error: fmt.Errorf("invalid token for config: %v", token),
						})
						if err != nil {
							log.Println(err)
						}
						next = idle
					}
				case 4:
					sn, err := bcd.Decode(token)
					if err != nil {
						err2 := w.Write(lemi025.AcquireResponse{
							Error: fmt.Errorf("invalid token for config: %v", token),
						})
						log.Println(err2)
					}
					err = w.Write(lemi025.AcquireResponse{
						Error: nil,
						Payload: lemi025.ReadConfigResponse{
							StationNumber: sn,
						},
					})
					if err != nil {
						log.Println(err)
					}

				}
			case readTime:
				if len(buffer) < 6 {
					buffer = append(buffer, token)
				}
				w.Write(lemi025.AcquireResponse{})
			case setTime:
				if len(buffer) < 6 {
					buffer = append(buffer, token)
				}
			case setCoefficients1:

			case readCoefficients1:
			case setCoefficients2:
			case readCoefficients2:
			case readGPSData:
			case stopSystem:
			case startSystem:
			case checkFLASH:
				0
			case setDACx:
			case setDACy:
			case setDACz:
			}
		}
		w.Write(event)
	}
}
