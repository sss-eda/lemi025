package lemi025

// ConfigReader TODO
type ConfigReader func() error

// ReadConfig TODO
func ReadConfig(
	configReader ConfigReader,
) error {
	return configReader()
}
