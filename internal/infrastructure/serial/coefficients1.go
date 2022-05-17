package serial

// Coefficients1 TODO
type Coefficients1 struct {
	Mode  uint8
	Uin   uint8
	Mode1 uint8
}

// func (coefficients1 *Coefficients1) UnmarshalBinary(data []byte) error {
// 	coefficients1 = &Coefficients1{}
// 	var err error

// 	if len(data) != 2 && len(data) != 4 {
// 		coefficients1 = nil
// 		return fmt.Errorf("unexpected data length: %v", len(data))
// 	}

// 	if data[0] != 0x00 {
// 		coefficients1 = nil
// 		return fmt.Errorf("unexpected value for byte 0: %v", data[0])
// 	}

// 	coefficients1.Mode = data[1]
// 	if err != nil {
// 		coefficients1 = nil
// 		return fmt.Errorf("failed to decode mode: %w", err)
// 	}

// 	coefficients1.Uin, err = bcd.Decode(data[1])
// 	if err != nil {
// 		coefficients1 = nil
// 		return fmt.Errorf("failed to decode uin: %w", err)
// 	}

// 	coefficients1.Mode1, err = bcd.Decode(data[2])
// 	if err != nil {
// 		coefficients1 = nil
// 		return fmt.Errorf("failed to decode mode1: %w", err)
// 	}

// 	return nil
// }
