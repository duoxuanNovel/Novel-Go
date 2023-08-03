package home

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

type UpdataPostDataLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewUpdataPostDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdataPostDataLogic {
	return &UpdataPostDataLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *UpdataPostDataLogic) UpdataPostData(req *types.UpdataPostDataReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:male:list:%d:%d:%d", req.Order, req.Page, req.Size)

	var articles []model.ArticleClassicTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Home.PostUpData {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Order, req.Page, req.Size)
	}

	if l.svcCtx.Config.CacheTimeFlag.Home.PostUpData {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Home.PostUpData)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil

}

func (l *UpdataPostDataLogic) getFromDB(order, page, size int) (articles []model.ArticleClassicTuiJian, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,author,sortid").Where("display <> 1 ").Where("words >= 0").Order(l.order(order)).Offset(utils.CalculateOffset(page, size)).Limit(size).Find(&articles).Error
	return
}

func (l *UpdataPostDataLogic) order(order int) string {
	switch order {
	case 1:
		return "lastupdate DESC"
	case 2:
		return "postdate DESC"
	case 3:
		return "monthvisit DESC"
	default:
		return "articleid DESC"
	}
}
