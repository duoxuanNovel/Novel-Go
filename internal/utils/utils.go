package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CalculateOffset 计算偏移量
func CalculateOffset(page, size int) int {
	return (page - 1) * size
}

const (
	remoteImgURL = "/files/article/image"
	remoteTXTURL = "/files/article/txt"
	SiteURL      = "/"
)

func GetImgURL(aid int) string {
	remoteImg := remoteImgURL + "/" + strconv.Itoa(aid/1000) + "/" + strconv.Itoa(aid) + "/" + strconv.Itoa(aid) + "s.jpg"
	return remoteImg
}

func nocoverURL() string {
	return SiteURL + "/img/nocover.jpg"
}

// GetCateName 根据分类ID 来获取分类名称
func GetCateName(cateId int) string {
	switch cateId {
	case 1:
		return "玄幻"
	case 2:
		return "武侠"
	case 3:
		return "都市"
	case 4:
		return "历史"
	case 5:
		return "网游"
	case 6:
		return "科幻"
	case 7:
		return "其他"
	default:
		return "其他"
	}
}

func GetChapterTableName(id int) string {
	tableName := "shipsay_article_chapter_" + strconv.Itoa((id-1)/10000+1)
	return tableName
}

// GetTxtURL 根据ID和CID 拼接出txt地址
func GetTxtURL(id, cid int) string {
	remoteTxt := remoteTXTURL + "/" + strconv.Itoa(id/1000) + "/" + strconv.Itoa(id) + "/" + strconv.Itoa(cid) + ".txt"
	return remoteTxt
}

// GetTxtContent 从远程.txt网址获取文本内容回来
func GetTxtContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http.Get error:", err)
		return "章节内容缺失或章节不存在！请稍后重新尝试！"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll error:", err)
		return "章节内容缺失或章节不存在！请稍后重新尝试！"
	}

	return string(body)
}

// ConvertToParagraph 将文本转换成段落
func ConvertToParagraph(text string) string {
	lines := strings.Split(text, "\n")

	var paragraphs []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		paragraphs = append(paragraphs, fmt.Sprintf("<p>%s</p>", line))
	}

	return strings.Join(paragraphs, "")
}
