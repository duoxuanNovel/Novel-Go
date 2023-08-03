package zs

import (
	"bytes"
	"ddxs-api/internal/config"
	"ddxs-api/internal/model"
	"ddxs-api/internal/utils"
	"fmt"
	json "github.com/bytedance/sonic"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type ZsConfig struct {
	Config config.Config
}

func NewZsConfig(c config.Config) *ZsConfig {
	return &ZsConfig{
		Config: c,
	}
}

type IndexConfig struct {
	Name        string      `json:"name"`
	StorageType string      `json:"storage_type"`
	ShardNum    int         `json:"shard_num"`
	Mappings    interface{} `json:"mappings"`
}

func (zs *ZsConfig) checkIndexExists(indexName string) (bool, error) {
	resp, err := http.Head(zs.Config.ZsConfig.Api.BaseURL + "/index/" + indexName)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	return resp.StatusCode == http.StatusOK, nil
}

func (zs *ZsConfig) deleteIndex(indexName string) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", zs.Config.ZsConfig.Api.BaseURL+"/index/"+indexName, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(zs.Config.ZsConfig.Api.Username, zs.Config.ZsConfig.Api.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", zs.Config.ZsConfig.Api.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Server response:", string(bodyBytes))
		}
	}
}

func (zs *ZsConfig) createOrUpdateIndex(indexConfig IndexConfig) {
	jsonValue, err := json.Marshal(indexConfig)
	if err != nil {
		fmt.Println("Error encoding index configuration:", err)
		return
	}

	client := &http.Client{}

	exists, err := zs.checkIndexExists(indexConfig.Name)
	if err != nil {
		fmt.Println("Error checking if index exists:", err)
		return
	}

	var req *http.Request

	if exists {
		req, err = http.NewRequest("PUT", zs.Config.ZsConfig.Api.BaseURL+"/index/"+indexConfig.Name, bytes.NewBuffer(jsonValue))
	} else {
		req, err = http.NewRequest("POST", zs.Config.ZsConfig.Api.BaseURL+"/index", bytes.NewBuffer(jsonValue))
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(zs.Config.ZsConfig.Api.Username, zs.Config.ZsConfig.Api.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", zs.Config.ZsConfig.Api.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Server response:", string(bodyBytes))
		}
	}
}

func (zs *ZsConfig) getLastSync() int64 {
	if _, err := os.Stat(zs.Config.ZsConfig.Api.LastSyncFile); os.IsNotExist(err) {
		// if the file doesn't exist, return 0
		return 0
	}

	data, err := os.ReadFile(zs.Config.ZsConfig.Api.LastSyncFile)
	if err != nil {
		fmt.Println("Error reading last sync file:", err)
		return 0
	}

	lastSync, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		fmt.Println("Error parsing last sync time:", err)
		return 0
	}

	return lastSync
}

func (zs *ZsConfig) setLastSync(lastSync int64) {
	// make sure the directory exists
	syncDir := filepath.Dir(zs.Config.ZsConfig.Api.LastSyncFile)
	if _, err := os.Stat(syncDir); os.IsNotExist(err) {
		err := os.MkdirAll(syncDir, os.ModePerm) // os.ModePerm is 0777
		if err != nil {
			fmt.Println("Error creating directories:", err)
			return
		}
	}

	// then write the file
	err := os.WriteFile(zs.Config.ZsConfig.Api.LastSyncFile, []byte(strconv.FormatInt(lastSync, 10)), 0666)
	if err != nil {
		fmt.Println("Error writing last sync file:", err)
	}
}

func (zs *ZsConfig) indexDocument(docID string, document map[string]interface{}) {
	client := &http.Client{}
	jsonValue, err := json.Marshal(document)
	if err != nil {
		fmt.Println("Error encoding document:", err)
		return
	}

	// Update the URL to match the API documentation
	req, err := http.NewRequest("PUT", zs.Config.ZsConfig.Api.BaseURL+"/"+zs.Config.ZsConfig.Api.IndexName+"/_doc/"+docID, bytes.NewBuffer(jsonValue))
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Request failed with status:", resp.Status)
			fmt.Println("Server response:", string(bodyBytes))
		}
	}
}

func (zs *ZsConfig) AppInit() {
	zs.deleteIndex("article")
	indexConfig := IndexConfig{
		Name:        "article",
		StorageType: "disk",
		ShardNum:    1,
		Mappings: map[string]interface{}{
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"content": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"status": map[string]interface{}{
					"type":         "keyword",
					"index":        true,
					"sortable":     true,
					"aggregatable": true,
				},
				"publish_date": map[string]interface{}{
					"type":         "date",
					"format":       "2006-01-02T15:04:05Z07:00",
					"index":        true,
					"sortable":     true,
					"aggregatable": true,
				},
			},
		},
	}

	zs.createOrUpdateIndex(indexConfig)

}

var syncMutex sync.Mutex

func (zs *ZsConfig) SyncGORMToIndexProtected(db *gorm.DB) {
	syncMutex.Lock()
	defer syncMutex.Unlock()

	zs.SyncGORMToIndex(db)
}

func (zs *ZsConfig) SyncGORMToIndexFull(db *gorm.DB) {
	var articles []model.ArticleArticle
	result := db.Select("articleid,articlename,author,intro,postdate").Find(&articles)

	if result.Error != nil {
		fmt.Println("Error querying articles:", result.Error)
		return
	}

	for _, article := range articles {
		doc := map[string]interface{}{
			"name":   article.Articlename,
			"id":     article.Articleid,
			"author": article.Author,
			"intro":  article.Intro,
			"cover":  utils.GetImgURL(article.Articleid),
		}

		zs.indexDocument(fmt.Sprint(article.Articleid), doc)
	}

	// Update the last sync time after a successful sync.
	if len(articles) > 0 {
		zs.setLastSync(int64(articles[len(articles)-1].Postdate))
	}
}

func (zs *ZsConfig) SyncGORMToIndex(db *gorm.DB) {
	lastSync := zs.getLastSync()

	var articles []model.ArticleArticle
	result := db.Select("articleid,articlename,author,intro,postdate").Where("postdate > ?", lastSync).Find(&articles)

	if result.Error != nil {
		fmt.Println("Error querying articles:", result.Error)
		return
	}

	for _, article := range articles {
		doc := map[string]interface{}{
			"name":   article.Articlename,
			"id":     article.Articleid,
			"author": article.Author,
			"intro":  article.Intro,
			"cover":  utils.GetImgURL(article.Articleid),
		}

		zs.indexDocument(fmt.Sprint(article.Articleid), doc)
	}

	// Update the last sync time after a successful sync.
	zs.setLastSync(time.Now().Unix())
}
