package errors

type NoValueDescriptor struct {
	Id string
}

func (e NoValueDescriptor) Error() string {
	return "There is no value descriptor for the following reading: " + e.Id
}
