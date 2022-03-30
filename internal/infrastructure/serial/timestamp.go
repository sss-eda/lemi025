package serial

import (
	"fmt"
	"time"

	"github.com/sss-eda/encoding/bcd"
)

// Timestamp TODO
type Timestamp struct {
	year   uint8
	month  uint8
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

// MarshalBinary TODO
func (timestamp *Timestamp) MarshalBinary() ([]byte, error) {
	_, err := time.Parse(
		"2006-01-02 15:04:05",
		fmt.Sprintf(
			"20%02d-%02d-%02d %02d:%02d:%02d",
			timestamp.year,
			timestamp.month,
			timestamp.day,
			timestamp.hour,
			timestamp.minute,
			timestamp.second,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp: %w", err)
	}

	year, err := bcd.Encode(timestamp.year)
	if err != nil {
		return nil, fmt.Errorf("failed to encode year: %w", err)
	}

	month, err := bcd.Encode(timestamp.month)
	if err != nil {
		return nil, fmt.Errorf("failed to encode month: %w", err)
	}

	day, err := bcd.Encode(timestamp.day)
	if err != nil {
		return nil, fmt.Errorf("failed to encode day: %w", err)
	}

	hour, err := bcd.Encode(timestamp.hour)
	if err != nil {
		return nil, fmt.Errorf("failed to encode hour: %w", err)
	}

	minute, err := bcd.Encode(timestamp.minute)
	if err != nil {
		return nil, fmt.Errorf("failed to encode minute: %w", err)
	}

	second, err := bcd.Encode(timestamp.second)
	if err != nil {
		return nil, fmt.Errorf("failed to encode second: %w", err)
	}

	return []byte{year, month, day, hour, minute, second}, nil
}

// UnmarshalBinary TODO
func (timestamp *Timestamp) UnmarshalBinary(data []byte) error {
	timestamp = &Timestamp{}
	var err error

	if len(data) != 6 {
		return fmt.Errorf("unexpected data length: %v", len(data))
	}

	timestamp.year, err = bcd.Decode(data[0])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode year: %w", err)
	}

	timestamp.month, err = bcd.Decode(data[1])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode month: %w", err)
	}

	timestamp.day, err = bcd.Decode(data[2])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode day: %w", err)
	}

	timestamp.hour, err = bcd.Decode(data[3])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode hour: %w", err)
	}

	timestamp.minute, err = bcd.Decode(data[4])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode minute: %w", err)
	}

	timestamp.second, err = bcd.Decode(data[5])
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to decode second: %w", err)
	}

	_, err = time.Parse(
		"2006-01-02 15:04:05",
		fmt.Sprintf(
			"20%02d-%02d-%02d %02d:%02d:%02d",
			timestamp.year,
			timestamp.month,
			timestamp.day,
			timestamp.hour,
			timestamp.minute,
			timestamp.second,
		),
	)
	if err != nil {
		timestamp = nil
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	return nil
}

// type readTimeCommand lemi025.ReadTimeCommandPayload

// // MarshalBinary TODO
// func (command readTimeCommand) MarshalBinary() ([]byte, error) {
// 	return []byte{sendToken, readTimeToken}, nil
// }

// type setTimeCommand lemi025.SetTimeCommandPayload

// func (command setTimeCommand) MarshalBinary() ([]byte, error) {
// 	timestamp, err := (&Timestamp{
// 		year:   command.Year,
// 		month:  command.Month,
// 		day:    command.Day,
// 		hour:   command.Hour,
// 		minute: command.Minute,
// 		second: command.Second,
// 	}).MarshalBinary()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return append([]byte{sendToken, setTimeToken}, timestamp...), nil
// }

// type timeReadEvent lemi025.TimeReadEventPayload

// func (event *timeReadEvent) UnmarshalBinary(data []byte) error {
// 	timestamp := &Timestamp{}
// 	err := timestamp.UnmarshalBinary(data[1:])
// 	if err != nil {
// 		return err
// 	}

// 	event.Timestamp = timestamp

// 	return nil
// }
