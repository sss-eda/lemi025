package eventsourcing

type command int

const (
	readConfig command = iota
	readTime
	setTime
)

// API TODO
type API struct {
	store    Store
	commands map[command]func() error
}

// NewAPI TODO
func NewAPI(
	store Store,
) error {
	api := &API{
		store: store,
		commands: {
			readConfig: ReadConfig(store),
			readTime: ReadTime(store),
		}
	}
	
	return nil
}

func (api *API) ReadConfig(
	input ReadConfigInput,
) error {
	return api.ReadConfig(input)
}
