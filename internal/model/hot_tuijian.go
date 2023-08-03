package model

import (
	"ddxs-api/internal/utils"
	json "github.com/bytedance/sonic"
)

type ArticleTuiJian struct {
	Articleid   int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Articlename string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Author      string `gorm:"column:author;NOT NULL" json:"author"`
	Intro       string `gorm:"column:intro" json:"intro"`
	Cover       string `json:"cover"`
}

func (article *ArticleTuiJian) TableName() string {
	return "shipsay_article_article"
}

func (article *ArticleTuiJian) getCover() string {
	return utils.GetImgURL(article.Articleid)
}

func (article *ArticleTuiJian) MarshalJSON() ([]byte, error) {
	type Alias ArticleTuiJian
	return json.Marshal(&struct {
		*Alias
		Cover string `json:"cover"`
	}{
		Cover: article.getCover(),
		Alias: (*Alias)(article),
	})
}

type ArticleBookSort struct {
	Count int64            `json:"count"`
	List  []ArticleTuiJian `json:"list"`
}
