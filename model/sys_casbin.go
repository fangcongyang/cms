package model

type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	Permission  string `json:"permission" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}
