package home

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"strings"
	"time"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotTuiJianLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewHotTuiJianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotTuiJianLogic {
	return &HotTuiJianLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *HotTuiJianLogic) HotTuiJian(req *types.HotTuiJianReq) (resp *types.Response, err error) {
	cacheKey := "home:hot:" + strings.Join(req.Ids, ",")

	var articles []model.ArticleTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Home.HotTuiJian {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Ids)
	}

	if l.svcCtx.Config.CacheTimeFlag.Home.HotTuiJian {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Home.HotTuiJian)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *HotTuiJianLogic) getFromDB(ids []string) (articles []model.ArticleTuiJian, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,intro,author").Where("display <> 1 ").Where("words >= 0").Where("articleid IN (?)", ids).Order("FIELD(articleid, " + strings.Join(ids, ",") + ")").Find(&articles).Error
	return
}
