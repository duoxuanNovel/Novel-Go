package bookInfo

import (
	"context"
	"ddxs-api/internal/utils"
	"net/http"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterContextLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChapterContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterContextLogic {
	return &ChapterContextLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterContextLogic) ChapterContext(req *types.ChapterContextReq) (resp *types.Response, err error) {
	txtUrl := utils.GetTxtURL(req.Id, req.Cid)
	content := utils.GetTxtContent(txtUrl)
	text := utils.ConvertToParagraph(content)
	return &types.Response{
		Code:    http.StatusOK,
		Message: "接口请求成功",
		Data:    text,
	}, nil
}
