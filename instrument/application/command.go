package readconfig

// // Command TODO
// type Command struct {
// 	InstrumentID  instrument.ID
// 	StationType   string
// 	StationNumber instrument.StationNumber
// }

// // Mutate TODO
// func Mutate(
// 	instruments instrument.Repository,
// ) func(Command) error {
// 	return func(command Command) error {
// 		instrument, err := instruments.Load(command.InstrumentID)
// 		if err != nil {
// 			return err
// 		}
// 		event, err := instrument.Config.Read(
// 			command.StationType,
// 			command.StationNumber,
// 		)
// 		err = instruments.Save(command.InstrumentID, event)
// 		if err != nil {
// 			return err
// 		}
// 		// log.Println(event)
// 		return nil
// 	}
// }
