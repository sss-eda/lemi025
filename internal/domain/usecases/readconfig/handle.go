package readconfig

import "context"

// // Handler TOD
// type Handler interface {
// 	Handle(context.Context, Request, ResponseHandler) error
// }

// // HandlerFunc TODO
// type HandlerFunc func(context.Context, Request, ResponseHandler) error

// // ResponseHandler TODO
// type ResponseHandler func(context.Context, Response) error

// UseCase - This is where all of the business logic needs to go for this
// usecase!
func UseCase(
	gateway Gateway,
) func(context.Context, Request) (Response, error) {
	return func(ctx context.Context, request Request) (Response, error) {
		err := gateway()
		if err != nil {
			return Response{}, err
		}

		return Response{}, nil
	}
}
