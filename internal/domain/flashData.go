package domain

// FLASHData TODO
type FLASHData struct {
	Size Size
	Free Free
}

// Size TODO
type Size uint16

// Free TODO
type Free uint8

// CheckFLASHDataCommand TODO
type CheckFLASHDataCommand struct{}
