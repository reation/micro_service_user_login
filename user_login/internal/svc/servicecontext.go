package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_user_login/model"
	"micro_service_user_login/user_login/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	userModel := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(userModel, nil),
	}
}
