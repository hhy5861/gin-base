package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hhy5861/gin-base/models"
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/gin_controller"
	"sync"
)

type (
	UserUtils struct {
		gin_controller.Controller
	}
)

var (
	u     *UserUtils
	once  sync.Once
	tools *common.Tools
)

func NewUtils() *UserUtils {
	once.Do(func() {
		u = &UserUtils{}
	})

	tools = common.NewTools()
	return u
}

func (ctl *UserUtils) GetValueUserInfo(ctx *gin.Context) (*models.AccountModel, error) {
	body := ctx.Value(ctx.GetHeader("x-auth-token"))
	if body == nil {
		return nil, errors.New("get user info is nil")
	}

	return body.(*models.AccountModel), nil
}
