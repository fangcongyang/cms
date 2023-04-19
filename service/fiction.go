package service

import (
	"cms/global"
	"cms/model"
	"cms/model/request"
	"io/ioutil"
	"os"
)

//@author: fcy
//@function: QueryFictionPageList
//@description: 分页获取小说信息
//@param: id int64
//@return: err error, fiction *model.Fiction

func QueryFictionSortPageList(info request.PageInfo) (list interface{}, err error, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNo - 1)
	db := global.GVA_DB.Model(&model.FictionSort{})
	var fictionSortList []model.FictionSort
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fictionSortList).Error
	return fictionSortList, err, total
}

func GetFictionSortIdByName(name string) (id int64, err error) {
	var localId int64
	var fs model.FictionSort
	e := global.GVA_DB.Where("name = ?", name).First(&fs).Error
	if fs == (model.FictionSort{}) {
		localId = global.GVA_WORKER.GetId()
		fs = model.FictionSort{
			Logo:      "",
			Name:      name,
			Status:    1,
			Type:      "male",
			BookCount: 0,
			Sort:      1,
		}
		fs.Id = localId
		e = global.GVA_DB.Create(&fs).Error
	} else {
		localId = fs.Id
	}
	return localId, e
}

func SaveFictionSortFiction(fictionId int64, fictionSortId int64) (err error) {
	var id int64
	id = global.GVA_WORKER.GetId()
	fs := model.FictionFictionSort{
		FictionId:     fictionId,
		FictionSortId: fictionSortId,
	}
	fs.Id = id
	e := global.GVA_DB.Create(&fs).Error
	return e
}

//@author: fcy
//@function: Get
//@description: 通过id获取小说信息
//@param: id int64
//@return: err error, user *model.SysUser

func GetFictionById(id int64) (fiction *model.Fiction, err error) {
	var f model.Fiction
	err = global.GVA_DB.Where("`id` = ?", id).First(&f).Error
	return &f, err
}

func SaveFiction(fiction model.Fiction) (id int64, err error) {
	id = global.GVA_WORKER.GetId()
	fiction.Id = id
	e := global.GVA_DB.Create(fiction).Error
	return id, e
}

//@author: fcy
//@function: QueryFictionPageList
//@description: 分页获取小说信息
//@param: id int64
//@return: err error, fiction *model.Fiction

func QueryFictionPageList(info request.FictionApiParam) (list interface{}, err error, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNo - 1)
	db := global.GVA_DB.Model(&model.Fiction{}).Where("id IN (SELECT fiction_id FROM fic_fiction_fiction_sort WHERE fiction_sort_id = ?)", info.SortId)
	var fictionList []model.Fiction
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fictionList).Error
	return fictionList, err, total
}

//@author: fcy
//@function: Get
//@description: 通过id获取小说信息
//@param: id int64
//@return: err error, user *model.SysUser

func GetFictionChapterById(id int64) (fiction *model.FictionChapter, err error) {
	var f model.FictionChapter
	err = global.GVA_DB.Where("`id` = ?", id).First(&f).Error
	file, err := os.Open(global.GVA_CONFIG.System.FictionBasePath + f.FilePath)
	if err == nil {
		var data, _ = ioutil.ReadAll(file)
		f.Context = string(data[:])
	}
	defer file.Close()
	return &f, err
}

//@author: fcy
//@function: QueryFictionPageList
//@description: 分页获取小说信息
//@param: id int64
//@return: err error, fiction *model.Fiction

func QueryFictionChapterPageList(fictionId int64) (list interface{}, err error) {
	db := global.GVA_DB.Model(&model.FictionChapter{})
	var fictionChapterList []model.FictionChapter
	err = db.Select("id, sort, name").Where("fiction_id = ?", fictionId).Find(&fictionChapterList).Error
	return fictionChapterList, err
}

func SaveFictionChapter(fictionChapter model.FictionChapter) (err error) {
	id := global.GVA_WORKER.GetId()
	fictionChapter.Id = id
	e := global.GVA_DB.Create(&fictionChapter).Error
	return e
}
