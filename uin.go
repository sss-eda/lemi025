package lemi025

import (
	"fmt"
	"strconv"
)

// Uin TODO
type Uin struct {
	uinx10 uint8
}

// MarshalText TODO
func (uin Uin) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", uin.uinx10/10)), nil
}

// UnmarshalText TODO
func (uin *Uin) UnmarshalText(text []byte) error {
	uinx10, err := strconv.ParseUint(string(text), 10, 8)
	if err != nil {
		return err
	}

	uin.uinx10 = uint8(uinx10)

	return nil
}
