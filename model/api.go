package model

import (
	"cms/global"
	"time"
)

type UserInfo struct {
	global.GVA_MODEL
	AppKey        string     `json:"appKey" gorm:"comment:app_key"`
	Mobile        string     `json:"mobile"  gorm:"comment:手机号码"`
	Email         string     `json:"email" gorm:"comment:邮箱" `
	NickName      string     `json:"nickName" gorm:"comment:昵称"`
	RealName      string     `json:"realName" gorm:"comment:真实姓名"`
	Logo          string     `json:"logo" gorm:"comment:头像"`
	Sex           int        `json:"sex" gorm:"default:0;comment:0未知，1女，2男"`
	Status        int        `json:"status" gorm:"default:1;comment:账户状态 0 正常 1冻结"`
	Sort          int        `json:"sort" gorm:"comment:排序，倒序"`
	WechatStatus  string     `json:"wechatStatus" gorm:"default:0;comment:微信绑定状态：0-未绑定，1-已绑定"`
	WechatId      string     `json:"wechatId" gorm:"comment:微信绑定id"`
	LoginLastTime *time.Time `json:"loginLastTime" gorm:"comment:最后登录时间"`
	TokenId       string     `json:"tokenId" gorm:"-"`
	//用户类型，0普通用户,1部门管理员
	Type string `json:"type" gorm:"-"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type UserLogin struct {
	global.GVA_MODEL
	AppKey   string `json:"appKey" gorm:"comment:app_key"`
	Mobile   string `json:"mobile"  gorm:"comment:手机号码"`
	Password string `json:"password" gorm:"comment:密码" `
	Salt     string `json:"salt" gorm:"comment:密码盐"`
}

func (UserLogin) TableName() string {
	return "user_login"
}

type SysBanner struct {
	global.GVA_MODEL
	Name           string     `json:"name" gorm:"comment:名称"`
	Logo           string     `json:"logo"  gorm:"comment:banner图片"`
	Type           int        `json:"type" gorm:"default:1;comment:0 启动页图片 1banner图片" `
	SkipType       int        `json:"skipType" gorm:"default:0;comment:跳转类型：0=不跳转，1=跳内页，2=h5,3=请求后台方法，4=电话热线"`
	SkipUrlIos     string     `json:"skipUrlIos" gorm:"comment:ios跳转的路径或方法"`
	SkipUrlAndroid string     `json:"skipUrlAndroid"  gorm:"comment:安卓跳转的路径或方法"`
	Status         int        `json:"status" gorm:"default:0;comment:banner状态:0-未启用，1-正在使用，2-已删除" `
	Sort           int        `json:"sort" gorm:"comment:排序，倒序（值越大越靠前）"`
	BeginTime      *time.Time `json:"beginTime" gorm:"comment:开始时间"`
	EndTime        *time.Time `json:"endTime"  gorm:"comment:结束时间"`
	SysUserId      int64      `json:"sysUserId" gorm:"comment:操作人ID" `
	SysUserName    string     `json:"sysUserName" gorm:"comment:操作人名称"`
}

func (SysBanner) TableName() string {
	return "sys_banner_info"
}

type UserFiction struct {
	global.GVA_MODEL
	UserId    int64 `json:"userId" gorm:"comment:用户id"`
	FictionId int64 `json:"fictionId"  gorm:"comment:小说id"`
}

func (UserFiction) TableName() string {
	return "user_fiction"
}
