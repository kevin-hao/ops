package forms

type JenkinsForm struct {
	Name string  `form:"name"`
	JenkinsTemplate int `form:"template_id"`
}