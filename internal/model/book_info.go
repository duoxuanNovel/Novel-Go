package model

import "ddxs-api/internal/utils"

type ArticleBookInfo struct {
	Articleid     int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Lastupdate    int    `gorm:"column:lastupdate;default:0;NOT NULL" json:"lastupdate"`
	Articlename   string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Author        string `gorm:"column:author;NOT NULL" json:"author"`
	Sortid        int    `gorm:"column:sortid;default:0;NOT NULL" json:"sortid"`
	Intro         string `gorm:"column:intro" json:"intro"`
	Fullflag      int    `gorm:"column:fullflag;default:0;NOT NULL" json:"fullflag"`
	Lastchapterid int    `gorm:"column:lastchapterid;default:0;NOT NULL" json:"lastchapterid"`
	Lastchapter   string `gorm:"column:lastchapter;NOT NULL" json:"lastchapter"`
	Cover         string `json:"cover"`
}

func (article *ArticleBookInfo) TableName() string {
	return "shipsay_article_article"
}

func (article *ArticleBookInfo) GetCover() string {
	return utils.GetImgURL(article.Articleid)
}
