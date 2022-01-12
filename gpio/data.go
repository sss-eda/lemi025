package gpio

import (
	"io"
	"log"
	"strconv"
)

type dataState struct {
	buffer []byte
}

func (s *dataState) next(
	token byte,
) state {
	switch len(s.buffer) {
	case 0:
		log.Println("invalid state, returning to idle")

		return idleState{}
	case 1:
		log.Println("invalid token, returning to idle")

		if token != 0x30 {
			return idleState{}
		}
	case 2:
		log.Println("invalid token, returning to idle")

		if token != 0x32 {
			return idleState{}
		}
	case 3:
		log.Println("invalid token, returning to idle")

		if token != 0x35 {
			return idleState{}
		}
	case 4:
		log.Println("invalid token, returning to idle")

		if token != 0x32 {
			return idleState{}
		}
	case 152:
		event := &lemi025.DataPointAcquiredEvent{
			StationNumber: s.buffer[4],
			Time: &lemi025.Time{
				Year:   uint16(s.buffer[5]) + 2000,
				Month:  s.buffer[6],
				Day:    s.buffer[7],
				Hour:   s.buffer[8],
				Minute: s.buffer[9],
				Second: s.buffer[10],
			},
			Tf: int16(strconv.ParseInt(string(s.buffer[11:12]), 10, 16)),
		}

		return idleState{}
	}

	s.buffer = append(s.buffer, token)

	return s
}

func (fsm *FSM) idle() {

}

type Event interface{}

type Lemi025Reader interface {
	Read(io.Reader) (Event, error)
}

type Lemi025Writer interface {
	Write(io.Writer, Command) error
}

func NewLemi025Writer(w io.Writer)
