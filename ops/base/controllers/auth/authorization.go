package auth

import (
	"fmt"
	"ops/base/controllers/base"

	"github.com/astaxie/beego"
)

// AuthorizationController 所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	base.BaseController
}

// Prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	fmt.Println("sessionkey: ", sessionKey)
	user := c.GetSession(sessionKey)
	fmt.Println("user: ", user)
	if user == nil {
		c.Data["json"] = map[string]interface{}{"code": "failed"}
		//c.ServeJSON()
		// action := beego.AppConfig.DefaultString("auth::LoginAction","AthController.Login")
		// c.Redirect(beego.URLFor(action), http.StatusFound)
	}
}
