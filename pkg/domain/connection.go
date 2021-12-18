package lemi025

// type Transmitter interface {
// 	ReadConfig() error
// 	ReadTime() error
// 	SetTime(Time) error
// }

// type Receiver interface {
// 	AcquireData(chan<- Datum) error
// }

// // Connection TODO
// type Connection struct {
// 	tx Transmitter
// 	rx Receiver
// }

// func Connect(
// 	events chan<- Event,
// 	commands <-chan Command,
// ) (*Connection, error) {
// 	connection := Connection{
// 		tx: transmitter,
// 		rx: receiver,
// 	}

// 	return &connection, nil
// }

// func (c *Connection) ReadConfig() error {
// 	c.tx.ReadConfig()
// }
