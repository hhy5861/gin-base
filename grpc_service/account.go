package grpc_service

import (
	"github.com/hhy5861/gin-base/form"
	accountProto "github.com/hhy5861/gin-base/proto"
	accountService "github.com/hhy5861/gin-base/services/account"
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/jwt"
	"golang.org/x/net/context"
)

type (
	AccountGrpcService struct {
		tools     *common.Tools
		JwtConfig *jwt.JwtConfig
	}
)

func NewAccountGrpcService(jwtConfig *jwt.JwtConfig) *AccountGrpcService {
	return &AccountGrpcService{
		tools:     common.NewTools(),
		JwtConfig: jwtConfig,
	}
}

func (svc *AccountGrpcService) Login(ctx context.Context, in *accountProto.LoginRequest) (*accountProto.Response, error) {
	formData := form.AccountLoginForm{
		Username: in.Username,
		Password: in.Password,
	}

	account, err := accountService.NewAccountLoginService(
		accountService.NewAccountService().GetAccountInfoByName,
		accountService.NewAccountSecurityService().GetAccountSecurityByUuid,
		svc.JwtConfig).Login(formData)

	if err != nil {
		return &accountProto.Response{Code: 100000, Message: "Login failed.", Data: nil}, err
	}

	var result accountProto.AccountResult
	err = svc.tools.CopyStruct(account, &result)
	if err != nil {
		return &accountProto.Response{Code: 100001, Message: err.Error(), Data: nil}, err
	}

	return &accountProto.Response{Code: 0, Message: "success", Data: &result}, nil
}
