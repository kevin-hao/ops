package controllers

import (
	"encoding/json"
	"ops/base/controllers/base"
	"ops/base/response"
	"ops/forms"
	"ops/services"
	"ops/utils"
)

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	form := &forms.LoginForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			form_data := string(c.Ctx.Input.RequestBody)
			data := make(map[string]string)
			err := json.Unmarshal([]byte(form_data), &data)
			if err == nil {
				form.Name = data["username"]
				form.Password = data["password"]
			}
			if user := services.UserService.GetByName(form.Name); user == nil {
				c.Data["json"] = map[string]interface{}{"code": 403, "message": "no user", "data": nil}
				c.ServeJSON()
			} else if user.ValidPassword(form.Password) {
				token := utils.GenerateToken(form.Name)
				// data := map[string]interface{}{"token": token}
				c.Data["json"] = map[string]interface{}{"code": 200, "message": "success", "token": token}
				c.ServeJSON()
			}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 400, "message": "failed", "data": nil}
			c.ServeJSON()
		}
	}
	c.Data["json"] = response.BadResquest
	c.ServeJSON()

}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Data["json"] = response.NewJsonResponse(200, "logout", nil)
	c.ServeJSON()
}
