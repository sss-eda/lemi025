package getting

// Service TODO
type Service interface {
	GetInstrumentByID(InstrumentID) (Instrument, error)
}

type service struct {
	repo Repository
}

func (svc *service) GetInstrumentByID(
	id InstrumentID,
)
