package home

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"ddxs-api/internal/utils"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassicTuiJianLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewClassicTuiJianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassicTuiJianLogic {
	return &ClassicTuiJianLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *ClassicTuiJianLogic) ClassicTuiJian(req *types.PageReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:classic:%d:%d", req.Page, req.Size)

	var articles []model.ArticleClassicTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Home.ClassicTuiJian {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Page, req.Size)
	}

	if l.svcCtx.Config.CacheTimeFlag.Home.ClassicTuiJian {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Home.ClassicTuiJian)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *ClassicTuiJianLogic) getFromDB(page, size int) (articles []model.ArticleClassicTuiJian, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,sortid,author").Where("display <> 1 ").Where("words >= 0").Order("monthvisit DESC").Offset(utils.CalculateOffset(page, size)).Limit(size).Find(&articles).Error
	return
}
