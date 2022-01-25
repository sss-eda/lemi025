package application

// Client TODO
type Client interface {
	ReadConfig() error
	ReadTime() error
	SetTime() error
}

// Server TODO
type Server interface {
	Serve() error
}
