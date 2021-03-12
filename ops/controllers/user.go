package controllers

import (
	"ops/base/controllers/base"
	"ops/models"
)

type UserController struct {
	base.BaseController
}

func (u *UserController) Query() {
	q := u.GetString("q")
	users := models.QueryUser(q)
	u.Data["json"] = map[string]interface{}{"code": 0, "msg": users}
	u.ServeJSON()
}
