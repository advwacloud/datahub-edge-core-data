package errors

type BadValueDescriptor struct {
	VName string
}

func (b BadValueDescriptor) Error() string {
	return "The value descriptor is either not unique or not formatted properly.  ValueDescriptorName: " + b.VName
}
