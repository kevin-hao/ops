package controllers

import (
	"fmt"
	"ops/base/controllers/base"
	"ops/services"

	"github.com/astaxie/beego"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	fmt.Println("sessionKey: ", sessionKey)
	sessionUser := c.GetSession(sessionKey)
	fmt.Println("sessionUser: ", sessionUser)
	// if sessionUser != nil {
	// 	fmt.Println(sessionUser)
	// 	// action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
	// 	// c.Redirect(beego.URLFor(action), http.StatusFound)
	// 	// return
	// }
	if c.Ctx.Input.IsPost() {
		username := c.GetString("username")
		password := c.GetString("password")

		user := services.UserService.GetByName(username)
		if user == nil {
			//c.Data["json"] = map[string]interface{}{"code": 1, "msg": "username not exist"}
			c.Data["json"] = map[string]interface{}{"code": 1}
			c.ServeJSON()

		} else if user.ValidPassword(password) {
			//c.Data["json"] = map[string]interface{}{"code": 2, "msg": user}
			sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
			c.SetSession(sessionKey, user.ID)
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

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Data["json"] = map[string]interface{}{"code": 0}
	c.ServeJSON()
}
