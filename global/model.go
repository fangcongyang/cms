package global

import (
	"time"
)

type GVA_MODEL struct {
	Id         int64      `json:"id,string" form:"id" gorm:"primarykey; comment:名称"`
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间"`
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间"`
	TokenId    string     `json:"tokenId" form:"tokenId" gorm:"-"`
}
