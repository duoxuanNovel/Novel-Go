package model

type LastChapter struct {
	Chapterid   int    `gorm:"column:chapterid;primary_key;AUTO_INCREMENT" json:"chapterid"`
	Chaptername string `gorm:"column:chaptername;NOT NULL" json:"chaptername"`
}
