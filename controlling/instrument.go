package controlling

// Instrument TODO
type Instrument struct{}

// ReadConfig TODO
func (instrument *Instrument) ReadConfig(
	configReader func(*ReadConfigCommand) error,
) func(*ReadConfigCommand) error {
	return func(command *ReadConfigCommand) error {
		return configReader(command)
	}
}

// Query TODO
func (instrument *Instrument) Query() error {

}

// Subscribe TODO
func (instrument *Instrument) Subscribe() error {

}
