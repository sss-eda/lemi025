package driving

// Device TODO
type Device interface {
	ReadConfig() error
	ReadTime() error
	SetTime(Time) error
}
