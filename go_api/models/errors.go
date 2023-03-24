package models

const (
	CodeInvalidArgument = iota
	CodeNotFound
	CodeInternalError
)

type InternalErrResp struct {
	Orig         error
	InternalCode int
}

func (e InternalErrResp) Error() string {
	return e.Orig.Error()
}

func (e InternalErrResp) Code() int {
	return e.InternalCode
}
