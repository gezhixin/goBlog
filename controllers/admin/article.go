package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"os"
	"specialTady/models"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	baseController
}

//添加
func (this *ArticleController) Add() {
	var list []*models.Tag
	var tag models.Tag

	var day = this.GetString("d")
	var month = this.GetString("m")

	tag.Query().OrderBy("-count").All(&list)
	this.Data["list"] = list
	this.Data["day"] = day
	this.Data["month"] = month

	this.display()
}

//管理
func (this *ArticleController) List() {
	var (
		list     []*models.Article
		day      string
		month    string
		listType string
		tag      string
		tagList  []*models.Tag
	)

	listType = this.GetString("t")

	if listType == "" || listType == "date" {
		day = this.GetString("d")
		month = this.GetString("m")

		if day == "" || month == "" {
			_, timeMonth, timeDay := this.getTime().Date()
			day = fmt.Sprintf("%d", timeDay)
			month = fmt.Sprintf("%d", timeMonth)
		}

		_, err := orm.NewOrm().Raw("SELECT * FROM tbl_article WHERE status != ? and happen_month = ? and happen_day = ?", 2, month, day).QueryRows(&list)
		if err != nil {
			fmt.Println(err)
		}
	} else if listType == "tag" {
		tag = this.GetString("tag")
		_, err := orm.NewOrm().Raw("SELECT * FROM tbl_article WHERE tag = ? and status != ?", tag, 2).QueryRows(&list)
		if err != nil {
			fmt.Println(err)
		}
		_, errTag := orm.NewOrm().Raw("SELECT * FROM tbl_tag").QueryRows(&tagList)
		if errTag != nil {
			fmt.Println(errTag)
		}
	}

	this.Data["list"] = list
	this.Data["happenDay"] = day
	this.Data["happenMonth"] = month
	this.Data["tag"] = tag
	this.Data["listType"] = listType
	this.Data["tagList"] = tagList
	this.display()
}

func (this *ArticleController) Trash() {
	var (
		list []*models.Article
	)
	_, err := orm.NewOrm().Raw("SELECT * FROM tbl_article WHERE status = ?", 2).QueryRows(&list)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["list"] = list
	this.display()
}

func (this *ArticleController) Destroy() {
	var (
		id      int64
		article models.Article
	)

	id, _ = this.GetInt64("id")
	article.Id = id

	if err := article.Read(); err != nil {
		fmt.Println(err)
	}
	article.Delete()

	this.Redirect(this.Ctx.Request.Referer(), 302)
}

//编辑
func (this *ArticleController) Edit() {
	var (
		list []*models.Tag
		tag  models.Tag
	)

	id, _ := this.GetInt64("id")
	post := models.Article{Id: id}
	if post.Read() != nil {
		this.Abort("404")
	}

	tag.Query().OrderBy("-count").All(&list)
	fmt.Println(list)

	this.Data["list"] = list
	this.Data["post"] = post
	this.display()
}

//保存
func (this *ArticleController) Save() {
	var (
		id          int64  = 0
		title       string = strings.TrimSpace(this.GetString("title"))
		content     string = this.GetString("content")
		status      int    = 0
		happenMonth string = strings.TrimSpace(this.GetString("happenMonth"))
		happenDay   string = strings.TrimSpace(this.GetString("happenDay"))
		article     models.Article
		tag         models.Tag
	)

	if title == "" {
		this.showmsg("标题不能为空！")
	}

	id, _ = this.GetInt64("id")
	status, _ = this.GetInt("status")

	tagName := this.GetString("tag")
	fmt.Println(tagName)
	tag.Name = tagName
	tag.Read("name")

	if status != 1 {
		status = 0
	}

	if id < 1 {
		article.AuthorId = this.userid
		article.AuthorName = this.username
		article.Tag = tagName
		err := article.Insert()
		if err != nil {
			fmt.Println(err)
		}

		tag.Count = tag.Count + 1
		tag.Update()
	} else {
		article.Id = id
		if article.Read() != nil {
			goto RD
		}
	}

	if article.Tag != tagName {
		desTag := models.Tag{Name: article.Tag}
		desTag.Read("name")
		desTag.Count--
		desTag.Update("count")

		tag.Count = tag.Count + 1
		tag.Update()

		article.Tag = tagName
	}

	article.Status = status
	article.Title = title
	article.Content = content
	article.HappenMonth = happenMonth
	article.HappenDay = happenDay
	article.Update("tag", "status", "title", "content", "happen_month", "happen_day")

RD:
	url := "/admin/article/list?t=date&m=" + happenMonth + "&d=" + happenDay
	this.Redirect(url, 302)
}

//delete
func (this *ArticleController) Delete() {
	var (
		id      int64
		article models.Article
	)

	id, _ = this.GetInt64("id")
	article.Id = id

	if err := article.Read(); err == nil {
		fmt.Println(err)
	}

	article.Status = 2
	article.Update("status")

	tag := models.Tag{Name: article.Tag}
	tag.Read("Name")
	tag.Count--
	tag.Update("count")

	this.Redirect(this.Ctx.Request.Referer(), 302)
}

//批处理
func (this *ArticleController) Batch() {
	ids := this.GetStrings("ids[]")
	op := this.GetString("op")

	if len(ids) == 0 {
		this.Redirect(this.Ctx.Request.Referer(), 302)
		this.StopRun()
	}

	idarr := make([]int64, 0)
	for _, v := range ids {
		if id, _ := strconv.ParseInt(v, 10, 64); id > 0 {
			idarr = append(idarr, id)
		}
	}

	var post models.Article
	var articleList []*models.Article

	switch op {
	case "topub": //移到已发布
		post.Query().Filter("id__in", idarr).All(&articleList)
		for _, v := range articleList {
			if v.Status == 2 {
				tag := models.Tag{Name: v.Tag}
				tag.Read("name")
				tag.Count++
				tag.Update("count")
			}
		}
		post.Query().Filter("id__in", idarr).Update(orm.Params{"status": 0})

	case "todrafts": //移到草稿箱
		post.Query().Filter("id__in", idarr).All(&articleList)
		for _, v := range articleList {
			if v.Status == 2 {
				tag := models.Tag{Name: v.Tag}
				tag.Read("name")
				tag.Count++
				tag.Update("count")
			}
		}
		post.Query().Filter("id__in", idarr).Update(orm.Params{"status": 1})

	case "totrash": //移到回收站
		post.Query().Filter("id__in", idarr).All(&articleList)
		for _, v := range articleList {
			if v.Status != 2 {
				tag := models.Tag{Name: v.Tag}
				tag.Read("name")
				tag.Count--
				tag.Update("count")
			}
		}
		post.Query().Filter("id__in", idarr).Update(orm.Params{"status": 2})
	case "todestroy":
		post.Query().Filter("id__in", idarr).Delete()
	}

	this.Redirect(this.Ctx.Request.Referer(), 302)
}

//上传文件
func (this *ArticleController) Upload() {
	_, header, err := this.GetFile("upfile")
	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	out := make(map[string]string)
	out["url"] = ""
	out["fileType"] = ext
	out["original"] = header.Filename
	out["state"] = "SUCCESS"
	if err != nil {
		out["state"] = err.Error()
	} else {
		savepath := "./static/upload/"
		if err := os.MkdirAll(savepath, os.ModePerm); err != nil {
			out["state"] = err.Error()
		} else {
			filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
			if err := this.SaveToFile("upfile", filename); err != nil {
				out["state"] = err.Error()
			} else {
				out["url"] = filename[1:]
			}
		}
	}

	this.Data["json"] = out
	this.ServeJson()
}
