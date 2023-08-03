package user

import (
	"context"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBookCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserBookCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBookCaseLogic {
	return &UserBookCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserBookCaseLogic) UserBookCase(req *types.UserBookCaseReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
