package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Collection struct {
	Id         int64
	UserId     int64 `orm:"index"`
	ArticleId  int64 `orm:"index"`
	Status     int   `orm:"index"`
	Tag        string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(Collection))
}

func (this *Collection) TableName() string {
	return "tbl_collection"
}
