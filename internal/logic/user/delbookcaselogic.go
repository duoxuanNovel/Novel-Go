package user

import (
	"context"
	"ddxs-api/internal/model"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBookCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelBookCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBookCaseLogic {
	return &DelBookCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelBookCaseLogic) DelBookCase(req *types.DelBookCaseReq) (resp *types.Response, err error) {
	var bookCase model.ArticleBookcase
	err = l.svcCtx.DB.Where("caseid = ?", req.CaseId).Delete(&bookCase).Error
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "删除失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "删除成功",
	}, nil
}
