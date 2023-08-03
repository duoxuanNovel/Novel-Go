package bookSort

import (
	"context"
	"ddxs-api/internal/helper"
	"ddxs-api/internal/model"
	"ddxs-api/internal/utils"
	"fmt"
	"strconv"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortBookListLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	redisCacheHelper *helper.RedisCacheHelper
}

func NewSortBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SortBookListLogic {
	return &SortBookListLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		redisCacheHelper: helper.NewRedisCacheHelper(ctx, svcCtx),
	}
}

func (l *SortBookListLogic) SortBookList(req *types.SortBookListReq) (resp *types.Response, err error) {
	cacheKey := fmt.Sprintf("home:sortbooklist:%d:%d:%d:%d", req.SortId, req.Status, req.Page, req.Size)

	var sortList model.ArticleBookSort
	var articles []model.ArticleTuiJian
	var count int64

	if l.svcCtx.Config.Redis.Enabled {
		err = l.redisCacheHelper.GetFromCache(cacheKey, &sortList)
	}

	if articles == nil {
		articles, count, err = l.getFromDB(req.SortId, req.Status, req.Page, req.Size)
		sortList.Count = count
		sortList.List = articles
	}

	if l.svcCtx.Config.Redis.Enabled {
		err = l.redisCacheHelper.SetToCache(cacheKey, sortList)
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    sortList,
	}, nil
}

func (l *SortBookListLogic) getFromDB(sortId, statusId, page, size int) (articles []model.ArticleTuiJian, count int64, err error) {
	err = l.svcCtx.DB.Select("articleid,articlename,intro,author").Where("display <> 1 ").Where("words >= 0").Where(l.getSort(sortId)).Where(l.getStatus(statusId)).Order("lastupdate DESC").Offset(utils.CalculateOffset(page, size)).Limit(size).Find(&articles).Error
	var article model.ArticleArticle
	err = l.svcCtx.DB.Where("display <> 1 ").Where("words >= 0").Where(l.getSort(sortId)).Where(l.getStatus(statusId)).Order("lastupdate DESC").Find(&article).Count(&count).Error
	return
}

func (l *SortBookListLogic) getSort(sortId int) string {
	if sortId == 0 {
		return ""
	}
	return "sortid = " + strconv.Itoa(sortId)
}

func (l *SortBookListLogic) getStatus(statusId int) string {
	if statusId == 2 {
		return ""
	}
	return "fullflag = " + strconv.Itoa(statusId)
}
