package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type JenkinsJob struct {
	ID         int        `json:"id" orm:"column(id)"`
	Name       string     `json:"name" orm:"size(64);unique"`
	CreatedAt  *time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt  *time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt  *time.Time `json:"deleted_at" orm:"null"`
	JenkinsTemplate *JenkinsTemplate `orm:"rel(fk)"`
}

func (j *JenkinsJob) TableName() string {
	return "jenkins_jobs"
}

type JenkinsTemplate struct{
	ID         int        `json:"id" orm:"column(id)"`
	Name       string     `json:"name" orm:"size(64);unique"`
	Content    string     `son:"name" orm:"type(text)"`
	CreatedAt  *time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt  *time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt  *time.Time `json:"deleted_at" orm:"null"`
	JenkinsJob []*JenkinsJob `orm:"reverse(many)"`
}

func (j *JenkinsTemplate) TableName() string {
	return "jenkins_templates"
}

func init() {
	orm.RegisterModel(new(JenkinsJob))
	orm.RegisterModel(new(JenkinsTemplate))
}