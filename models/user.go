package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int64
	Name       string `orm:"size(50);unique"`
	Email      string `orm:"size(50);unique"`
	Sex        int64
	Status     int
	Tag        string
	Password   string
	LoginCount int64
	LastIp     string
	Active     int8
	ImgUrl     string
	LastLogin  time.Time `orm:"type(datetime)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *User) TableName() string {
	return "tbl_user"
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
