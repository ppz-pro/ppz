package api

type _Error struct {
	Code  int
	Error error
	Msg   string
}

type _Result struct {
	Data any
	Err  _Error
}

func ErrResult(msg string, err error) _Result {
	return ErrorResult(0, msg, err)
}
func ErrorResult(code int, msg string, err error) _Result {
	// log.error(code, msg, err)
	return _Result{
		Err: _Error{
			Msg:   msg,
			Error: err,
			Code:  code,
		},
	}
}

func NewResult(data any) _Result {
	return _Result{Data: data}
}
