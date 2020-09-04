package errors

type NotFound struct {
}

func (e NotFound) Error() string {
	return "Can't find the resource"
}
