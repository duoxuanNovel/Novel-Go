package model

type BookCaseList struct {
	Articleid       int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Articlename     string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Author          string `gorm:"column:author;NOT NULL" json:"author"`
	Lastchapterid   int    `gorm:"column:lastchapterid;default:0;NOT NULL" json:"lastchapterid"`
	Lastchapter     string `gorm:"column:lastchapter;NOT NULL" json:"lastchapter"`
	BookChapterId   int    `gorm:"column:chapterid;default:0;NOT NULL" json:"chapterid"`
	BookChapterName string `gorm:"column:chaptername;NOT NULL" json:"chaptername"`
	Cover           string `gorm:"column:cover;NOT NULL" json:"cover"`
	CaseId          int    `gorm:"column:caseid;primary_key;AUTO_INCREMENT" json:"caseid"`
}
