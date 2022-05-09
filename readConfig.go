package lemi025

// // ReadConfigCommand TODO
// type ReadConfigCommand struct{}

// // ReadConfigCommandHandler TODO
// type ReadConfigCommandHandler interface {
// 	Handle(context.Context, *ReadConfigCommand) error
// }

// // StationRepository TODO
// type StationRepository interface {
// 	GetStationByID(id uint16) (*Station, error)
// 	SaveStation(station *Station) error
// }

// // ReadConfigCommand TODO
// type readConfigCommandHandler struct {
// 	stations StationRepository
// }

// // Handle TODO
// func (handler *readConfigCommandHandler) Handle(ctx context.Context, command *ReadConfigCommand) error {
// 	station, err := handler.stations.GetStationByID(command.StationID)
// 	if err != nil {
// 		return err
// 	}

// 	station.ReadConfig()

// 	repository.SaveStation(station)
// }
