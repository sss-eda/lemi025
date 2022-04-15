package serial

const (
	idleToken              byte = 0x00
	acquireToken           byte = 0x4C
	sendToken              byte = 0x3D
	receiveToken           byte = 0x3F
	readConfigToken        byte = 0x30
	readTimeToken          byte = 0x31
	setTimeToken           byte = 0x32
	setCoefficients1Token  byte = 0x33
	readCoefficients1Token byte = 0x34
)
