package errors

type UnsupportedPublisher struct {
}

func (e UnsupportedPublisher) Error() string {
	return "The publisher type is not supported"
}
