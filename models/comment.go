package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id           int64
	Content      string
	Status       int
	ArticleId    int64 `orm:"index"`
	Tag          string
	AuthorId     int64 `orm:"index"`
	AuthorName   string
	SupportCount int64
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

func (this *Comment) TableName() string {
	return "tbl_comment"
}
