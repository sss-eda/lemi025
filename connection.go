package lemi025

type Connection struct{}

func (conn *Connection) ReadConfig(
	configReader ConfigReader,
) error {
	return configReader()
}
