package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Redis struct {
		Host            string `json:"host"`
		Port            int    `json:"port"`
		Password        string `json:"password"`
		Db              int    `json:"db"`
		Enabled         bool   `json:"enabled"`
		CacheExpiration int    `json:"cacheExpiration"`
	}
	Mysql struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Dbname      string `json:"dbname"`
		TablePrefix string `json:"tablePrefix"`
	}
	SiteConfig struct {
		ImgUrl string `json:"imgUrl"`
	}
	ZsConfig struct {
		Api struct {
			BaseURL      string `json:"baseURL"`
			Username     string `json:"username"`
			Password     string `json:"password"`
			UserAgent    string `json:"userAgent"`
			LastSyncFile string `json:"lastSyncFile"`
			IndexName    string `json:"indexName"`
		}
	}
	CacheTime struct {
		Home struct {
			HotTuiJian     int `json:"hotTuiJian"`
			ClassicTuiJian int `json:"classicTuiJian"`
			PostUpData     int `json:"postUpData"`
		}
		Book struct {
			BookInfo         int `json:"bookInfo"`
			BookLastChapter  int `json:"bookLastChapter"`
			ChapterList      int `json:"chapterList"`
			HotBookTuiJian   int `json:"hotBookTuiJian"`
			XiangGuanTuiJian int `json:"xiangGuanTuiJian"`
			AuthorALl        int `json:"authorALl"`
		}
	}
	CacheTimeFlag struct {
		Home struct {
			HotTuiJian     bool `json:"hotTuiJian"`
			ClassicTuiJian bool `json:"classicTuiJian"`
			PostUpData     bool `json:"postUpData"`
		}
		Book struct {
			BookInfo         bool `json:"bookInfo"`
			BookLastChapter  bool `json:"bookLastChapter"`
			ChapterList      bool `json:"chapterList"`
			HotBookTuiJian   bool `json:"hotBookTuiJian"`
			XiangGuanTuiJian bool `json:"xiangGuanTuiJian"`
			AuthorALl        bool `json:"authorALl"`
		}
	}
}
