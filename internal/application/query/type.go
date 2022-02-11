package query

// Type TODO
type Type interface {
	ReadConfig | ReadTime
}

type ReadConfig struct{}

type ReadTime struct{}