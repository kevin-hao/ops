package controllers

import (
	"ops/base/controllers/base"
	"ops/services"
)

type MenusController struct {
	base.BaseController
}

func (c *MenusController) GetAllMenus() {
	res := services.MenuService.GetAllMenus()
	if res == nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "data": nil}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{"code": 0, "data": res}
	c.ServeJSON()

}
