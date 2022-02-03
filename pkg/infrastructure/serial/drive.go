package serial

import (
	"bufio"
	"io"

	"github.com/sss-eda/lemi025"
)

// Listen TODO
func Listen(
	r io.Reader,
	strategy lemi025.ListeningStrategy,
) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)

	var fsm state = idleState{}

	for scanner.Scan() {
		b := scanner.Bytes()[0]
		fsm.next(b)
	}

	// strategy.OnConfigRead(lemi025.OnConfigReadInput{})
}

type state interface {
	next(byte) state
}

type idleState struct{}

func (s idleState) next(
	token byte,
) state {
	switch token {
	case 0x3F:
		return responseState{}
	case 0x4C:
	}

	return idleState{}
}

type responseState struct{}

func (s responseState) next(
	token byte,
) state {
	var nextState state

	switch token {
	case 0x30:
		nextState = readConfigState{}
	case 0x31:
		nextState = readTimeState{}
	case 0x32:
		nextState = setTimeState{}
	case 0x33:
		nextState = setCoefficients1State{}
	case 0x34:
		nextState = readCoefficients1State{}
	case 0x35:
		nextState = setCoefficients2State{}
	case 0x36:
		nextState = readCoefficients2State{}
	case 0x37:
		nextState = readGPSDataState{}
	case 0x38:
		nextState = stopSystemState{}
	case 0x39:
		nextState = startSystemState{}
	case 0x3A:
		nextState = checkFLASHState{}
	case 0x3D:
		nextState = setDACxState{}
	case 0x3E:
		nextState = setDACyState{}
	case 0x3F:
		nextState = setDACzState{}
	default:
		nextState = idleState{}
	}

	return nextState
}

type readConfigState struct {
	buffer []byte
}

func (s readConfigState) next(
	token byte,
) state {
	var newState state

	switch len(s.buffer) {
	case 0:
		if token == 0x30 {
			s.buffer = append(s.buffer, token)
			newState = s
		} else {
			newState = idleState{}.next(token)
		}
	case 1:
		if token == 0x32 {
			s.buffer = append(s.buffer, token)
			newState = s
		} else {
			newState = idleState{}.next(token)
		}
	case 2:
		if token == 0x35 {
			s.buffer = append(s.buffer, token)
			newState = s
		} else {
			newState = idleState{}.next(token)
		}
	case 3:
		if token == 0x20 {
			s.buffer = append(s.buffer, token)
			newState = s
		} else {
			newState = idleState{}.next(token)
		}
	case 4:
		if token == 0x20 {
			s.buffer = append(s.buffer, token)
			newState = s
		} else {
			newState = idleState{}.next(token)
		}
	default:
		newState = idleState{}.next(token)
	}

	return newState
}

type readTimeState struct{}

func (s readTimeState) next(
	token byte,
) state {
	return idleState{}
}

type setTimeState struct{}

func (s setTimeState) next(
	token byte,
) state {
	return idleState{}
}

type setCoefficients1State struct{}

func (s setCoefficients1State) next(
	token byte,
) state {
	return idleState{}
}

type readCoefficients1State struct{}

func (s readCoefficients1State) next(
	token byte,
) state {
	return idleState{}
}

type setCoefficients2State struct{}

func (s setCoefficients2State) next(
	token byte,
) state {
	return idleState{}
}

type readCoefficients2State struct{}

func (s readCoefficients2State) next(
	token byte,
) state {
	return idleState{}
}

type readGPSDataState struct{}

func (s readGPSDataState) next(
	token byte,
) state {
	return idleState{}
}

type stopSystemState struct{}

func (s stopSystemState) next(
	token byte,
) state {
	return idleState{}
}

type startSystemState struct{}

func (s startSystemState) next(
	token byte,
) state {
	return idleState{}
}

type checkFLASHState struct{}

func (s checkFLASHState) next(
	token byte,
) state {
	return idleState{}
}

type setDACxState struct{}

func (s setDACxState) next(
	token byte,
) state {
	return idleState{}
}

type setDACyState struct{}

func (s setDACyState) next(
	token byte,
) state {
	return idleState{}
}

type setDACzState struct{}

func (s setDACzState) next(
	token byte,
) state {
	return idleState{}
}
