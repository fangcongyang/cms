package model

import (
	"cms/global"
)

type SysRole struct {
	global.GVA_MODEL
	Name     		string        	`json:"name" gorm:"comment:角色名称"`
	Type        	int 			`json:"type" gorm:"comment:0功能权限，1其他"`
	Status    		int  			`json:"status" gorm:"comment:状态：0正常，1删除"`
	ParentId   		int64         	`json:"parentId,string" gorm:"comment:父级ID"`
	SearchKey   	string         	`json:"searchKey" gorm:"comment:搜索,0是顶级"`
	Permission   	string         	`json:"permission" gorm:"comment:权限"`
	Sort   			int         	`json:"sort" gorm:"comment:排序，默认倒序"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

type SysRoleIdPermissionVo struct {
	RoleId 			string
	Permission   	string
}