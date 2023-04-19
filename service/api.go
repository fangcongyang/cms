package service

import (
	"cms/global"
	"cms/model"
	"cms/model/request"
	"cms/utils"
	"errors"
	"strings"
	"time"
)

//@author: fcy
//@function: loginByPassword
//@description: app登录
//@param: id int64
//@return: err error, fiction *model.Fiction

func LoginByPassword(mobile string, password string) (userInfo model.UserInfo, err error) {
	if mobile == "" || password == "" {
		return model.UserInfo{}, errors.New("账号密码错误")
	}
	db := global.GVA_DB.Model(&model.UserLogin{})
	var userLogin model.UserLogin
	err = db.Where("mobile = ?", mobile).Find(&userLogin).Error
	if userLogin == (model.UserLogin{}) {
		return model.UserInfo{}, errors.New("账号密码错误")
	}
	pwd := utils.MD5V([]byte(strings.ToUpper(password + userLogin.Salt)))
	if !strings.EqualFold(pwd, userLogin.Password) {
		return model.UserInfo{}, errors.New("账号密码错误")
	}
	err = global.GVA_DB.Model(&model.UserInfo{}).Where("mobile = ?", mobile).Set("login_last_time = ?", time.Time{}.Local()).Error
	global.GVA_DB.Model(&model.UserInfo{}).Where("mobile = ?", mobile).Find(&userInfo)
	if userInfo != (model.UserInfo{}) && userInfo.Status != 0 {
		return model.UserInfo{}, errors.New("账号密码错误")
	}
	return userInfo, err
}

func SelectBannerList(bannerType int) (list interface{}, err error) {
	db := global.GVA_DB.Model(&model.SysBanner{})
	var bannerList []model.SysBanner
	err = db.Where("status = ? and type = ?", "1", bannerType).Find(&bannerList).Error
	return bannerList, err
}

func ListBookShelf(userId int64, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNo - 1)
	db := global.GVA_DB.Table("user_fiction as uf")
	var fictionList []model.Fiction
	db = db.Select("ff.id, ff.name, ff.logo, uf.id AS userFictionId").
		Joins("left join fic_fiction as ff on uf.fiction_id = ff.id").
		Where("uf.user_id = ?", userId).Order("uf.create_time desc")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fictionList).Error
	return fictionList, total, err
}

func AddBookShelf(userFiction model.UserFiction) (err error) {
	err = global.GVA_DB.Create(&userFiction).Error
	return err
}
