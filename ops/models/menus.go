package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Menus struct {
	ID         int        `json:"id" orm:"column(id)"`
	Pid        int        `json:"pid" orm:"column(pid)"`
	Type       int        `json:"type" orm:"column(type)"`
	Icon       string     `json:"icon" orm:"size(128)"`
	Title      string     `json:"title" orm:"size(128)"`
	Path       string     `json:"path" orm:"size(128)"`
	Component  string     `json:"component" orm:"size(128)"`
	AlwaysShow string     `json:"alwaysShow" orm:"size(128)"`
	Redirect   string     `json:"redirect" orm:"size(128)"`
	NoCache    string     `json:"noCache" orm:"size(128)"`
	Name       string     `json:"name" orm:"size(128)"`
	Hidden     string     `json:"hidden" orm:"size(128)"`
	CreatedAt  *time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt  *time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt  *time.Time `json:"deleted_at" orm:"null"`
}

func init() {
	orm.RegisterModel(new(Menus))
}
