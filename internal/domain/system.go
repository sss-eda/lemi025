package domain

// System TODO
type System bool

// StopSystemInput TODO
type StopSystemInput struct{}

// Stop TODO
func (system *System) Stop(
	input StopSystemInput,
) error {
	return nil
}

// StartSystemInput TODO
type StartSystemInput struct{}

// Start TODO
func (system *System) Start(
	input StartSystemInput,
) error {
	return nil
}
