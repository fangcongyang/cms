package main

import (
	"cms/common/config"
	"cms/core"
	"cms/global"
	"cms/initialize"
	"cms/websocket"
	"net/http"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	worker, _ := config.NewWorker(1)
	global.GVA_WORKER = worker        // 初始化雪花算法
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB)
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go startWebsocket()
	core.RunWindowsServer()
}

func startWebsocket() {
	http.HandleFunc("/ws", websocket.WsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
