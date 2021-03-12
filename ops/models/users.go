package models

import (
	"ops/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID         int        `json:"id" orm:"column(id)"`
	StaffID    string     `json:"staff_id" orm:"column(stafff_id);size(32)"`
	Name       string     `json:"name" orm:"szie(64)"`
	Nickname   string     `json:"nickname" orm:"size(64)"`
	Password   string     `json:"_" orm:"size(1024)"`
	Gender     int        `json:"gender" orm:""`
	Tel        string     `json:"tel" orm:"size(32)"`
	Addr       string     `json:"addr" orm:"size(128)"`
	Email      string     `json:"email" orm:"size(32)"`
	Department string     `json:"department" orm:"size(128)"`
	Status     int        `json:"status" orm:""`
	CreatedAt  *time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt  *time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt  *time.Time `json:"deleted_at" orm:"null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}

func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "在职"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}

func QueryUser(q string) []*User {
	var users []*User
	queryset := orm.NewOrm().QueryTable(&User{})
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("department__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&users)
	queryset.Exclude("password", true)
	return users
}

func init() {
	orm.RegisterModel(new(User))
}
