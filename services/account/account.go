package account

import (
	"errors"
	"github.com/hhy5861/gin-base/form"
	"github.com/hhy5861/gin-base/models"
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/jwt"
	"github.com/hhy5861/go-common/mysql"
)

type (
	AccountService struct {
		tools *common.Tools
	}

	AccountLoginService struct {
		GetAccount         GetAccountFunc
		GetAccountSecurity GetAccountSecurityFunc
		JwtConfig          *jwt.JwtConfig
	}

	LoginResult struct {
		models.AccountModel
		Token string `json:"token"`
	}

	GetAccountFunc func(formData form.AccountLoginForm) (*models.AccountModel, error)

	GetAccountSecurityFunc func(uuid string) (*models.AccountSecurityModel, error)
)

const (
	DisableErrors  = "User has been disabled"
	PasswordErrors = "Incorrect password"
)

func NewAccountService() *AccountService {
	return &AccountService{tools: common.NewTools()}
}

func NewAccountLoginService(accountFunc GetAccountFunc, securityFunc GetAccountSecurityFunc, jwtConfig *jwt.JwtConfig) *AccountLoginService {
	return &AccountLoginService{
		GetAccount:         accountFunc,
		GetAccountSecurity: securityFunc,
		JwtConfig:          jwtConfig,
	}
}

func (svc *AccountLoginService) Login(formData form.AccountLoginForm) (*LoginResult, error) {
	accountInfo, err := svc.GetAccount(formData)
	if err != nil {
		return nil, err
	}

	if accountInfo.Status == 0 {
		return nil, errors.New(DisableErrors)
	}

	security, err := svc.GetAccountSecurity(accountInfo.Uuid)
	if err != nil {
		return nil, err
	}

	ok := common.NewPassword(&common.Password{
		Password: formData.Password,
		Salt:     security.Salt,
		Mode:     security.Mode,
	}).Hash256().EqualsPassword(security.Password)
	if !ok {
		return nil, errors.New(PasswordErrors)
	}

	token, err := jwt.NewJwtPackage(svc.JwtConfig).CreateToken(&jwt.StandardClaims{
		Id:   uint(accountInfo.Id),
		Uuid: accountInfo.Uuid,
	})
	if err != nil {
		return nil, err
	}

	return &LoginResult{*accountInfo, token}, nil
}

func (svc *AccountService) GetAccountInfoByName(formData form.AccountLoginForm) (*models.AccountModel, error) {
	var account models.AccountModel

	err := mysql.NewDBClient().Where("name =? OR email = ? OR mobile = ?", formData.Username, formData.Username, formData.Username).First(&account).Error
	return &account, err
}

func (svc *AccountService) GetAccountList(formData form.AccountListForm) (*mysql.QueryResult, error) {
	var list []*models.AccountModel

	query := mysql.NewDBClient()

	query = query.Order("created_at DESC")
	return new(mysql.Service).QueryResList(query,
		new(models.AccountModel).TableName(), &list,
		svc.tools.Int64Value(formData.Page),
		svc.tools.Int64Value(formData.Size))
}
