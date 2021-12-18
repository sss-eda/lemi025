package lemi025

type Config struct {
	StationNumber uint8
}

type ConfigReader interface {
	ReadConfig() error
}

func (config *Config) Read(
	adapter ConfigReader,
	command ReadConfigCommand,
) error {
	if 

	err := adapter.ReadConfig()
}
