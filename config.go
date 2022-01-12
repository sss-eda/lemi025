package lemi025

import "sync"

// Config TODO
type Config struct {
	sync.Mutex
	syncStatus    bool
	StationNumber uint8
	events        []Event
}

func (config *Config) OnConfigRead(
	event *ConfigReadEvent,
) {
	config.Lock()
	defer config.Unlock()

	config.StationNumber = event.StationNumber
	config.syncStatus = true
}

// ConfigReader TODO
type ConfigReader interface {
	ReadConfig() error
}

// Read TODO
func (config *Config) Read(
	reader ConfigReader,
) error {
	return reader.ReadConfig()
}

// Sync TODO
func (config *Config) Sync() error {
	config.syncStatus = true

	return nil
}

// Desync TODO
func (config *Config) Desync() error {
	config.syncStatus = false

	return nil
}
