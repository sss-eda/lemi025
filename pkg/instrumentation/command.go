package instrumentation

// // CommandType TODO
// type CommandType int

// const (
// 	// ReadConfig TODO
// 	ReadConfig CommandType = iota
// 	// ReadTime TODO
// 	ReadTime
// 	// SetTime TODO
// 	SetTime
// )

// // String TODO
// func (ct CommandType) String() string {
// 	switch ct {
// 	case ReadConfig:
// 		return "LEMI025 -> Driving -> Read Config"
// 	case ReadTime:
// 		return "LEMI025 -> Driving -> Read Time"
// 	case SetTime:
// 		return "LEMI025 -> Driving -> Set Time"
// 	}

// 	return ""
// }

// // Command TODO
// type Command interface {
// 	Type() CommandType
// 	Execute() error
// 	Expose(Server) error
// }

// // Server TODO
// type Server interface {
// 	Serve(Command) error
// }

// // ReadTimeCommand TODO
// type ReadTimeCommand struct {
// 	w io.Writer
// }

// // Type TODO
// func (command *ReadTimeCommand) Type() CommandType {
// 	return ReadTime
// }

// // Execute TODO
// func (command *ReadTimeCommand) Execute() error {
// 	return nil
// }

// // Expose TODO
// func (command *ReadTimeCommand) Expose(
// 	server Server,
// ) error {
// 	return server.Serve(command)
// }
