package api

type ErrCode interface {
	Code() int
	String() string
}

type Err struct {
	code int
	msg  string
}

func (e Err) Code() int {
	return e.code
}

func (e Err) String() string {
	return e.msg
}
