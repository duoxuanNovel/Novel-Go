package bookInfo

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type XiangGuanTuiJianLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewXiangGuanTuiJianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *XiangGuanTuiJianLogic {
	return &XiangGuanTuiJianLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *XiangGuanTuiJianLogic) XiangGuanTuiJian(req *types.XiangGuanTuiJianReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:xiangguantuijian:%d:%d", req.Id, req.Limit)

	var articles []model.ArticleHotBookTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Book.XiangGuanTuiJian {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Id, req.Limit)
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.XiangGuanTuiJian {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.XiangGuanTuiJian)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *XiangGuanTuiJianLogic) getFromDB(id, limit int) (articles []model.ArticleHotBookTuiJian, err error) {
	// Get books with ID greater than the given ID first
	err = l.svcCtx.DB.Select("articleid, articlename").Where("display <> 1").Where("words >= 0").Where("articleid > ?", id).Order("articleid ASC").Limit(limit).Find(&articles).Error
	// If not enough books are found, get more books with ID less than the given ID
	if len(articles) < limit {
		var extraBooks []model.ArticleHotBookTuiJian
		err = l.svcCtx.DB.Select("articleid,articlename").Where("display <> 1 ").Where("words >= 0").Where("articleid < ?", id).Order("articleid DESC").Limit(limit - len(articles)).Find(&extraBooks).Error
		articles = append(extraBooks, articles...)
	}
	return
}
