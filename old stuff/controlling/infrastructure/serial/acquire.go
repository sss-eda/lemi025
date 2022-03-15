package serial

import (
	"bufio"
	"context"
	"io"
	"log"
)

type state int

const (
	idle state = iota
	data
	response
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

// Acquire TODO
func Acquire(
	ctx context.Context,
	reader io.Reader,
	onConfigRead func([]byte) error,
	onTimeRead func([]byte) error,
	onTimeSet func([]byte) error,
	onDataFrameAcquired func([]byte) error,
) error { // Where request specify which event to respond to
	byteStream := make(chan byte, 1024)
	go func(ctx context.Context, r io.Reader) {
		defer close(byteStream)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			for _, b := range scanner.Bytes() {
				select {
				case byteStream <- b:
				case <-ctx.Done():
					return
				}
			}
		}
		return
	}(ctx, reader)

	next := idle
	buffer := []byte{}
	for {
		select {
		case b := <-byteStream:
			switch next {
			case idle:
				if len(buffer) > 0 {
					log.Printf("byte dropped: %v", buffer[0])
					buffer = buffer[1:]
				}
				buffer = append(buffer, b)
				for _, b := range buffer {
					switch b {
					case 0x3F:
						next = response
						break
					case 0x4C:
						next = data
						break
					default:
						log.Printf("byte dropped: %v", buffer[0])
						buffer = buffer[1:]
					}
				}
			case data:
				if len(buffer) < 152 {
					buffer = append(buffer, b)
				}
				err := onDataFrameAcquired(buffer)
				if err != nil {
					log.Println(err)
					next = idle
				}
			case response:
				buffer = append(buffer, b)
				switch b {
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
					next = idle
				}
			case readConfig:
				if len(buffer) < 5 {
					buffer = append(buffer, b)
				}
				err := onConfigRead(buffer)
				if err != nil {
					log.Println(err)
					next = idle
				}
			case readTime:
				if len(buffer) < 8 {
					buffer = append(buffer, b)
				}
				err := onTimeRead(buffer)
				if err != nil {
					log.Println(err)
					next = idle
				}
			case setTime:
				if len(buffer) < 8 {
					buffer = append(buffer, b)
				}
				err := onTimeSet(buffer)
				if err != nil {
					log.Println(err)
					next = idle
				}
			case setCoefficients1:
				if len(buffer) < 8 {
					buffer = append(buffer, b)
				}
				err := onTimeSet(buffer)
				if err != nil {
					log.Println(err)
					next = idle
				}
			case readCoefficients1:
			case setCoefficients2:
			case readCoefficients2:
			case readGPSData:
			case stopSystem:
			case startSystem:
			case checkFLASH:
			case setDACx:
			case setDACy:
			case setDACz:
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
