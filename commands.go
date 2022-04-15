package lemi025

// CommandPayloads TODO
type CommandPayloads interface {
	ReadConfigCommandPayload |
		ReadTimeCommandPayload |
		SetTimeCommandPayload |
		SetCoefficients1CommandPayload |
		ReadCoefficients1CommandPayload
}

// ReadConfigCommandPayload TODO
type ReadConfigCommandPayload struct{}

// ReadTimeCommandPayload TODO
type ReadTimeCommandPayload struct{}

// SetTimeCommandPayload TODO
type SetTimeCommandPayload struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// SetCoefficients1CommandPayload TODO
type SetCoefficients1CommandPayload struct {
	Mode Mode
}

// ReadCoefficients1CommandPayload TODO
type ReadCoefficients1CommandPayload struct {
	Mode  Mode
	Uin   Uin
	Mode1 Mode1
}
