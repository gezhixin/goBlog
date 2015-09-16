package api

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"specialTady/models"
	// "strings"
)

type ArticleController struct {
	baseController
}

type IndexStruct struct {
	ShortTadayDescr string
	ArticleList     interface{} `json:"articleList"`
}

func (this *ArticleController) Index() {
	var (
		articleList []*models.Article
	)
	var taday = IndexStruct{}
	taday.ShortTadayDescr = "这是一个美好的早晨"
	_, errTag := orm.NewOrm().Raw("SELECT content FROM tbl_article").QueryRows(&articleList)
	if errTag != nil {
		fmt.Println(errTag)
	}
	taday.ArticleList = &articleList
	this.setOutContent(taday)
	this.JsonOutput()
}
