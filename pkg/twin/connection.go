package twin

// This is the root
// type Connection struct {
// 	commands chan<- interface{}
// 	events   <-chan interface{}
// }

// func Connect(
// 	commands chan<- interface{},
// 	events <-chan interface{},
// ) (*Connection, error) {
// 	connection := Connection{
// 		commands: commands,
// 		events:   events,
// 	}

// 	go connection.loop()

// 	return &connection, nil
// }

// func (connection *Connection) loop() error {
// 	for command := range connection.commands {
// 		switch c := command.(type) {
// 		case ReadConfigCommand:
// 			connection.readConfig(c)
// 			return nil
// 		case ReadTimeCommand:
// 			connection.readTime(c)
// 			return nil
// 		case SetTimeCommand:
// 			connection.setTime(c)
// 			return nil
// 		}

// 		log.Printf("unsupported command type: %v\n", command)
// 	}

// 	return fmt.Errorf("commands channel closed by host")
// }
