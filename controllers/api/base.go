package api

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type OutStruct struct {
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	Content   interface{} `json:"content"`
}

type baseController struct {
	beego.Controller
	userid         int
	username       string
	moduleName     string
	controllerName string
	actionName     string
	mobileOs       string
	osVersion      string
	OutJson        OutStruct
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "mobileAPI"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.mobileOs = this.GetString("os")
	this.osVersion = this.GetString("sv")
	fmt.Println("request ----> " + this.Ctx.Request.URL.String() + "  os:" + this.mobileOs + " verson: " + this.osVersion)
}

func (this *baseController) setOutError(errorCode int, errorMsg string) {
	this.OutJson.ErrorCode = errorCode
	this.OutJson.ErrorMsg = errorMsg
}

func (this *baseController) setOutContent(content interface{}) {
	this.OutJson.Content = content
}

func (this *baseController) JsonOutput() {
	this.Data["json"] = this.OutJson
	this.ServeJson()
}
