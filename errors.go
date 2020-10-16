package antchain

type AccessError struct {
	message string
}

func (e AccessError) Error() string {
	return e.message
}
