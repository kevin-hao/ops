package forms

type LoginForm struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}
