package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"specialTady/models"
	"specialTady/routers"
	"specialTady/util"
)

func init() {
	var config_file string
	flag.StringVar(&config_file, "conf", "", "the path of the config file")
	flag.Parse()
	if config_file != "" {
		beego.AppConfigPath, _ = filepath.Abs(config_file)
		beego.ParseConfig()
	} else {
		if config_file = os.Getenv("BEEGO_APP_CONFIG_FILE"); config_file != "" {
			beego.AppConfigPath, _ = filepath.Abs(config_file)
			beego.ParseConfig()
		}
	}

	util.Factory.Set("cache", func() (interface{}, error) {
		mc := util.NewLruCache(1000)
		return mc, nil
	})

	models.Init()
	routers.Init()
	models.CreateAdmin()
	// models.CreateDate()
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = false
	}

	beego.Run()
}
