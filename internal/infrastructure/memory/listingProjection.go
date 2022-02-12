package memory

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// ListingProjection TODO
type ListingProjection struct{}

// NewListingProjection TODO
func NewListingProjection() (*ListingProjection, error) {
	return &ListingProjection{}, nil
}

// Project TODO
func (projection *ListingProjection) Project(
	event Event,
) error {

}

func (projection *ListingProjection) OnInstrumentAdded() {

}

func (projection *ListingProjection) OnConfigRead(
	event instrument.ConfigReadEvent,
) {
	projection.repository.
}
