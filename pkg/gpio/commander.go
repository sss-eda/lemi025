package gpio

// Commander TODO - analog for the TV in example. This is the concrete receiver
type Commander struct {
	*Connection
}

// NewCommander TODO
func NewCommander(
	connection *Connection,
) (*Commander, error) {
	return &Commander{
		connection,
	}, nil
}

// ReadConfig TODO
func (commander *Commander) ReadConfig() error {
	commander.Lock()
	defer commander.Unlock()

	_, err := commander.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}

// ReadTime TODO
func (commander *Commander) ReadTime() error {
	_, err := commander.Write([]byte{0x3D, 0x31})
	if err != nil {
		return err
	}

	return nil
}

// SetTime TODO
func (commander *Commander) SetTime(
	year uint16,
	month uint8,
	day uint8,
	hour uint8,
	minute uint8,
	second uint8,
) error {
	_, err := commander.Write([]byte{0x3D, 0x32})
	if err != nil {
		return err
	}

	return nil
}
