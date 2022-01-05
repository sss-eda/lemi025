package lemi025

// Controller TODO
// type Controller interface {
// 	Commander
// 	Observer
// }

// Commander TODO - Actions
type Commander interface {
	ReadConfig() error
	ReadTime() error
	SetTime(
		year uint16,
		month uint8,
		day uint8,
		hour uint8,
		minute uint8,
		second uint8,
	) error
}

// Driver TODO - Actor
type Driver interface {
	Drive(Commander)
}

// Projector TODO - Actions
type Projector interface {
	OnConfigRead(ConfigReadEvent)
	OnTimeRead(TimeReadEvent)
	OnTimeSet(TimeSetEvent)
}

// Projectionist TODO - Actor
type Projectionist interface {
	Project(Projector)
}
