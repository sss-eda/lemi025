package getconfig

import "context"

// UseCase TODO
func UseCase(
	repository Repository,
) func(context.Context, Request) (Response, error) {
	return func(ctx context.Context, request Request) (Response, error) {
		instr, err := repository.GetInstrumentByID(request.InstrumentID)
		if err != nil {
			return Response{}, err
		}

		return Response{
			Config: instr.Config,
		}, nil
	}
}
