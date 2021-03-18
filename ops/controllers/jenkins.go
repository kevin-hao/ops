package controllers

import (
	"fmt"
	"ops/base/controllers/base"
	"ops/services"
	"ops/forms"

)

type JenkinsController struct {
	base.BaseController
}

func (c *JenkinsController) CreateJob() {
	form := &forms.JenkinsForm{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil{
			
			if res := services.JenkinsService.GetJobByName(form.Name); res != nil {
				c.Data["json"] = map[string]interface{}{"code": 0, "msg": fmt.Sprintf("job %s 已存在", form.Name)}
				c.ServeJSON()
			} else {
				jenkins_job := services.JenkinsService.GetJob(form.Name)
				if jenkins_job != nil {
					c.Data["json"] = map[string]interface{}{"code": 3, "msg": fmt.Sprintf("job %s 已存在于Jenkins", form.Name)}
					c.ServeJSON()
				} else {
					if job := services.JenkinsService.CreateJob(form); job == nil {
						c.Data["json"] = map[string]interface{}{"code": 3, "msg": "create failed"}
						c.ServeJSON()
					} else {
						c.Data["json"] = map[string]interface{}{"code": 0, "msg": fmt.Sprintf("job %s 创建成功", form.Name)}
						c.ServeJSON()
					}
				}
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "msg": "bad request method"}
	c.ServeJSON()
}

func (c *JenkinsController) DeleteJob(){
	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("id")
		services.JenkinsService.DeleteJob(id)
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除成功"}
	    c.ServeJSON()
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "msg": "bad request method"}
	c.ServeJSON()
}

type JenkinsTemplateController struct{
	base.BaseController
}

func (c *JenkinsTemplateController) CreateTemplate(){
	if c.Ctx.Input.IsPost() {
		name := c.GetString("name")
		content := c.GetString("content")
		if err := services.JenkinsTemplateService.GetTemplateByName(name); err == nil {
			res := services.JenkinsTemplateService.CreateTemplate(name, content)
			if res != nil{
				c.Data["json"] = map[string]interface{}{"code": 0, "msg": fmt.Sprintf("%s 创建成功", name)}
		    	c.ServeJSON()
			}
			c.Data["json"] = map[string]interface{}{"code": 3, "msg": fmt.Sprintf("%s 创建失败, %s", name, res)}
		    c.ServeJSON()
		}
		c.Data["json"] = map[string]interface{}{"code": 2, "msg": fmt.Sprintf("%s 模版已存在", name)}
		c.ServeJSON()
		
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "msg": "bad request method"}
	c.ServeJSON()
}

func (c *JenkinsTemplateController) Modify(){
	form := &forms.JenkinsTemplateForm{}
	if c.Ctx.Input.IsPost(){
		if err := c.ParseForm(form); err == nil{
			services.JenkinsTemplateService.Modify(form)
			c.Data["json"] = map[string]interface{}{"code": 0, "msg": "update"}
	    	c.ServeJSON()
		}
		
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "failed"}
		c.ServeJSON()
		
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "msg": "bad request method"}
	c.ServeJSON()
}
