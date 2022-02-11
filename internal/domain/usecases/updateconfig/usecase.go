package updateconfig

// UseCase TODO - In the case of commands there is not much business logic in
// this layer, mostly orchestration. However, this will also be the anti-
// corruption layer. So input needs to be validated here. Input validation is
// not busness logic.
func UseCase(
	repository Repository,
) func(Request) (Response, error) {
	return func(request Request) (Response, error) {
		instr, err := repository.Load(request.InstrumentID)
		if err != nil {
			return Response{}, err
		}

		err := instr.Config.Update(request.NewConfig)
		if err != nil {
			return Response{}, err
		}

		err = repository.Save(request.InstrumentID, instr)
		if err != nil {
			return Response{}, err
		}

		return Response{}, nil
	}
}
