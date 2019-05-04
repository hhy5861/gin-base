package models

import "github.com/hhy5861/go-common/mysql"

type (
	AccountSecurityModel struct {
		Id        int64  `json:"id" gorm:"column:id"`
		Uuid      string `json:"uuid" gorm:"column:uuid"`
		AccountId int64  `json:"accountId" gorm:"column:accountId"`
		Salt      string `json:"salt" gorm:"column:salt"`
		Mode      string `json:"mode" gorm:"column:mode"`
		Password  string `json:"password" gorm:"column:password"`
		mysql.Model
	}
)


func (m *AccountSecurityModel) TableName() string {
	return "t_account_security"
}