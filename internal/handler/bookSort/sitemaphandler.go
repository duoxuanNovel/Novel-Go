package bookSort

import (
	"net/http"

	"ddxs-api/internal/logic/bookSort"
	"ddxs-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SiteMapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := bookSort.NewSiteMapLogic(r.Context(), svcCtx)
		resp, err := l.SiteMap()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
