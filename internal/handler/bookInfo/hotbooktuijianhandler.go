package bookInfo

import (
	"net/http"

	"ddxs-api/internal/logic/bookInfo"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HotBookTuiJianHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HotBookTuiJianReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bookInfo.NewHotBookTuiJianLogic(r.Context(), svcCtx)
		resp, err := l.HotBookTuiJian(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
