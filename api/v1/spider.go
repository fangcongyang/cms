package v1

import (
	bytes2 "bytes"
	"cms/global"
	"cms/model"
	"cms/model/response"
	"cms/service"
	"cms/spider"
	"cms/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// @Tags SysUser
// @Summary 开启小说爬虫抓取
// @Security ApiKeyAuth
// @accept formdata
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spider/getUserList [get]

func StartFictionSpider(c *gin.Context) {
	isLock := utils.Lock("fictionSpider", "fictionSpider")
	if isLock {
		spider.Ltengxsw(c.Query("pageNo"))
		response.Ok(c)
	} else {
		response.FailWithMessage("正在爬取小说数据!", c)
	}
}

// @Tags SysUser
// @Summary 读取目录小说文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]

func SaveFictionToDataBase(c *gin.Context) {
	isLock := utils.Lock("SaveFictionToDataBase", "SaveFictionToDataBase")
	if isLock {
		filePath := c.Query("filePath")

		files, e := ioutil.ReadDir(filePath)
		if e != nil {
			utils.UnLock("SaveFictionToDataBase")
			response.FailWithMessage("当前目录不存在!", c)
		}
		for _, v := range files {
			fileName := filepath.Join(filePath, v.Name())
			bytes, _ := ioutil.ReadFile(fileName)
			fictionStr := string(bytes)
			ficArr := strings.Split(fictionStr, "fangcy")
			tx := global.GVA_DB.Begin()
			sortId, e := service.GetFictionSortIdByName(strings.ReplaceAll(strings.ReplaceAll(ficArr[0], " ", ""), "\n", ""))
			if e != nil {
				tx.Rollback()
				utils.UnLock("SaveFictionToDataBase")
				response.FailWithMessage("获取小说分类id失败", c)
				return
			}
			introduce := strings.ReplaceAll(strings.ReplaceAll(ficArr[2], " ", ""), "\n", "")
			chapterIntroduce := strings.ReplaceAll(strings.ReplaceAll(strings.Split(ficArr[3], "fang|c|y")[1], " ", ""), "\n", "")
			if introduce == "" {
				if utf8.RuneCountInString(chapterIntroduce) > 100 {
					rs := []rune(chapterIntroduce)
					introduce = string(rs[0:100])
				} else {
					introduce = chapterIntroduce
				}
			}
			fiction := model.Fiction{
				Name: strings.ReplaceAll(strings.ReplaceAll(strings.Split(v.Name(),".")[0], " ", ""), "\n", ""),
				Status: 1,
				Score: 5.0,
				Author: strings.ReplaceAll(strings.ReplaceAll(ficArr[1], " ", ""), "\n", ""),
				Introduce: "  " + introduce,
			}

			fictionId, e := service.SaveFiction(fiction)
			if e != nil {
				tx.Rollback()
				utils.UnLock("SaveFictionToDataBase")
				response.FailWithMessage("保存小说失败", c)
				return
			}
			e = service.SaveFictionSortFiction(fictionId, sortId)
			if e != nil {
				tx.Rollback()
				utils.UnLock("SaveFictionToDataBase")
				response.FailWithMessage("保存小说分类小说关系失败", c)
				return
			}
			length := len(ficArr)
			for i := 3;  i < length; i++ {
				chapters := strings.Split(ficArr[i],"fang|c|y")
				if len(chapters) != 2 {
					continue
				}
				chapter := strings.Split(chapters[0], " ")
				var chapterName bytes2.Buffer
				if len(chapter) >= 2 {
					for j := 0; j < len(chapter); j++ {
						if j == 0 {
							continue
						}
						chapterName.WriteString(chapter[j])
					}
				}
				compile, _ := regexp.Compile("\\d+")
				sort, _ := strconv.Atoi(compile.FindString(chapter[0]))
				chapterContext := chapters[1]
				chapterNameStr := strings.ReplaceAll(strings.ReplaceAll(chapterName.String(), "\n", "")," ", "")
				chapterDto := model.FictionChapter{
					Name: chapterNameStr,
					FictionId: fictionId,
					Sort: sort,
					WordCount: len(strings.ReplaceAll(chapterContext, " ", "")),
					Context: chapterContext,
				}
				e = service.SaveFictionChapter(chapterDto)
				if e != nil {
					tx.Rollback()
					utils.UnLock("SaveFictionToDataBase")
					response.FailWithMessage("保存小说章节失败", c)
					return
				}
			}
			err := os.Remove(fileName)
			if err != nil {
				global.GVA_LOG.Info(v.Name() + "文件删除成功")
			}
		}
		utils.UnLock("SaveFictionToDataBase")
		response.Ok(c)
	} else {
		response.FailWithMessage("正在同步小说数据!", c)
	}
}