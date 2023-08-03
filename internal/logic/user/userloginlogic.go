package user

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	md5HashUtil *helper.HashUtil
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		md5HashUtil: helper.NewHashUtil(),
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.Response, err error) {
	var user model.SystemUsers
	err = l.svcCtx.DB.Where("uname = ?", req.UserName).First(&user).Error
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "用户名不存在",
		}, nil
	}

	l.md5HashUtil.Salt = user.Salt
	// Check the password
	if l.md5HashUtil.CheckPassword(req.PassWord, user.Pass) {
		return &types.Response{
			Code:    200,
			Message: "登录成功",
			Data:    user,
		}, nil
	} else {
		return &types.Response{
			Code:    500,
			Message: "密码错误",
		}, nil
	}
}
