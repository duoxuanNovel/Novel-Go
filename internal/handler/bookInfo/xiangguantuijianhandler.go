package bookInfo

import (
	"net/http"

	"ddxs-api/internal/logic/bookInfo"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func XiangGuanTuiJianHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.XiangGuanTuiJianReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bookInfo.NewXiangGuanTuiJianLogic(r.Context(), svcCtx)
		resp, err := l.XiangGuanTuiJian(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
