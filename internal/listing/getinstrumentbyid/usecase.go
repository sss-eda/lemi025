package getinstrumentbyid

import "github.com/sss-eda/lemi025/internal/listing"

// Request TODO
type Request struct {
	InstrumentID listing.InstrumentID
}

// Response TODO
type Response struct {
	Instrument listing.Instrument
}

// GetInstrumentByID TODO
func UseCase(
	repo listing.Repository,
) func(Request) (Response, error) {
	return func(request Request) (Response, error) {
		instrument, err := repo.GetInstrumentByID(request.InstrumentID)
		if err != nil {
			return Response{}, err
		}

		return Response{
			Instrument: instrument,
		}, nil
	}
}
