package model

type ArticleChapter struct {
	Chapterid    int    `gorm:"column:chapterid;primary_key;AUTO_INCREMENT" json:"chapterid"`
	Articleid    int    `gorm:"column:articleid;default:0;NOT NULL" json:"articleid"`
	Articlename  string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Posterid     int    `gorm:"column:posterid;default:0;NOT NULL" json:"posterid"`
	Poster       string `gorm:"column:poster;NOT NULL" json:"poster"`
	Postdate     int    `gorm:"column:postdate;default:0;NOT NULL" json:"postdate"`
	Lastupdate   int    `gorm:"column:lastupdate;default:0;NOT NULL" json:"lastupdate"`
	Chaptername  string `gorm:"column:chaptername;NOT NULL" json:"chaptername"`
	Chapterorder int    `gorm:"column:chapterorder;default:0;NOT NULL" json:"chapterorder"`
	Words        int    `gorm:"column:words;default:0;NOT NULL" json:"words"`
	Chaptertype  int    `gorm:"column:chaptertype;default:0;NOT NULL" json:"chaptertype"`
	Attachment   string `gorm:"column:attachment" json:"attachment"`
	Summary      string `gorm:"column:summary" json:"summary"`
	Isimage      int    `gorm:"column:isimage;default:0;NOT NULL" json:"isimage"`
	Volumeid     int    `gorm:"column:volumeid;default:0;NOT NULL" json:"volumeid"`
	Pushed       int    `gorm:"column:pushed;default:0;NOT NULL" json:"pushed"`
}

type ChapterInfo struct {
	CurrentChapter ArticleChapter `json:"current_chapter"`
	NextChapter    ArticleChapter `json:"next_chapter"`
	PrevChapter    ArticleChapter `json:"prev_chapter"`
}
