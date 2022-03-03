package controlling

// Instrument TODO
type Instrument struct{}

// SendCommand TODO
func ExecuteCommand[Command Commands](
	execute CommandHandlerFunc[Command],
) CommandHandlerFunc[Command] {
	// This now is kind of like the read model. But just acts as a proxy in
	// this case. Note sure - what are the business rules that need to be
	// adhered to before a command is executed? There aren't any, really.
	return execute
}

// RaiseEvent TODO
func RaiseEvent[Event Events](
	raise EventHandlerFunc[Event],
) EventHandlerFunc[Event] {
	return raise
}
