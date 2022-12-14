// Code generated by protoc-gen-bic. DO NOT EDIT.
// versions:1.1.3

package api

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserServiceHttpHandler(g *gin.RouterGroup, srvs UserService) {
	tmp := &x_UserService{xx: srvs}
	g.GET("/users/v1/", tmp.GetUsers)
	g.POST("/users/v1/", tmp.PostUsers)
	g.GET("/users/v1/:id/", tmp.GetUserInfo)
	g.GET("/users/v1/:id/", tmp.GetUserInfoDownload)
}

type UserService interface {
	GetUsers(ctx *gin.Context, in *RequestUsers) (out *ResponseUsers, code ErrCode)
	PostUsers(ctx *gin.Context, in *User) (out *ResponseNil, code ErrCode)
	GetUserInfo(ctx *gin.Context, in *RequestNil, id uint) (out *User, code ErrCode)
	GetUserInfoDownload(ctx *gin.Context, in *RequestNil, id uint) (out *ResponseNil, code ErrCode)
}

// generated http handle
type UserServiceHttpHandler interface {
	GetUsers(ctx *gin.Context)
	PostUsers(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	GetUserInfoDownload(ctx *gin.Context)
}

type x_UserService struct {
	xx UserService
}

// @Summary 获取用户列表
// @Tags User-Service
// @Produce json
// @Param page query uint32 false "页码"
// @Param page_size query uint32 false "每页数量"
// @Success 200 {object} ResponseUsers
// @Failure 401 {string} string "header need Authorization data"
// @Failure 403 {string} string "no api permission or no obj permission"
// @Router /users/v1/ [GET]
func (x *x_UserService) GetUsers(ctx *gin.Context) {
	req := &RequestUsers{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "detail": " request error"})
		return
	}
	rsp, code := x.xx.GetUsers(ctx, req)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   int(code),
		"detail": code.String(),
		"data":   rsp,
	})
}

// @Summary 添加用户
// @Tags User-Service
// @Produce json
// @Param id body uint32 false "主键ID"
// @Param username body string false "用户名"
// @Param email body string false "邮箱"
// @Success 200 {object} ResponseNil
// @Failure 401 {string} string "header need Authorization data"
// @Failure 403 {string} string "no api permission or no obj permission"
// @Router /users/v1/ [POST]
func (x *x_UserService) PostUsers(ctx *gin.Context) {
	req := &User{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "detail": " request error"})
		return
	}
	rsp, code := x.xx.PostUsers(ctx, req)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   int(code),
		"detail": code.String(),
		"data":   rsp,
	})
}

// @Summary 获取用户详情
// @Tags User-Service
// @Produce json
// @Param id path uint true "some id"
// @Success 200 {object} User
// @Failure 401 {string} string "header need Authorization data"
// @Failure 403 {string} string "no api permission or no obj permission"
// @Router /users/v1/:id/ [GET]
func (x *x_UserService) GetUserInfo(ctx *gin.Context) {
	req := &RequestNil{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "detail": " request error"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   int(ErrCode_param_error),
			"detail": "param id should be int",
		})
	}
	rsp, code := x.xx.GetUserInfo(ctx, req, uint(id))

	ctx.JSON(http.StatusOK, gin.H{
		"code":   int(code),
		"detail": code.String(),
		"data":   rsp,
	})
}

// @Summary 用户下载
// @Tags User-Service
// @Produce json
// @Param id path uint true "some id"
// @Success 200 {object} ResponseNil
// @Failure 401 {string} string "header need Authorization data"
// @Failure 403 {string} string "no api permission or no obj permission"
// @Router /users/v1/:id/ [GET]
func (x *x_UserService) GetUserInfoDownload(ctx *gin.Context) {
	req := &RequestNil{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "detail": " request error"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   int(ErrCode_param_error),
			"detail": "param id should be int",
		})
	}
	_, _ = x.xx.GetUserInfoDownload(ctx, req, uint(id))
}
