package controlling

// Queries TODO
// type Queries interface {
// 	readconfig.Query | readtime.Query | settime.Query
// }

// // HandlerFuncType TODO
// type HandlerFuncType interface {
// 	readconfig.HandlerFunc | readtime.HandlerFunc | settime.HandlerFunc
// }

// // HandlerFunc TODO
// type HandlerFunc func(func(*Response) error, *Request)

// // Response TODO
// type Response struct {
// 	Error error
// }

// // Request TODO
// type Request struct {
// 	Type    string
// 	Message []byte
// }

// // Queries TODO
// func Queries(
// 	readConfig readconfig.HandlerFunc,
// 	readTime readtime.HandlerFunc,
// 	setTime settime.HandlerFunc,
// ) HandlerFunc {
// 	return func(respond func(*Response) error, request *Request) {
// 		switch request.Type {
// 		case "Read Config":
// 			msg := readconfig.Request{}
// 			err := json.Unmarshal(request.Message, &msg)
// 			if err != nil {
// 				respond(&Response{Error: err})
// 				return
// 			}

// 			err = readConfig(
// 				func(response *readconfig.Response) error { return nil },
// 				&msg,
// 			)
// 			err = respond(&Response{
// 				Error: err,
// 			})
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		case "Read Time":
// 			return readTime
// 		}
// 	}
// }
