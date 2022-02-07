package application

// Application TODO
type Application struct {
	Mutations     Mutations
	Queries       Queries
	Subscriptions Subscriptions
}

// Mutations TODO
type Mutations struct {
	ReadConfig mutation.ReadConfigHandler
	ReadTime   mutation.ReadTimeHandler
	SetTime    mutation.SetTimeHandler
}

// Queries TODO
type Queries struct {
}

// Subscriptions TODO
type Subscriptions struct {
	Data       event.Data
	ConfigRead event.ConfigRead
	TimeRead   event.TimeRead
	TimeSet    event.TimeSet
}
