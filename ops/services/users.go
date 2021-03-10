package services

import (
	"ops/models"

	"github.com/astaxie/beego/orm"
)

type userService struct {
}

func (u *userService) GetByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

var UserService = new(userService)
