package getconfig

// This usecase is is a Query and thus only returns values stored in a repo.
// The projector that puts the data into this repo will also serve other
// usecases, so it needs to be in a different place?

// The same read model can be used for listing instruments, searching for
// instruments, etc.

// // UseCase TODO
// func UseCase(
// 	repository Repository,
// ) func(context.Context, Request) (Response, error) {
// 	return func(ctx context.Context, request Request) (Response, error) {
// 		instr, err := repository.GetInstrumentByID(request.InstrumentID)
// 		if err != nil {
// 			return Response{}, err
// 		}

// 		return Response{
// 			Config: instr.Config,
// 		}, nil
// 	}
// }
