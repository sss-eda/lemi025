package cli

import "os"

type Application struct {
	args []string
}

func NewApp() *Application {
	app := Application{
		args: os.Args,
	}

	return &app
}

func (app *Application) Run() error {
	switch app.args {
	case "acquire":
		acquireCommand.Parse(app.args[2:])
	}

	return nil
}
