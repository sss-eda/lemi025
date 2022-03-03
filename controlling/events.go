package controlling

import "github.com/sss-eda/lemi025/controlling/application/acquire"

// Events TODO
type Events interface {
	acquire.ConfigReadEvent | acquire.TimeReadEvent | acquire.TimeSetEvent
}

// EventHandlerFunc TODO
type EventHandlerFunc[Event Events] func(*Event)
