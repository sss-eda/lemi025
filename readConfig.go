package lemi025

// Command TODO
type Command func() error

// ReadConfig TODO
func ReadConfig(
	configReader ConfigReader,
) Command {
	return configReader.ReadConfig
}

// ConfigReader TODO
type ConfigReader interface {
	ReadConfig() error
}

// // ReadConfigCommand TODO
// type ReadConfigCommand struct {
// 	ConfigReader `json:"-"`
// }

// // Execute TODO
// func (command *ReadConfigCommand) Execute() error {
// 	return command.ReadConfig()
// }
