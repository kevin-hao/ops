package forms

type JenkinsTemplateForm struct{
	ID  int `form:"id"`
	Content string `form:"content"`
}