package controlling

import (
	"github.com/sss-eda/lemi025/controlling/application/readconfig"
	"github.com/sss-eda/lemi025/controlling/application/readtime"
	"github.com/sss-eda/lemi025/controlling/application/settime"
)

// Commands TODO
type Commands interface {
	readconfig.Command | readtime.Command | settime.Command
}

type CommandHandlerFunc[Command Commands] func(*Command) error
