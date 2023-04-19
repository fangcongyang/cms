package service

import (
	"cms/global"
	"cms/model"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

func ParseInfoList2Excel(infoList []model.SysMenu, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "路由Name", "路由Path", "父节点", "排序"})
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.Id,
			menu.Name,
			menu.Url,
			menu.ParentId,
			menu.Sort,
		})
	}
	excel.SaveAs(filePath)
	return nil
}

func ParseExcel2InfoList() ([]model.SysMenu, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]model.SysMenu, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id := global.GVA_WORKER.GetId()
		//hidden, _ := strconv.ParseBool(row[3])
		parentId, _ := strconv.ParseInt(row[3], 10, 64)
		sort, _ := strconv.Atoi(row[4])
		menu := model.SysMenu{
			GVA_MODEL: global.GVA_MODEL{
				Id: id,
			},
			Name:      row[1],
			Url:      row[2],
			//Hidden:    hidden,
			ParentId:  parentId,
			Sort:      sort,
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
