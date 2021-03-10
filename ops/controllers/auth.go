package controllers

import (
	"ops/base/controllers/base"
	"ops/services"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	if c.Ctx.Input.IsPost() {
		username := c.GetString("username")
		password := c.GetString("password")
		user := services.UserService.GetByName(username)
		if user == nil {
			c.Ctx.WriteString("用户不存在")
		} else if user.ValidPassword(password) {
			c.Ctx.WriteString("hello")
		} else {
			c.Ctx.WriteString("用户名密码错误")
		}
	} else {
		c.Ctx.WriteString("wrong method")
	}

}
