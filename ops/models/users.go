package models

import (
	"ops/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(stafff_id);size(32)"`
	Name       string     `orm:"szie(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(32)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

func init() {
	orm.RegisterModel(new(User))
}
