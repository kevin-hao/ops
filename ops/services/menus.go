package services

import (
	// "ops/models"

	"ops/models"

	"github.com/astaxie/beego/orm"
)

type menuService struct {
}

type Meta struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
}

type Menu struct {
	Id         int64  `json:"id"`
	Pid        int64  `json:"pid"`
	Hidden     bool   `json:"hidden"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	AlwaysShow bool   `json:"alwaysShow"`
	Redirect   string `json:"redirect"`
	Name       string `json:"name"`
	Meta       Meta   `json:"meta"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Children   []Menu `json:"children"`
}

func (c *menuService) getAllParentMenu() []*models.Menus {
	var p_menus []*models.Menus

	ormer := orm.NewOrm()
	if _, err := ormer.QueryTable("menus").Filter("pid", "0").All(&p_menus); err == nil {
		return p_menus
	}
	return nil

}

func (c *menuService) getAllMenu() []*models.Menus {
	var menus []*models.Menus
	ormer := orm.NewOrm()
	if _, err := ormer.QueryTable("menus").All(&menus); err == nil {
		return menus
	}
	return nil

}

func (c *menuService) hasChild(allMenus []*models.Menus, parentMenu *models.Menus) (childs []*models.Menus, yes bool) {
	for _, child := range allMenus {
		if child.Pid == parentMenu.ID {
			childs = append(childs, child)
		}
	}

	if childs != nil {
		yes = true
	}
	return
}

func (c *menuService) genMenu(parent *models.Menus) []map[string]interface{} {
	// var Menus []Menu
	var allMenu []*models.Menus

	var t_allMenus []map[string]interface{}
	var t_menus map[string]interface{}
	var hidden bool
	var alwaysShow bool
	var noCache bool

	if allMenus := c.getAllMenu(); allMenus != nil {
		allMenu = allMenus
	}

	if parent.Hidden == "1" {
		hidden = false
		
	} else {
		hidden = true
		
	}
	
	if parent.AlwaysShow == "" {
		alwaysShow = false
	} else {
		alwaysShow = true
	}
	if parent.NoCache == "0" {
		noCache = true
	} else {
		noCache = false
	}

	t_menus = map[string]interface{}{"id": parent.ID, "pid": parent.Pid, "hidden": hidden, "type": parent.TypeText(), "path": parent.Path, "component": parent.Component, "alwaysShow": alwaysShow, "redirect": parent.Redirect, "name": parent.Name, "created_at": parent.CreatedAt, "updated_at": parent.UpdatedAt, "meta": map[string]interface{}{"icon": parent.Icon, "title": parent.Title, "noCache": noCache}}
	

	t_allMenus = append(t_allMenus, t_menus)

	childs, _ := c.hasChild(allMenu, parent)

	if childs != nil {
		var children []map[string]interface{}
		for _, child := range childs {
			

			var c_menus map[string]interface{}
			var c_hidden bool
			var c_alwaysShow bool
			var c_noCache bool

			if child.Hidden == "1" {
				c_hidden = false
			} else {
				c_hidden = true
				
			}
			
			if child.AlwaysShow == "" {
				c_alwaysShow = false
			} else {
				c_alwaysShow = true
			}
			if child.NoCache == "0" {
				c_noCache = true
			} else {
				c_noCache = false
			}
			c_menus = map[string]interface{}{"id": child.ID, "pid": child.Pid, "hidden": c_hidden, "type": child.TypeText(), "path": child.Path, "component": child.Component, "alwaysShow": c_alwaysShow, "redirect": child.Redirect, "name": child.Name, "created_at": child.CreatedAt, "updated_at": child.UpdatedAt, "meta": map[string]interface{}{"icon": child.Icon, "title": child.Title, "noCache": c_noCache}}

			children = append(children, c_menus)
			_, has := c.hasChild(allMenu, child)
			if has {
				c.genMenu(child)
			}
		}
		if len(children) > 0 {
			t_menus["children"] = children
		}

	}
	t_allMenus = append(t_allMenus, t_menus)

	if t_allMenus != nil {
		return t_allMenus
	}
	return nil
}

func (c *menuService) GetAllMenus() []map[string]interface{} {
	var Menus []map[string]interface{}
	var allParentMenus []*models.Menus
	if parentMenus := c.getAllParentMenu(); parentMenus != nil {
		allParentMenus = parentMenus
	}
	for _, parent := range allParentMenus {
		menu := c.genMenu(parent)
		Menus = append(Menus, menu[0])

	}

	return Menus

}

var MenuService = new(menuService)
