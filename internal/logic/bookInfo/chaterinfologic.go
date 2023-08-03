package bookInfo

import (
	"context"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChaterInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChaterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChaterInfoLogic {
	return &ChaterInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChaterInfoLogic) ChaterInfo(req *types.ChapterContextReq) (resp *types.Response, err error) {
	return
}
