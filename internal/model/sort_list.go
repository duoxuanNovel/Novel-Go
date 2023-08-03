package model

type ArticleSortList struct {
	Articleid   int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Articlename string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Author      string `gorm:"column:author;NOT NULL" json:"author"`
	Intro       string `gorm:"column:intro" json:"intro"`
	Cover       string `json:"cover"`
}
