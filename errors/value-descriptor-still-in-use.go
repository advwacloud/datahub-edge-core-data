package errors

type ValueDescriptorStillInUse struct {
}

func (e ValueDescriptorStillInUse) Error() string {
	return "Can't delete the value descriptor, its still referenced by readings"
}
