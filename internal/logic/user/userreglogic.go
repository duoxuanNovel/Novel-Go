package user

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	md5HashUtil *helper.HashUtil
}

func NewUserRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegLogic {
	return &UserRegLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		md5HashUtil: helper.NewHashUtil(),
	}
}

func (l *UserRegLogic) UserReg(req *types.UserRegReq) (resp *types.Response, err error) {
	encryptedPassword := l.md5HashUtil.EncryptPassword(req.PassWord)
	//获取现在的时间的时间戳转为int
	var user = model.SystemUsers{
		Uname:   req.UserName,
		Name:    req.UserName,
		Email:   req.Email,
		Salt:    l.md5HashUtil.Salt,
		Pass:    encryptedPassword,
		Groupid: 3,
		Regdate: int(time.Now().Unix()),
	}
	err = l.svcCtx.DB.Create(&user).Error

	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "注册失败",
		}, nil
	}
	
	return &types.Response{
		Code:    200,
		Message: "注册成功",
	}, nil
}
