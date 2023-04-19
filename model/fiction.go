package model

import (
	"cms/global"
)

type Fiction struct {
	global.GVA_MODEL
	Name      string  `json:"name" gorm:"comment:名称"`
	Logo      string  `json:"logo"  gorm:"comment:图标"`
	Status    int     `json:"status" gorm:"comment:状态 0禁用 1启用" `
	Score     float64 `json:"score" gorm:"default:null;comment:评分 格式：9.3"`
	Author    string  `json:"author" gorm:"default:'';comment:作者"`
	Introduce string  `json:"introduce" gorm:"default:888;comment:小说简介"`
}

func (Fiction) TableName() string {
	return "fic_fiction"
}

type FictionChapter struct {
	global.GVA_MODEL
	Name      string `json:"name" gorm:"comment:名称"`
	FictionId int64  `json:"fictionId"  gorm:"comment:小说id"`
	Sort      int    `json:"sort" gorm:"comment:章节排序" `
	WordCount int    `json:"wordCount" gorm:"comment:章节字数"`
	FilePath  string `json:"filePath" gorm:"comment:章节保存路径"`
	Context   string `json:"context" gorm:"-"`
}

func (FictionChapter) TableName() string {
	return "fic_fiction_chapter"
}

type FictionFictionSort struct {
	global.GVA_MODEL
	FictionId     int64 `json:"fiction_id" gorm:"comment:小说id"`
	FictionSortId int64 `json:"status"  gorm:"comment:小说分类id"`
}

func (FictionFictionSort) TableName() string {
	return "fic_fiction_fiction_sort"
}

type FictionSort struct {
	global.GVA_MODEL
	Logo      string `json:"logo" gorm:"comment:logo图标"`
	Name      string `json:"name" gorm:"comment:分类名"`
	Status    int    `json:"status"  gorm:"comment:分类状态 0=不可用 1=可用"`
	BookCount int    `json:"bookCount" gorm:"comment:当前分类小说数"`
	Type      string `json:"type" gorm:"comment:分类类型 ‘male’ 男 ‘female’ 女"`
	Sort      int    `json:"sort" gorm:"comment:排序" `
}

func (FictionSort) TableName() string {
	return "fic_fiction_sort"
}
