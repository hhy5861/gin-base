package account

import (
	"github.com/hhy5861/gin-base/models"
	"github.com/hhy5861/go-common/mysql"
)

type (
	AccountSecurityService struct {
	}
)

func NewAccountSecurityService() *AccountSecurityService {
	return &AccountSecurityService{}
}

func (svc *AccountSecurityService) GetAccountSecurityByUuid(uuid string) (*models.AccountSecurityModel, error) {
	var accountSecurity models.AccountSecurityModel

	err := mysql.NewDBClient().Where("uuid = ? AND valid = 0", uuid).First(&accountSecurity).Error

	return &accountSecurity, err
}
