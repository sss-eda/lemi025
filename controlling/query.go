package controlling

import (
	"github.com/sss-eda/lemi025/controlling/application/readconfig"
	"github.com/sss-eda/lemi025/controlling/application/readtime"
	"github.com/sss-eda/lemi025/controlling/application/settime"
)

// HandlerFuncType TODO
type HandlerFuncType interface {
	readconfig.HandlerFunc | readtime.HandlerFunc | settime.HandlerFunc
}

// Query - This stage serves no purpose
func Query(
	readConfig readconfig.HandlerFunc,
	readTime readtime.HandlerFunc,
	setTime settime.HandlerFunc,
) {
	switch HandlerFunc.(type) {
	case readconfig.HandlerFunc:
		return readConfig
	case readtime.HandlerFunc:
		return readTime
	}
}
