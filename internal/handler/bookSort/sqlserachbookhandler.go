package bookSort

import (
	"net/http"

	"ddxs-api/internal/logic/bookSort"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SqlSerachBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SqlSerachBookReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bookSort.NewSqlSerachBookLogic(r.Context(), svcCtx)
		resp, err := l.SqlSerachBook(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
