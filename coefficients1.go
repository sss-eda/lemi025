package lemi025

import "fmt"

// Coefficients1 TODO
type Coefficients1 struct {
	mode  Mode
	uin   Uin
	mode1 Mode1
}

// NewCoefficients1 TODO
func NewCoefficients1(
	mode Mode,
	uin Uin,
	mode1 Mode1,
) (Coefficients1, error) {
	// Since all of the parameters are always valid, the result of this function
	// will also be always valid.
	return Coefficients1{mode, uin, mode1}, nil
}

// MarshalText TODO
func (coefficients1 Coefficients1) MarshalText() ([]byte, error) {
	mode, err := coefficients1.mode.MarshalText()
	if err != nil {
		return nil, err
	}
	uin, err := coefficients1.uin.MarshalText()
	if err != nil {
		return nil, err
	}
	mode1, err := coefficients1.mode1.MarshalText()
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("Coefficients1:\n  %s  %s  %s", mode, uin, mode1)), nil
}

// UnmarshalText TODO
func (coefficients1 Coefficients1) UnmarshalText(text []byte) error {
	return nil
}
