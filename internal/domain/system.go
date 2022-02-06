package domain

type System bool

type StopSystemCommand struct{}

func (system *System) Stop() {}

type StartSystemCommand struct{}

func (system *System) Start() {}
