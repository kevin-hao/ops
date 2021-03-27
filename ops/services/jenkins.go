package services

import (
	"ops/forms"
	"ops/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/bndr/gojenkins"

	"strings"
)

type jenkinsService struct {
}

//创建jenkins对象
func (c *jenkinsService) GetJenkins() *gojenkins.Jenkins {
	JenkinsUrl := beego.AppConfig.DefaultString("jenkins::JenkinsUrl", "JenkinsUrl")
	JenkinsUser := beego.AppConfig.DefaultString("jenkins::JenkinsUser", "JenkinsUser")
	JenkinsPW := beego.AppConfig.DefaultString("jenkins::JenkinsPW", "JenkinsPW")
	jenkins := gojenkins.CreateJenkins(nil, JenkinsUrl, JenkinsUser, JenkinsPW)
	_, err := jenkins.Init()
	if err != nil {
		panic(err)
	}
	return jenkins
}

func (c *jenkinsService) GetJobByName(name string) *models.JenkinsJob {
	jenkin_job := &models.JenkinsJob{Name: name}
	ormer := orm.NewOrm()

	if err := ormer.Read(jenkin_job, "Name"); err == nil {
		return jenkin_job
	}
	return nil

}

func (c *jenkinsService) CreateJob(form *forms.JenkinsForm) *models.JenkinsJob {
	jenkins := c.GetJenkins()
	template := &models.JenkinsTemplate{ID: form.JenkinsTemplate}
	ormer := orm.NewOrm()
	if err := ormer.Read(template); err == nil {
		job := c.GetJob(template.Name)
		conf, _ := job.GetConfig()
		conf = strings.Replace(conf, template.Name, form.Name, -1)
		jenkin_job := &models.JenkinsJob{
			Name:            form.Name,
			JenkinsTemplate: template,
			Content:         conf,
		}
		if _, err := orm.NewOrm().Insert(jenkin_job); err == nil {
			_, err = jenkins.CreateJob(conf, form.Name)
			if err == nil {
				return jenkin_job
			}
		}
	}
	return nil
}

func (c *jenkinsService) getJobById(id int) *models.JenkinsJob {
	ormer := orm.NewOrm()
	job := &models.JenkinsJob{ID: id}
	if err := ormer.Read(job); err == nil {
		return job
	}
	return nil

}

func (c *jenkinsService) DeleteJob(id int) *models.JenkinsJob {
	jenkins := c.GetJenkins()
	ormer := orm.NewOrm()
	job := c.getJobById(id)
	if job != nil {
		id := job.ID
		name := job.Name
		if _, err := ormer.Delete(&models.JenkinsJob{ID: id}); err == nil {
			if _, res := jenkins.DeleteJob(name); res == nil {
				return nil
			}
		}
	}
	return job
}

func (c *jenkinsService) GetJob(name string) *gojenkins.Job {
	jenkins := c.GetJenkins()
	job, err := jenkins.GetJob(name)
	if err == nil {
		return job
	}
	return nil
}

type jenkinsTemplateService struct {
}

func (c *jenkinsTemplateService) CreateTemplate(name, content string) *models.JenkinsTemplate {
	template := &models.JenkinsTemplate{
		Name:    name,
		Content: content,
	}
	if _, err := orm.NewOrm().Insert(template); err == nil {
		return template
	}
	return nil
}

func (c *jenkinsTemplateService) GetTemplateByName(name string) *models.JenkinsTemplate {
	template := &models.JenkinsTemplate{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(template); err != nil {
		return template
	}
	return nil
}

func (c *jenkinsTemplateService) GetById(id int) *models.JenkinsTemplate {
	template := &models.JenkinsTemplate{ID: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(template); err != nil {
		return template
	}
	return nil
}

func (c *jenkinsTemplateService) Modify(form *forms.JenkinsTemplateForm) *models.JenkinsTemplate {

	if template := c.GetById(form.ID); template != nil {

		return template
	}
	return nil
}

var JenkinsService = new(jenkinsService)
var JenkinsTemplateService = new(jenkinsTemplateService)
