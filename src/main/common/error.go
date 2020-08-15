package common

type Error interface {
	error
	ECode() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) ECode() int {
	return se.Code
}
