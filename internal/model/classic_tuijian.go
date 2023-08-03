package model

type ArticleClassicTuiJian struct {
	Articleid   int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Articlename string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Author      string `gorm:"column:author;NOT NULL" json:"author"`
	Sortid      int    `gorm:"column:sortid;default:0;NOT NULL" json:"sortid"`
}

func (article *ArticleClassicTuiJian) TableName() string {
	return "shipsay_article_article"
}
