package api

import (
	"fmt"
	"specialTady/models"
	"specialTady/util"
	"strings"
)

type UserController struct {
	baseController
}

func (this *UserController) Login() {
	var (
		OutJson OutStruct
	)
	account := strings.TrimSpace(this.GetString("u"))
	password := strings.TrimSpace(this.GetString("p"))
	fmt.Println("account: " + account + "  password: " + password)
	if account != "" && password != "" {
		var user models.User
		user.Name = account
		if user.Read("name") != nil || user.Password != util.Md5([]byte(password)) {
			OutJson.ErrorCode = 1000
			OutJson.ErrorMsg = "帐号或密码错误"
		} else if user.Active == 0 {
			OutJson.ErrorCode = 1000
			OutJson.ErrorMsg = "该帐号未激活"
		} else {
			user.LoginCount += 1
			user.Update()
			OutJson.Content = &user
		}
	} else {
		OutJson.ErrorCode = 1001
		OutJson.ErrorMsg = "账号或者密码不能为空"
	}
	this.Data["json"] = &OutJson
	this.ServeJson()
}

func (this *UserController) Logout() {
	this.Data["json"] = "logout"
	this.ServeJson()
}
