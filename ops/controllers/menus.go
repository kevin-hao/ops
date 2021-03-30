package controllers

import (
	"ops/base/controllers/base"
	"ops/base/response"
	"ops/services"
	"ops/utils"

	"github.com/astaxie/beego"
)

type MenusController struct {
	base.BaseController
}

func (c *MenusController) GetAllMenus() {
	if c.Ctx.Input.IsGet() {
		token := c.Ctx.Request.Header["Token"]

		KEY := beego.AppConfig.DefaultString("token::Key", "OPS")
		if len(token) == 0 {
			c.Data["json"] = response.NotAcceptable
			c.ServeJSON()
		}
		claims, err := utils.ParseToken(token[0], []byte(KEY))
		if err != nil {
			c.Data["json"] = response.InvalidToken
			c.ServeJSON()
		}
		username := claims["username"]
		if user := services.UserService.GetByName(username.(string)); user != nil {
			res := services.MenuService.GetAllMenus()
			c.Data["json"] = map[string]interface{}{"code": 200, "message": "获取路由成功", "data": res}
			c.ServeJSON()
		}
		c.Data["json"] = response.InvalidToken
		c.ServeJSON()

	}
	c.Data["json"] = response.BadResquest
	c.ServeJSON()

}

func (c *MenusController) GetMenuList() {
	if c.Ctx.Input.IsGet() {
		token := c.Ctx.Request.Header["Token"]

		KEY := beego.AppConfig.DefaultString("token::Key", "OPS")
		if len(token) == 0 {
			c.Data["json"] = response.NotAcceptable
			c.ServeJSON()
		}
		claims, err := utils.ParseToken(token[0], []byte(KEY))
		if err != nil {
			c.Data["json"] = response.InvalidToken
			c.ServeJSON()
		}
		username := claims["username"]

		id, _ := c.GetInt("id")
		if user := services.UserService.GetByName(username.(string)); user != nil {
			if menus := services.MenuService.GetMenuList(id); menus != nil {
				c.Data["json"] = response.NewJsonResponse(200, "获取路由成功", menus)
				c.ServeJSON()
			}
		}

		c.Data["json"] = response.InvalidToken
		c.ServeJSON()
	}
	c.Data["json"] = response.BadResquest
	c.ServeJSON()
}
