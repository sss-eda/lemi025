package text

// // Projector TODO
// type Projector interface {
// 	Project(context.Context, <-chan Event)
// }

// type projector struct{}

// // Project TODO
// func (prj *projector) Project(
// 	ctx context.Context,
// 	events <-chan Event,
// ) error {
// 	for {
// 		select {
// 		case event := <-events:
// 			switch event.Type {
// 			case acquiring.NewDataPointAcquired:
// 				e := event.Payload.Deserialize()
// 			}
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		}
// 	}
// }

// // Events TODO
// type Events map[EventType]Event
