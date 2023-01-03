package logic

import (
	"context"
	"micro_service_user_login/model"
	"micro_service_user_login/tool"
	"micro_service_user_login/user_login/internal/svc"
	"micro_service_user_login/user_login/user_login"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user_login.Request) (*user_login.Response, error) {
	// todo: add your logic here and delete this line
	checkLogin, err := l.userLoginCheck(in.UserName, in.Password)
	if err != nil {
		return &user_login.Response{
			Id: 0,
		}, err
	}
	return &user_login.Response{
		Id: int32(checkLogin),
	}, nil
}

func (l *UserLoginLogic) userLoginCheck(userName string, passWord string) (result int64, err error) {
	userInfo, err := l.svcCtx.UserModel.FindInfoByUserName(l.ctx, userName)
	switch err {
	case nil:
	case model.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}

	if userInfo.Password != tool.MD5(passWord) {
		return 0, nil
	}

	return userInfo.Id, nil
}
