package readconfig

import "context"

// UseCase - This is where all of the business logic needs to go for this
// usecase!
func Query(
	gateway func() error,
) func(context.Context, Request) (Response, error) {
	return func(ctx context.Context, request Request) (Response, error) {
		err := gateway()
		if err != nil {
			return Response{}, err
		}

		return Response{}, nil
	}
}
