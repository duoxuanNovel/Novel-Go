package bookInfo

import (
	"net/http"

	"ddxs-api/internal/logic/bookInfo"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChaterInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChapterContextReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bookInfo.NewChaterInfoLogic(r.Context(), svcCtx)
		resp, err := l.ChaterInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
