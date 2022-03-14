package persisting

// File - Our projection
type File struct{}

// OnDataPointAcquired TODO
func (file *File) OnDataPointAcquired(
	event DataPointAcquiredEvent,
	writer Writer,
) {
	file.encoding.Encode(event)
}
