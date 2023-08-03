package model

type SystemUsers struct {
	Uid       int    `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"uid"`
	Uname     string `gorm:"column:uname;NOT NULL" json:"uname"`
	Name      string `gorm:"column:name;NOT NULL" json:"name"`
	Pass      string `gorm:"column:pass;NOT NULL" json:"pass"`
	Salt      string `gorm:"column:salt;NOT NULL" json:"salt"`
	Groupid   int    `gorm:"column:groupid;default:0;NOT NULL" json:"groupid"`
	Regdate   int    `gorm:"column:regdate;default:0;NOT NULL" json:"regdate"`
	Lastlogin int    `gorm:"column:lastlogin;default:0;NOT NULL" json:"lastlogin"`
	Email     string `gorm:"column:email;NOT NULL" json:"email"`
}

func (systemUsers *SystemUsers) TableName() string {
	return "shipsay_system_users"
}
