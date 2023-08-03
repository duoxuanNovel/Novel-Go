package bookInfo

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

type HotBookTuiJianLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewHotBookTuiJianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotBookTuiJianLogic {
	return &HotBookTuiJianLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *HotBookTuiJianLogic) HotBookTuiJian(req *types.HotBookTuiJianReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:hotbooktuijian:%d:%d:%d", req.Order, req.Page, req.Size)

	var articles []model.ArticleHotBookTuiJian

	if l.svcCtx.Config.CacheTimeFlag.Book.HotBookTuiJian {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &articles)
	}

	if articles == nil {
		articles, err = l.getFromDB(req.Order, req.Page, req.Size)
	}

	if l.svcCtx.Config.CacheTimeFlag.Book.HotBookTuiJian {
		err = l.redisCacheHelper.SetToCache(cacheKey, articles, time.Duration(l.svcCtx.Config.CacheTime.Book.HotBookTuiJian)*time.Second)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    articles,
	}, nil
}

func (l *HotBookTuiJianLogic) getFromDB(order, page, size int) (articles []model.ArticleHotBookTuiJian, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename").Where("display <> 1 ").Where("words >= 0").Order(l.getOrder(order)).Offset(utils.CalculateOffset(page, size)).Limit(size).Find(&articles).Error
	return
}

func (l *HotBookTuiJianLogic) getOrder(order int) string {
	switch order {
	case 1:
		return "monthvisit DESC"
	case 2:
		return "monthvisit ASC"
	case 3:
		return "lastupdate DESC"
	case 4:
		return "lastupdate ASC"
	case 5:
		return "dayvisit DESC"
	case 6:
		return "dayvisit ASC"
	default:
		return "articleid DESC"
	}
}
