package request

import "cms/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
