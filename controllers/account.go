package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hhy5861/gin-base/form"
	accountProto "github.com/hhy5861/gin-base/proto"
	accountService "github.com/hhy5861/gin-base/services/account"
	"github.com/hhy5861/go-common/gin_controller"
	"github.com/hhy5861/go-common/google_grpc"
	"golang.org/x/net/context"
	"time"
)

type (
	AccountControllers struct {
		gin_controller.Controller
	}
)

func NewAccountControllers(ctl *gin_controller.Controller) *AccountControllers {
	return &AccountControllers{gin_controller.Controller{JwtConf: ctl.JwtConf}}
}

func (ctl *AccountControllers) Login(ctx *gin.Context) {
	var formData form.AccountLoginForm

	if err := ctx.ShouldBindJSON(&formData); err != nil {
		ctl.ParamsExecption(ctx, err)
		return
	}

	account, err := accountService.NewAccountLoginService(
		accountService.NewAccountService().GetAccountInfoByName,
		accountService.NewAccountSecurityService().GetAccountSecurityByUuid,
		ctl.JwtConf).Login(formData)

	if err != nil {
		ctl.ServiceCodeExecption(ctx, 100000, err)
		return
	}

	ctl.Response(ctx, account)
}

func (ctl *AccountControllers) List(ctx *gin.Context) {
	var formData form.AccountListForm

	if err := ctx.ShouldBindQuery(&formData); err != nil {
		ctl.ParamsExecption(ctx, err)
		return
	}

	list, err := accountService.NewAccountService().GetAccountList(formData)
	if err != nil {
		ctl.ServiceCodeExecption(ctx, 100010, err)
		return
	}

	ctl.Response(ctx, list)
}

func (ctl *AccountControllers) GrpcLogin(ctx *gin.Context) {
	var formData form.AccountLoginForm

	if err := ctx.ShouldBindJSON(&formData); err != nil {
		ctl.ParamsExecption(ctx, err)
		return
	}

	c := accountProto.NewAccountClient(google_grpc.GetGrpcConn())

	contexts, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(contexts, &accountProto.LoginRequest{
		Username: formData.Username,
		Password: formData.Password,
	})
	if err != nil {
		ctl.ServiceCodeExecption(ctx, 100020, err)
		return
	}

	ctl.Response(ctx, r.Data)
}
