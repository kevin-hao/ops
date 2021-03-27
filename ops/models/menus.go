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

func (m *Menus) TypeText() string {
	switch m.Type {
	case 1:
		return "目录"
	case 2:
		return "菜单"
	case 3:
		return "按钮"
	}
	return "未知"
}

func init() {
	orm.RegisterModel(new(Menus))
}
