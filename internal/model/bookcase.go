package model

type ArticleBookcase struct {
	Caseid      int    `gorm:"column:caseid;primary_key;AUTO_INCREMENT" json:"caseid"`
	Articleid   int    `gorm:"column:articleid;default:0;NOT NULL" json:"articleid"`
	Articlename string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Userid      int    `gorm:"column:userid;default:0;NOT NULL" json:"userid"`
	Username    string `gorm:"column:username;NOT NULL" json:"username"`
	Chapterid   int    `gorm:"column:chapterid;default:0;NOT NULL" json:"chapterid"`
	Chaptername string `gorm:"column:chaptername;NOT NULL" json:"chaptername"`
}

func (articleBookcase *ArticleBookcase) TableName() string {
	return "shipsay_article_bookcase"
}
