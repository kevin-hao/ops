package controllers

import (
	"ops/base/controllers/base"
	"ops/base/response"
	"ops/models"

	"ops/services"
	"ops/utils"

	"github.com/astaxie/beego"
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

func (u *UserController) GetUserInfo() {
	if u.Ctx.Input.IsPost() {
		token := u.Ctx.Request.Header["Token"][0]
		KEY := beego.AppConfig.DefaultString("token::Key", "OPS")
		claims, err := utils.ParseToken(token, []byte(KEY))
		if err != nil {
			u.Data["json"] = response.InvalidToken
			u.ServeJSON()
		}

		username := claims["username"]
		user := services.UserService.GetByName(username.(string))
		data := map[string]interface{}{"name": user.Name, "email": user.Email, "gender": user.GenderText()}
		u.Data["json"] = response.NewJsonResponse(200, "success", data)
		u.ServeJSON()
	}
	u.Data["json"] = response.BadResquest
	u.ServeJSON()
}
