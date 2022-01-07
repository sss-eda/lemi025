package gpio

// SetTimeCommand TODO
// type SetTimeCommand struct {
// 	writer io.Writer
// 	time   *instrumentation.Time
// 	// device *lemi025.Device // Don't use interface here, then the domain will have an obligation to fullfill this application level interface.
// }

// // Execute TODO
// func (command *SetTimeCommand) Execute() error {
// 	// year := command.time.Year()

// 	// Write the command to the command's writer interface.
// 	_, err := command.writer.Write([]byte{0x3D, 0x32})
// 	if err != nil {
// 		return err
// 	}

// 	return command.time.Set()
// }
