package model

type ArticleArticle struct {
	Articleid     int    `gorm:"column:articleid;primary_key;AUTO_INCREMENT" json:"articleid"`
	Siteid        int    `gorm:"column:siteid;default:0;NOT NULL" json:"siteid"`
	Postdate      int    `gorm:"column:postdate;default:0;NOT NULL" json:"postdate"`
	Lastupdate    int    `gorm:"column:lastupdate;default:0;NOT NULL" json:"lastupdate"`
	Infoupdate    int    `gorm:"column:infoupdate;default:0;NOT NULL" json:"infoupdate"`
	Articlename   string `gorm:"column:articlename;NOT NULL" json:"articlename"`
	Articlecode   string `gorm:"column:articlecode;NOT NULL" json:"articlecode"`
	Backupname    string `gorm:"column:backupname;NOT NULL" json:"backupname"`
	Keywords      string `gorm:"column:keywords;NOT NULL" json:"keywords"`
	Roles         string `gorm:"column:roles;NOT NULL" json:"roles"`
	Initial       string `gorm:"column:initial;NOT NULL" json:"initial"`
	Authorid      int    `gorm:"column:authorid;default:0;NOT NULL" json:"authorid"`
	Author        string `gorm:"column:author;NOT NULL" json:"author"`
	Posterid      int    `gorm:"column:posterid;default:0;NOT NULL" json:"posterid"`
	Poster        string `gorm:"column:poster;NOT NULL" json:"poster"`
	Sortid        int    `gorm:"column:sortid;default:0;NOT NULL" json:"sortid"`
	Typeid        int    `gorm:"column:typeid;default:0;NOT NULL" json:"typeid"`
	Intro         string `gorm:"column:intro" json:"intro"`
	Lastchapterid int    `gorm:"column:lastchapterid;default:0;NOT NULL" json:"lastchapterid"`
	Lastchapter   string `gorm:"column:lastchapter;NOT NULL" json:"lastchapter"`
	Lastsummary   string `gorm:"column:lastsummary" json:"lastsummary"`
	Chapters      int    `gorm:"column:chapters;default:0;NOT NULL" json:"chapters"`
	Words         int    `gorm:"column:words;default:0;NOT NULL" json:"words"`
	Lastvisit     int    `gorm:"column:lastvisit;default:0;NOT NULL" json:"lastvisit"`
	Dayvisit      int    `gorm:"column:dayvisit;default:0;NOT NULL" json:"dayvisit"`
	Weekvisit     int    `gorm:"column:weekvisit;default:0;NOT NULL" json:"weekvisit"`
	Monthvisit    int    `gorm:"column:monthvisit;default:0;NOT NULL" json:"monthvisit"`
	Allvisit      int    `gorm:"column:allvisit;default:0;NOT NULL" json:"allvisit"`
	Lastvote      int    `gorm:"column:lastvote;default:0;NOT NULL" json:"lastvote"`
	Dayvote       int    `gorm:"column:dayvote;default:0;NOT NULL" json:"dayvote"`
	Weekvote      int    `gorm:"column:weekvote;default:0;NOT NULL" json:"weekvote"`
	Monthvote     int    `gorm:"column:monthvote;default:0;NOT NULL" json:"monthvote"`
	Allvote       int    `gorm:"column:allvote;default:0;NOT NULL" json:"allvote"`
	Goodnum       int    `gorm:"column:goodnum;default:0;NOT NULL" json:"goodnum"`
	Fullflag      int    `gorm:"column:fullflag;default:0;NOT NULL" json:"fullflag"`
	Imgflag       int    `gorm:"column:imgflag;default:0;NOT NULL" json:"imgflag"`
	Display       int    `gorm:"column:display;default:0;NOT NULL" json:"display"`
	Rgroup        int    `gorm:"column:rgroup;default:0;NOT NULL" json:"rgroup"`
	Lastvolume    string `gorm:"column:lastvolume;NOT NULL" json:"lastvolume"`
	Lastvolumeid  int    `gorm:"column:lastvolumeid;default:0;NOT NULL" json:"lastvolumeid"`
	Pushed        int    `gorm:"column:pushed;default:0;NOT NULL" json:"pushed"`
	Ratenum       int    `gorm:"column:ratenum;default:0;NOT NULL" json:"ratenum"`
	Ratesum       int    `gorm:"column:ratesum;default:0;NOT NULL" json:"ratesum"`
}

func (article *ArticleArticle) TableName() string {
	return "shipsay_article_article"
}
