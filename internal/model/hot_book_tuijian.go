package model

type ArticleHotBookTuiJian struct {
	Articleid   int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Articlename string `gorm:"column:articlename;NOT NULL" json:"articlename"`
}

func (article *ArticleHotBookTuiJian) TableName() string {
	return "shipsay_article_article"
}
