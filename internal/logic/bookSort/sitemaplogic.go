package bookSort

import (
	"context"
	"ddxs-api/internal/model"
	"net/http"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SiteMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSiteMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SiteMapLogic {
	return &SiteMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SiteMapLogic) SiteMap() (resp *types.Response, err error) {
	var sitemaps []model.ArticleSiteMap
	err = l.svcCtx.DB.Select("articleid,lastupdate").Order("articleid DESC").Find(&sitemaps).Error

	return &types.Response{
		Code:    http.StatusOK,
		Data:    sitemaps,
		Message: "接口请求成功",
	}, nil
}
