package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id              int64
	Title           string `orm:"type(text)";unique`
	Content         string `orm:"type(text)"`
	Status          int    `orm:"index"`
	Tag             string
	AuthorId        int64 `orm:"index"`
	AuthorName      string
	HappenMonth     string `orm:"index"`
	HappenDay       string `orm:"index"`
	ReadCount       int64
	CollectionCount int64
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) TableName() string {
	return "tbl_article"
}

func (m *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
