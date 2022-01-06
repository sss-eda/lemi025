package lemi025

// Command TODO
type Command interface {
	Execute() error
}

// Receiver: The connection to the lemi

// Command: The command (must embed the receiver)

// Invoker: Can be the CLI, web api, ...

// Client: It creates the command with the appropriate receiver bypassing the
// receiver to the command’s constructor. After that, it also associates the
// resulting command with an invoker.
