package serial

// // Listen TODO
// func Listen(
// 	reader io.Reader,
// 	instrument lemi025.Instrument,
// ) {
// 	scanner := bufio.NewScanner(r)
// 	scanner.Split(bufio.ScanBytes)

// 	for scanner.Scan() {
// 		b := scanner.Bytes()[0]

// 		switch b {
// 		case 0x4c:

// 		case 0x3F:
// 		}
// 	}
// }

// type state interface {
// 	next(byte) state
// }

// type baseState struct {
// 	lemi025.ListeningStrategy
// }

// type idleState struct {
// 	baseState
// }

// func (s idleState) next(
// 	token byte,
// ) state {
// 	switch token {
// 	case 0x3F:
// 		return responseState{s.baseState}
// 	case 0x4C:
// 	}

// 	return idleState{s.baseState}
// }

// type responseState struct {
// 	baseState
// }

// func (s responseState) next(
// 	token byte,
// ) state {
// 	var nextState state

// 	switch token {
// 	case 0x30:
// 		request := [5]byte{}
// 		n, err := reader.Read(request)
// 		nextState = readConfigState{s.baseState, []byte{}}
// 	case 0x31:
// 		nextState = readTimeState{s.baseState}
// 	case 0x32:
// 		nextState = setTimeState{s.baseState}
// 	case 0x33:
// 		nextState = setCoefficients1State{s.baseState}
// 	case 0x34:
// 		nextState = readCoefficients1State{s.baseState}
// 	case 0x35:
// 		nextState = setCoefficients2State{s.baseState}
// 	case 0x36:
// 		nextState = readCoefficients2State{s.baseState}
// 	case 0x37:
// 		nextState = readGPSDataState{s.baseState}
// 	case 0x38:
// 		nextState = stopSystemState{s.baseState}
// 	case 0x39:
// 		nextState = startSystemState{s.baseState}
// 	case 0x3A:
// 		nextState = checkFLASHState{s.baseState}
// 	case 0x3D:
// 		nextState = setDACxState{s.baseState}
// 	case 0x3E:
// 		nextState = setDACyState{s.baseState}
// 	case 0x3F:
// 		nextState = setDACzState{s.baseState}
// 	default:
// 		nextState = idleState{s.baseState}
// 	}

// 	return nextState
// }

// type readConfigState struct {
// 	baseState
// 	buffer []byte
// }

// func (s readConfigState) next(
// 	token byte,
// ) state {
// 	var newState state

// 	switch len(s.buffer) {
// 	case 0:
// 		if token == 0x30 {
// 			s.buffer = append(s.buffer, token)
// 			newState = s
// 		} else {
// 			newState = idleState{s.baseState}.next(token)
// 		}
// 	case 1:
// 		if token == 0x32 {
// 			s.buffer = append(s.buffer, token)
// 			newState = s
// 		} else {
// 			newState = idleState{s.baseState}.next(token)
// 		}
// 	case 2:
// 		if token == 0x35 {
// 			s.buffer = append(s.buffer, token)
// 			newState = s
// 		} else {
// 			newState = idleState{s.baseState}.next(token)
// 		}
// 	case 3:
// 		if token == 0x20 {
// 			s.buffer = append(s.buffer, token)
// 			newState = s
// 		} else {
// 			newState = idleState{s.baseState}.next(token)
// 		}
// 	case 4:
// 		sn, err := bcd.Decode(token)
// 		if err != nil {
// 			newState = idleState{s.baseState}.next(token)
// 		}
// 		s.baseState.OnConfigRead(lemi025.OnConfigReadInput{
// 			StationType:   string(s.buffer),
// 			StationNumber: uint8(sn),
// 		})
// 		newState = idleState{s.baseState}
// 	default:
// 		newState = idleState{s.baseState}.next(token)
// 	}

// 	return newState
// }

// // type readTimeState struct {
// // 	baseState
// // }

// // func (s readTimeState) next(
// // 	token byte,
// // ) state {
// // 	return idleState{}
// // }

// func onTimeRead(
// 	strategy interface{ OnTimeRead(lemi025.OnTimeReadInput) },
// 	buffer []byte,
// ) func(byte) state {
// 	return func(token byte) state {
// 		if len(buffer)
// 		switch len(buffer) {
// 		case 0:
// 			buffer = append(buffer, token)
// 		case 1:
// 		case 2:
// 		case 3:
// 		case 4:
// 		case 5:
// 		}
// 	}
// }

// type setTimeState struct {
// 	baseState
// }

// func (s setTimeState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type setCoefficients1State struct {
// 	baseState
// }

// func (s setCoefficients1State) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type readCoefficients1State struct {
// 	baseState
// }

// func (s readCoefficients1State) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type setCoefficients2State struct {
// 	baseState
// }

// func (s setCoefficients2State) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type readCoefficients2State struct {
// 	baseState
// }

// func (s readCoefficients2State) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type readGPSDataState struct {
// 	baseState
// }

// func (s readGPSDataState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type stopSystemState struct {
// 	baseState
// }

// func (s stopSystemState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type startSystemState struct {
// 	baseState
// }

// func (s startSystemState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type checkFLASHState struct {
// 	baseState
// }

// func (s checkFLASHState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type setDACxState struct {
// 	baseState
// }

// func (s setDACxState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type setDACyState struct {
// 	baseState
// }

// func (s setDACyState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }

// type setDACzState struct {
// 	baseState
// }

// func (s setDACzState) next(
// 	token byte,
// ) state {
// 	return idleState{}
// }
