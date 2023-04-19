package model

import (
	"cms/global"
)

type SysUser struct {
	global.GVA_MODEL
	Account    		string       	`json:"account" gorm:"comment:账号"`
	Password    	string       	`json:"-"  gorm:"comment:用户登录密码"`
	Logo    		string       	`json:"logo" gorm:"default:系统用户;comment:用户昵称" `
	Name    		string       	`json:"name" gorm:"comment:用户名称" `
	Permission    	string      	`json:"permission"  gorm:"-"`
	RoleId			string 			`json:"-"  gorm:"-"`
}

func (SysUser) TableName() string {
	return "sys_user"
}