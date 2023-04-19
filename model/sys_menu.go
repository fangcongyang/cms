package model

import (
	"cms/global"
)

type SysMenu struct {
	global.GVA_MODEL
	Name     		string        	`json:"name" gorm:"comment:菜单名称"`
	Icon   			string        	`json:"icon" gorm:"comment:菜单ICON"`
	Url        		string         	`json:"url" gorm:"comment:菜单跳转地址"`
	AuthId 			string 			`json:"authId" gorm:"comment:授权(多个用逗号分隔，如：user:list,user:create)"`
	Type        	int 			`json:"type" gorm:"comment:0主菜单（目录没连接），1菜单项，3按钮，4其他"`
	Status    		int  			`json:"status" gorm:"comment:状态：0正常，1删除"`
	ParentId   		int64         	`json:"parentId,string" gorm:"comment:父级ID"`
	Sort   			int         	`json:"sort" gorm:"comment:排序，默认倒序"`
	Alias   		string         	`json:"alias" gorm:"comment:alias"`
	SearchKey   	string         	`json:"searchKey" gorm:"comment:搜索,0是顶级"`
	Children        []SysMenu       `json:"children" gorm:"-"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}