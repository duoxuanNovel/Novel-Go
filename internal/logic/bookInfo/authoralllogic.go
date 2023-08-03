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

type AuthorAllLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewAuthorAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorAllLogic {
	return &AuthorAllLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *AuthorAllLogic) AuthorAll(req *types.AuthorAllReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("book:author:all:%s", req.Name)
	var articles []model.ArticleTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Book.AuthorALl {
		err = l.redisCacheHelper.GetFromCache(cacheKey, articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Name) // 注意这里，我们要确保从数据库获取的是一个指针
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.AuthorALl {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.AuthorALl)*time.Second) // 缓存需要存放结构体，不是指针
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *AuthorAllLogic) getFromDB(name string) (articles []model.ArticleTuiJian, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,intro,author").Where("display <> 1 ").Where("words >= 0").Where("author = ?", name).Order("articleid DESC").Find(&articles).Error
	return
}
