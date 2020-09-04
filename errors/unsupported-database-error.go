package errors

type UnsupportedDatabaseError struct {
}

func (e UnsupportedDatabaseError) Error() string {
	return "The database type is not supported"
}
