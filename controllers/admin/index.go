package admin

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username

	this.TplNames = this.moduleName + "/index/index.html"
}
