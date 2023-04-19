package response

import "cms/model"

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus"`
}

