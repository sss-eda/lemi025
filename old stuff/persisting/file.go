package persisting

// File - Our projection
type File struct{}

type Writer interface{}

// OnDataPointAcquired TODO
func (file *File) OnDataPointAcquired(
	event DataPointAcquiredEvent,
	writer Writer,
) {
	file.encoding.Encode(event)
}
