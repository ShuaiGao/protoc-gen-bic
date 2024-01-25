package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type ErrCode interface {
	Code() int
	EqualTo(b ErrCode) bool
	String() string
	IsSuccess() bool
}

type Err struct {
	code   int
	msg    string
	detail error
}

func (e Err) IsSuccess() bool {
	return e.code == ECSuccess.code
}

func (e Err) Code() int {
	return e.code
}

func (e Err) String() string {
	if gin.Mode() == gin.ReleaseMode {
		return e.msg
	}
	if e.detail != nil {
		return e.msg + ", detail: " + e.detail.Error()
	}
	return e.msg
}

func (e Err) Wrap(err interface{}) ErrCode {
	switch err.(type) {
	case error:
		e.detail = err.(error)
	case string:
		e.detail = errors.New(err.(string))
	}
	return e
}

func (e Err) EqualTo(b ErrCode) bool {
	return e.Code() == b.Code()
}

func NewErr(e Err, err interface{}) Err {
	switch err.(type) {
	case error:
		e.detail = err.(error)
	case string:
		e.detail = errors.New(err.(string))
	}
	return e
}

func Equal(a, b ErrCode) bool {
	return a.Code() == b.Code()
}

var (
	ECSuccess  = Err{code: 0, msg: "success"}
	ECInternal = Err{code: 500, msg: "服务器内部错误"}
	ECDbFind   = Err{code: 1000, msg: "数据库查询错误"}
	ECDbFirst  = Err{code: 1001, msg: "数据库查找错误"}
	ECDbSave   = Err{code: 1002, msg: "数据库保存错误"}
	ECDbCreate = Err{code: 1003, msg: "数据库创建错误"}
	ECDbDelete = Err{code: 1003, msg: "数据库删除错误"}
	ECDbUpdate = Err{code: 1003, msg: "数据库更新错误"}
	ECDbCommit = Err{code: 1004, msg: "数据库事务提交错误"}
)
