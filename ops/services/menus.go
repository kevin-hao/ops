package services

import (
	// "ops/models"

	"github.com/astaxie/beego/orm"
)

type menuService struct {
}

//第一种方式
// func (c *menuService) GetAllMenus() []*models.Menus {

// 	ormer := orm.NewOrm()
// 	var menu []*models.Menus
// 	_, err := ormer.QueryTable("menus").RelatedSel().All(&menu)

// 	if err == nil {
// 		return menu
// 	}
// 	return nil
// }

//第二种方式
type Meta struct {
	title   string
	icon    string
	noCache bool
}

type Menu struct {
	id         int64
	pid        int64
	hidden     bool
	Type       string `json:"type"`
	title      string
	path       string
	component  string
	alwaysShow bool
	redirect   string
	name       string
	meta       Meta
	created_at string
	updated_at string
}

func (c *menuService) GetAllMenus() []orm.Params {
	ormer := orm.NewOrm()
	var menus []orm.Params
	//_, err := ormer.Raw("SELECT s.*, p.pid parent  from menus s left join menus p on p.pid=s.id;").Values(&menu)
	_, err := ormer.Raw("SELECT *  from menus;").Values(&menus)
	if err == nil {
		return menus
	}
	return nil
}

var MenuService = new(menuService)
