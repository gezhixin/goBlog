package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"specialTady/models"
	"specialTady/util"
)

type TagController struct {
	baseController
}

//标签列表
func (this *TagController) List() {
	var page int
	var pagesize int = 10
	var list []*models.Tag
	var tag models.Tag

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := tag.Query().Count()
	if count > 0 {
		tag.Query().OrderBy("-count").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize, "/admin/tag", true).ToString()
	this.display("tag/list")
}

func (this *TagController) Delete() {
	var tagId, _ = this.GetInt64("id")
	var tag = models.Tag{Id: tagId}
	tag.Read()
	var deleteTagName = tag.Name
	if deleteTagName == "未分组" {
		this.StopRun()
	}
	tag.Delete()

	o := orm.NewOrm()
	o.Raw("UPDATE tbl_article SET tag = ? WHERE tag = ?", "未分组", deleteTagName).Exec()
	var count int64 = 0
	orm.NewOrm().Raw("SELECT COUNT(*) FROM tbl_article WHERE tag = ?", "未分组").QueryRow(&count)

	noGroupTag := models.Tag{Name: "未分组"}
	noGroupTag.Read("name")
	noGroupTag.Count = count
	noGroupTag.Update("count")

	this.Redirect("/admin/tag/list", 301)
}

func (this *TagController) Add() {
	var tag名字 = this.GetString("newTagName")
	fmt.Println(tag名字)
	var tag = models.Tag{Name: tag名字}
	tag.Insert()
	this.Redirect("/admin/tag/list", 301)
}
