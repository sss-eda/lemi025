package serial

type mux struct {
	io.ReadWriteCloser
}

func (m *mux) HandleCommands (
	done <-chan interface{},
	commands <-chan lemi025.Command,
) <-chan lemi025.Event {
	eventStream := make(chan lemi025.Event)
	go func() {
		defer close(eventStream)
		var err error

		for command := range commands {
			switch command := command.(type) {
			case lemi025.ReadConfigCommand:
				err = m.ReadConfig(command)
			case lemi025.ReadTimeCommand:
				err = m.ReadTime(command)
			case lemi025.SetTimeCommand:
				err = m.SetTime()
			}

			select {
			case <-done:
				return
				case eventStream <- 
			}
		}
	}()

	return eventStream
}