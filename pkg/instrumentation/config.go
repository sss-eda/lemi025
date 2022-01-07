package instrumentation

// Config TODO
type Config struct {
	reader        Command
	sync          bool
	StationNumber uint8 `json:"station-number"`
}

// NewConfig TODO
func NewConfig(
	reader Command,
) *Config {
	return &Config{
		reader:        reader,
		sync:          false,
		StationNumber: 0,
	}
}

func (config *Config) Read() error {
	err := config.reader.Execute()
	if err != nil {
		return err
	}

	config.sync = false

	return nil
}
