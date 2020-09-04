package errors

type UnexpectedError struct {
	Message string
}

func (e UnexpectedError) Error() string {
	return e.Message
}
