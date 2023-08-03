package user

import (
	"context"
	"ddxs-api/internal/model"
	"ddxs-api/internal/utils"

	"ddxs-api/internal/svc"
	"ddxs-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryBookCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBookCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBookCaseLogic {
	return &QueryBookCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryBookCaseLogic) QueryBookCase(req *types.QueryBookCaseReq) (resp *types.Response, err error) {
	var bookCase []model.ArticleBookcase
	err = l.svcCtx.DB.Where("userid = ?", req.Uid).Find(&bookCase).Error
	if err != nil {
		// 这里需要处理错误
	}

	// 将查出来的文章ID放到数组里面
	var bookCaseIds []int
	for _, v := range bookCase {
		bookCaseIds = append(bookCaseIds, v.Articleid)
	}

	//然后再查文章表
	var articles []model.ArticleArticle
	err = l.svcCtx.DB.Where("articleid in (?)", bookCaseIds).Find(&articles).Error
	if err != nil {
		// 这里需要处理错误
	}

	// 将查出来的收藏表数据放入一个映射中，方便后续查找
	bookCaseMap := make(map[int]model.ArticleBookcase)
	for _, v := range bookCase {
		bookCaseMap[v.Articleid] = v
	}

	//将文章表的数据和收藏表的数据进行合并
	var bookCaseList = make([]model.BookCaseList, 0)
	for _, v := range articles {
		bc, ok := bookCaseMap[v.Articleid]
		if !ok {
			// 没有找到对应的书架记录，可能需要处理此错误
			continue
		}
		bookCaseList = append(bookCaseList, model.BookCaseList{
			Articleid:       v.Articleid,
			Articlename:     v.Articlename,
			Author:          v.Author,
			Lastchapter:     v.Lastchapter,
			Lastchapterid:   v.Lastchapterid,
			Cover:           utils.GetImgURL(v.Articleid),
			BookChapterId:   bc.Chapterid,
			BookChapterName: bc.Chaptername,
			CaseId:          bc.Caseid,
		})
	}

	return &types.Response{
		Code:    200,
		Message: "查询成功",
		Data:    bookCaseList,
	}, nil
}
