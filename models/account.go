package models

import "github.com/hhy5861/go-common/mysql"

type (
	AccountModel struct {
		Id              uint64 `json:"id" gorm:"column:id"`
		Uuid            string `json:"uuid" gorm:"column:uuid"`
		Name            string `json:"name" gorm:"column:name"`
		ActualName      string `json:"actualName" gorm:"column:actual_name"`
		CreatedByUserId uint32 `json:"createdByUserId" gorm:"column:created_by_user_Id"`
		Department      string `json:"department" gorm:"column:department"`
		Email           string `json:"email" gorm:"column:email"`
		Phone           string `json:"phone" gorm:"column:phone"`
		RoleId          string `json:"roleId" gorm:"column:role_id"`
		IsAdmin         uint32 `json:"isAdmin" gorm:"column:is_admin"`
		Status          uint32 `json:"status" gorm:"column:status"`
		mysql.Model
	}
)

func (m *AccountModel) TableName() string {
	return "t_account"
}
