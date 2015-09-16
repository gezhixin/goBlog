package models

import (
	"github.com/astaxie/beego/orm"
)

type Day struct {
	Id          int64
	String      string `orm:"index"`
	Int         int
	MonthString string `orm:"index"`
	MonthInt    int
}

func init() {
	orm.RegisterModel(new(Day))
}

func (this *Day) TableName() string {
	return "tbl_day"
}

func (m *Day) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Day) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Day) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Day) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Day) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
