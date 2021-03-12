package controllers

import (
	"fmt"
	"ops/base/controllers/base"
	"ops/services"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	if c.Ctx.Input.IsGet() {
		username := c.GetString("username")
		password := c.GetString("password")
		fmt.Println(username, password)
		user := services.UserService.GetByName(username)
		if user == nil {
			//c.Data["json"] = map[string]interface{}{"code": 1, "msg": "username not exist"}
			c.Data["json"] = map[string]interface{}{"code": 1}
			c.ServeJSON()

		} else if user.ValidPassword(password) {
			//c.Data["json"] = map[string]interface{}{"code": 2, "msg": user}
			c.Data["json"] = map[string]interface{}{"code": 0}
			c.ServeJSON()

		} else {

			//c.Data["json"] = map[string]interface{}{"code": 2, "msg": "username or password is wrong"}
			c.Data["json"] = map[string]interface{}{"code": 2}
			c.ServeJSON()
		}
	} else {
		//c.Data["json"] = map[string]interface{}{"code": 3, "msg": "bad request method"}
		c.Data["json"] = map[string]interface{}{"code": 3}
		c.ServeJSON()

	}

}
