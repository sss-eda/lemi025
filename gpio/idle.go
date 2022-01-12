package gpio

type idleState struct{}

func (state idleState) next(
	token byte,
) state {
	switch token {
	case 0x3F:
		return responseState{}
	case 0x4c:
		return dataState{
			buffer: []byte{token},
		}
	default:
		return idleState{}
	}
}
