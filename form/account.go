package form

import "github.com/hhy5861/go-common/mysql"

type (
	AccountLoginForm struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	AccountListForm struct {
		mysql.QueryPage
	}
)
