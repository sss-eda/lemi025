package iaga

import "encoding"

// Encoding TODO
type Encoding interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// File TODO
type File struct {
	Encoding
}

// Write TODO
func (file *File) Write(p []byte) (n int, err error) {
	datapoint := listing.DataPoint{}
	err = datapoint.Unmarshal(&p)
	if err != nil {
		return
	}

	file.Write()

	return
}

// Read TODO
func (file *File) Read(p []byte) (n, int, err error) {
	return
}
