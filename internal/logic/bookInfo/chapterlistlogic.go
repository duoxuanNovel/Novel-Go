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

type ChapterListLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterListLogic {
	return &ChapterListLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *ChapterListLogic) ChapterList(req *types.ChapterListReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:chapter:list:%d", req.Id)

	var articles []model.LastChapter

	if l.svcCtx.Config.CacheTimeFlag.Book.ChapterList {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Id)
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.ChapterList {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.ChapterList)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *ChapterListLogic) getFromDB(id int) (articles []model.LastChapter, err error) {
	err = l.svcCtx.DB.Select("chapterid,chaptername").Table(utils.GetChapterTableName(id)).Where("articleid = ?", id).Order("chapterorder ASC").Find(&articles).Error
	return
}
