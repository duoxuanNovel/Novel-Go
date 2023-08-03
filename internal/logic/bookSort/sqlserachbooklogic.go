package bookSort

import (
	"context"
	"ddxs-api/internal/model"
	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"net/http"
	"strings"
)

type SqlSerachBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSqlSerachBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SqlSerachBookLogic {
	return &SqlSerachBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SqlSerachBookLogic) SqlSerachBook(req *types.SqlSerachBookReq) (resp *types.Response, err error) {
	//组合“from”和“max_results”以允许分页。
	from := (req.Page - 1) * req.Size
	size := req.Size
	query := fmt.Sprintf(`{
    "search_type": "match",
    "query": {
        "term": "%s"
    },
    "from": %d,
    "max_results": %d,
    "_source": []
}`, req.Key, from, size)

	httpReq, err := http.NewRequest("POST", "http://localhost:4080/api/article/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	httpReq.SetBasicAuth("admin", "Complexpass#123")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	respHttp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(respHttp.Body)

	body, err := io.ReadAll(respHttp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var zincSearchResponse model.ZincSearchResponse
	err = sonic.Unmarshal(body, &zincSearchResponse)
	return &types.Response{
		Code:    http.StatusOK,
		Message: "接口请求成功",
		Data:    zincSearchResponse.Hits,
	}, nil
}
