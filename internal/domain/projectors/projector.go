package projectors

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// ListingProjector TODO
type ListingProjector struct {
	repository ListingRepository
}

// Run TODO
func (projector *ListingProjector) Project(
	events <-chan instrument.Event,
) error {
	for event := range events {

	}
}

// Or we can define a func on each event called Project(repository)
