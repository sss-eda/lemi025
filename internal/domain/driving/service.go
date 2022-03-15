package driving

// Service TODO
type Service interface {
	ReadConfig() error
	ReadTime() error
	SetTime() error
	SetCoefficients1() error
	ReadCoefficients1() error
	SetCoefficients2() error
	ReadCoefficients2() error
}
