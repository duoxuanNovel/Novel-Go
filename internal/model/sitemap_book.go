package model

type ArticleSiteMap struct {
	Articleid  int     `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Lastupdate IntTime `gorm:"column:lastupdate;default:0;NOT NULL" json:"lastupdate"`
}

func (article *ArticleSiteMap) TableName() string {
	return "shipsay_article_article"
}
