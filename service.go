package lemi025

import "github.com/sss-eda/lemi025/pkg/driving"

// Driver TODO
type Driver interface {
	Drive(driving.Drive) error
}

type Client interface {
	ReadConfig(driving.ReadConfig) error
	ReadTime(driving.ReadTime) error
	SetTime(driving.SetTime) error
}
