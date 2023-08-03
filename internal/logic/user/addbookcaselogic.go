package user

import (
	"context"
	"ddxs-api/internal/model"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBookCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBookCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBookCaseLogic {
	return &AddBookCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBookCaseLogic) AddBookCase(req *types.AddBookCaseReq) (resp *types.Response, err error) {
	var bookCase = model.ArticleBookcase{
		Articleid:   req.Articleid,
		Articlename: req.Articlename,
		Userid:      req.Userid,
		Username:    req.Username,
		Chapterid:   req.Chapterid,
		Chaptername: req.Chaptername,
	}
	// 判断是否已经加入书架 如果已经加入书架则更新数据
	var bookCaseCount int64
	l.svcCtx.DB.Model(&model.ArticleBookcase{}).Where("articleid = ? and userid = ?", req.Articleid, req.Userid).Count(&bookCaseCount)
	if bookCaseCount > 0 {
		err = l.svcCtx.DB.Model(&model.ArticleBookcase{}).Where("articleid = ? and userid = ?", req.Articleid, req.Userid).Update("chapterid", req.Chapterid).Update("chaptername", req.Chaptername).Error
		if err != nil {
			return &types.Response{
				Code:    500,
				Message: "加入书架失败",
			}, nil
		}
		return &types.Response{
			Code:    200,
			Message: "更新书架成功",
		}, nil
	}
	err = l.svcCtx.DB.Create(&bookCase).Error
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "加入书架失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "加入书架成功",
	}, nil
}
