package serial

import (
	"encoding/binary"
	"io"
	"math"

	"github.com/sss-eda/lemi025/internal/domain"
)

// SetCoefficients2 TODO
func SetCoefficients2(
	w io.Writer,
) func(domain.SetCoefficients2Command) error {
	return func(command domain.SetCoefficients2Command) error {
		buffer := make([]byte, 84)

		buffer = []byte{0x3D, 0x35, 0xFF, 0xFF}
		binary.LittleEndian.PutUint32(buffer[4:8], math.Float32bits(command.Ax1))
		binary.LittleEndian.PutUint32(buffer[8:12], math.Float32bits(command.Ay1))
		binary.LittleEndian.PutUint32(buffer[12:16], math.Float32bits(command.Az1))
		binary.LittleEndian.PutUint32(buffer[16:20], math.Float32bits(command.Beta))
		binary.LittleEndian.PutUint32(buffer[20:24], math.Float32bits(command.Gamma))
		binary.LittleEndian.PutUint32(buffer[24:28], math.Float32bits(command.Xi))
		binary.LittleEndian.PutUint32(buffer[28:32], math.Float32bits(command.Exy))
		binary.LittleEndian.PutUint32(buffer[32:36], math.Float32bits(command.Eyz))
		binary.LittleEndian.PutUint32(buffer[36:40], math.Float32bits(command.Exz))
		binary.LittleEndian.PutUint32(buffer[40:44], math.Float32bits(command.K1x))
		binary.LittleEndian.PutUint32(buffer[44:48], math.Float32bits(command.K1y))
		binary.LittleEndian.PutUint32(buffer[48:52], math.Float32bits(command.K1z))
		binary.LittleEndian.PutUint32(buffer[52:56], math.Float32bits(command.K2x))
		binary.LittleEndian.PutUint32(buffer[56:60], math.Float32bits(command.K2y))
		binary.LittleEndian.PutUint32(buffer[60:64], math.Float32bits(command.K2z))
		binary.LittleEndian.PutUint32(buffer[64:68], math.Float32bits(command.KTF))
		binary.LittleEndian.PutUint32(buffer[68:72], math.Float32bits(command.KTE))
		binary.LittleEndian.PutUint32(buffer[72:76], math.Float32bits(command.KTF0))
		binary.LittleEndian.PutUint32(buffer[76:80], math.Float32bits(command.KTE0))
		binary.LittleEndian.PutUint32(buffer[80:84], math.Float32bits(command.KVBAT))

		_, err := w.Write(buffer)
		if err != nil {
			return err
		}

		return nil
	}
}
