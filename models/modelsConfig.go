package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"specialTady/util"
	"strconv"
	"time"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RunSyncdb("default", false, true)
}

func CreateAdmin() {
	admin := User{}
	admin.Name = "admin"
	admin.Id = 1
	admin.Password = util.Md5([]byte("admin"))
	admin.Email = "x_inge@163.com"
	admin.Active = 1
	timezone, _ := strconv.ParseFloat("8", 64)
	add := timezone * float64(time.Hour)
	admin.LastLogin = time.Now().UTC().Add(time.Duration(add))
	err := admin.Insert()
	fmt.Println(err)
}

func CreateDate() {
	var (
		dcount  int
		mstring string
		dstring string
	)

	for i := 1; i <= 12; i++ {
		if i < 10 {
			mstring = fmt.Sprintf("0%d", i)
		} else {
			mstring = fmt.Sprintf("%d", i)
		}
		if i == 2 {
			dcount = 29
		} else if i == 4 || i == 6 || i == 9 || i == 11 {
			dcount = 30
		} else {
			dcount = 31
		}
		for j := 1; j <= dcount; j++ {

			if j < 10 {
				dstring = fmt.Sprintf("0%d", j)
			} else {
				dstring = fmt.Sprintf("%d", j)
			}

			day := Day{}
			day.String = dstring
			day.Int = j
			day.MonthString = mstring
			day.MonthInt = i
			day.Insert()
		}
	}
}
