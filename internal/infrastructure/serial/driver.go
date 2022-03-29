package serial

// // DriverAdapter TODO
// type DriverAdapter struct {
// 	listen     func(context.Context, io.Reader) error
// 	readConfig func(context.Context, *lemi025.ReadConfigCommandPayload) error
// 	readTime   func(context.Context, *lemi025.ReadTimeCommandPayload) error
// 	setTime    func(context.Context, *lemi025.SetTimeCommandPayload) error
// }

// // NewDriverAdapter TODO
// func NewDriverAdapter(
// 	onConfigRead func(context.Context, *lemi025.ConfigReadEventPayload) error,
// 	onTimeRead func(context.Context, *lemi025.TimeReadEventPayload) error,
// 	onTimeSet func(context.Context, *lemi025.TimeSetEventPayload) error,
// ) (*DriverAdapter, error) {
// 	adapter := &DriverAdapter{
// 		Listen(
// 			onConfigRead,
// 			onTimeRead,
// 			onTimeSet,
// 		),
// 		ReadConfig(port),
// 		ReadTime(port),
// 		SetTime(port),
// 	}

// 	go adapter.listener(ctx, port)

// 	return adapter, nil
// }

// // ReadConfig TODO
// func (adapter *DriverAdapter) ReadConfig(
// 	ctx context.Context,
// 	payload *lemi025.ReadConfigCommandPayload,
// ) error {
// 	return adapter.readConfig(ctx, payload)
// }

// // ReadTime TODO
// func (adapter *DriverAdapter) ReadTime(
// 	ctx context.Context,
// 	payload *lemi025.ReadTimeCommandPayload,
// ) error {
// 	return adapter.readTime(ctx, payload)
// }
