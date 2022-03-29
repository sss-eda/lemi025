package nats

// // Control TODO
// func Control[Request encoding.BinaryUnmarshaler](
// 	ctx context.Context,
// 	nc *natsio.Conn,
// 	subject string,
// 	usecase func(Request) error,
// ) error {
// 	subscription, err := nc.Subscribe(subject, func(msg *natsio.Msg) {
// 		var request Request
// 		err := request.UnmarshalBinary(msg.Data)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		err = usecase(request)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer subscription.Unsubscribe()

// 	<-ctx.Done()
// 	return ctx.Err()
// }
