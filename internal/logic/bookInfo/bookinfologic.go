package bookInfo

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"fmt"
	"time"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BookInfoLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewBookInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BookInfoLogic {
	return &BookInfoLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *BookInfoLogic) BookInfo(req *types.BookInfoReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("book:info:%d", req.Id)
	var articles model.ArticleBookInfo

	if l.svcCtx.Config.CacheTimeFlag.Book.BookInfo {
		err = l.redisCacheHelper.GetFromCache(cacheKey, articles)
	}

	if articles.Articleid == 0 {
		articles, err = l.getFromDB(req.Id) // 注意这里，我们要确保从数据库获取的是一个指针
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.BookInfo {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.BookInfo)*time.Second) // 缓存需要存放结构体，不是指针
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *BookInfoLogic) getFromDB(id int) (articles model.ArticleBookInfo, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,intro,lastupdate,author,sortid,fullflag,lastchapterid,lastchapter").Where("display <> 1 ").Where("words >= 0").Where("articleid = ?", id).Order("articleid DESC").First(&articles).Error
	if err == nil {
		articles.Cover = articles.GetCover()
	}
	return
}
