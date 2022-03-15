package serial

import (
	"encoding/binary"
	"io"
	"math"

	"github.com/sss-eda/lemi025/instrument"
)

// SetCoefficients2 TODO
func SetCoefficients2(
	writer io.Writer,
) func(instrument.Coefficients2) error {
	return func(coefficients2 instrument.Coefficients2) error {
		buffer := make([]byte, 84)

		buffer[0] = 0x3D // Token "=" indicating send command
		buffer[1] = 0x35 // Token "5" indicating set coefficients2
		buffer[2] = 0xFF // This byte doesn't matter
		buffer[3] = 0xFF // This byte doesn't matter

		// Encode and append Ax1
		binary.LittleEndian.PutUint32(
			buffer[4:8],
			math.Float32bits(coefficients2.Ax1),
		)

		// Encode and append Ay1
		binary.LittleEndian.PutUint32(
			buffer[8:12],
			math.Float32bits(coefficients2.Ay1),
		)

		// Encode and append Az1
		binary.LittleEndian.PutUint32(
			buffer[12:16],
			math.Float32bits(coefficients2.Az1),
		)

		// Encode and append Beta
		binary.LittleEndian.PutUint32(
			buffer[16:20],
			math.Float32bits(coefficients2.Beta),
		)

		// Encode and append Gamma
		binary.LittleEndian.PutUint32(
			buffer[20:24],
			math.Float32bits(coefficients2.Gamma),
		)

		// Encode and append Xi
		binary.LittleEndian.PutUint32(
			buffer[24:28],
			math.Float32bits(coefficients2.Xi),
		)

		// Encode and append Exy
		binary.LittleEndian.PutUint32(
			buffer[28:32],
			math.Float32bits(coefficients2.Exy),
		)

		// Encode and append Eyz
		binary.LittleEndian.PutUint32(
			buffer[32:36],
			math.Float32bits(coefficients2.Eyz),
		)

		// Encode and append Exz
		binary.LittleEndian.PutUint32(
			buffer[36:40],
			math.Float32bits(coefficients2.Exz),
		)

		// Encode and append K1x
		binary.LittleEndian.PutUint32(
			buffer[40:44],
			math.Float32bits(coefficients2.K1x),
		)

		// Encode and append K1y
		binary.LittleEndian.PutUint32(
			buffer[44:48],
			math.Float32bits(coefficients2.K1y),
		)

		// Encode and append K1z
		binary.LittleEndian.PutUint32(
			buffer[48:52],
			math.Float32bits(coefficients2.K1z),
		)

		// Encode and append K2x
		binary.LittleEndian.PutUint32(
			buffer[52:56],
			math.Float32bits(coefficients2.K2x),
		)

		// Encode and append K2y
		binary.LittleEndian.PutUint32(
			buffer[56:60],
			math.Float32bits(coefficients2.K2y),
		)

		// Encode and append K2z
		binary.LittleEndian.PutUint32(
			buffer[60:64],
			math.Float32bits(coefficients2.K2z),
		)

		// Encode and append KTF
		binary.LittleEndian.PutUint32(
			buffer[64:68],
			math.Float32bits(coefficients2.KTF),
		)

		// Encode and append KTE
		binary.LittleEndian.PutUint32(
			buffer[68:72],
			math.Float32bits(coefficients2.KTE),
		)

		// Encode and append KTF0
		binary.LittleEndian.PutUint32(
			buffer[72:76],
			math.Float32bits(coefficients2.KTF0),
		)

		// Encode and append KTE0
		binary.LittleEndian.PutUint32(
			buffer[76:80],
			math.Float32bits(coefficients2.KTE0),
		)

		// Encode and append KVBAT
		binary.LittleEndian.PutUint32(
			buffer[80:84],
			math.Float32bits(coefficients2.KVBAT),
		)

		_, err := writer.Write(buffer)
		return err
	}
}
