package bookInfo

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/utils"
	"fmt"
	"time"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BookLastChapterLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewBookLastChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BookLastChapterLogic {
	return &BookLastChapterLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *BookLastChapterLogic) BookLastChapter(req *types.LastChapterReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:chapter:last:%d:%d", req.Id, req.Limit)

	var articles []model.LastChapter

	if l.svcCtx.Config.CacheTimeFlag.Book.BookLastChapter {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Id, req.Limit)
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.BookLastChapter {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.BookLastChapter)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *BookLastChapterLogic) getFromDB(id, limit int) (articles []model.LastChapter, err error) {
	err = l.svcCtx.DB.Select("chapterid,chaptername").Table(utils.GetChapterTableName(id)).Where("articleid = ?", id).Order("chapterorder DESC").Limit(limit).Find(&articles).Error
	return
}
