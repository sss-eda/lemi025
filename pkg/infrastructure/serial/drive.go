package serial

import "io"

// Drive TODO
func Drive(
	r io.Reader,
	strategy lemi025.DriveClientStrategy,
) {
	strategy.OnConfigRead(lemi025.OnConfigReadInput{})
}
