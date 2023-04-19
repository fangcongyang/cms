package websocket

import (
	"cms/global"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		spiderProcessDataInstance := GetSpiderProcessDataInstance()
		if data != nil {
			var jsonData map[string]string
			err := json.Unmarshal(data, &jsonData)
			if err != nil {
				fmt.Println("JsonToMapDemo err: ", string(data))
			}
			taskId := jsonData["taskId"]
			currSpiderData, ok := spiderProcessDataInstance.spiderData[taskId]
			if !ok {
				currSpiderDataMap := make(map[string]interface{})
				currSpiderDataMap["message"] = jsonData["message"]
				eventType := jsonData["eventType"]
				var isDebug = "debug" == eventType
				currSpiderDataMap["isRun"] = isDebug || "test" == eventType
				currSpiderDataMap["isDebug"] = isDebug
				if isDebug {
					currSpiderDataMap["runNextNode"] = make(chan bool)
					currSpiderDataMap["runNextNode"].(chan bool) <- true
				}
				spiderProcessDataInstance.spiderData[taskId] = currSpiderDataMap
				go handlerSpider(taskId, currSpiderDataMap, nil, conn)
			} else {
				currSpiderDataMap := currSpiderData.(map[string]interface{})
				currSpiderDataMap["runNextNode"].(chan bool) <- true
			}
		}
	}

ERR:
	conn.Close()

}

type SpiderProcessData struct {
	spiderData map[string]interface{}
}

var spiderProcessDataInstance *SpiderProcessData
var spiderProcessDataOnce sync.Once

func GetSpiderProcessDataInstance() *SpiderProcessData {
	spiderProcessDataOnce.Do(func() {
		spiderProcessDataInstance = &SpiderProcessData{}
		spiderProcessDataInstance.spiderData = make(map[string]interface{})
	})
	return spiderProcessDataInstance
}

func handlerSpider(taskId string, currSpiderDataMap map[string]interface{}, runNextNode chan string, conn *Connection) {
	var xml = currSpiderDataMap["message"].(string)
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(":", zap.Any("err", err))
		}
	}()
	if currSpiderDataMap["isRun"].(bool) {
		if xml != "" {
			var v = ParseSpiderFlowXml(xml)
			var node = BindFlow(v.MxRoot.MxCells)
			for e := node.NextNodes.Front(); e != nil; e = e.Prev() {
				if currSpiderDataMap["isDebug"].(bool) {
					<-runNextNode
				}
				node = e.Value.(*SpiderNode)
				var shape = node.JsonProperty["shape"]
				switch shape {
				case "request":
					InvokeObjectMethod(new(SpiderNode), "Request", taskId)
					fmt.Println("Golang")
				case "Rust":
					fmt.Println("Rust")
				default:
					fmt.Println("PHP是世界上最好的语言")
				}
				var socketMessage = new(SocketMessage)
				socketMessage.NodeId = node.NodeId
				socketMessage.EventType = "log"
				socketMessage.Message.Level = ""
				socketMessage.Message.Message = "执行节点" + node.NodeId + "{}节点名称:" + node.NodeName
				socketMessage.Message.Variables = ""
				jsons, _ := json.Marshal(socketMessage)
				if err := conn.WriteMessage(jsons); err != nil {
					goto ERR
				}
			}
			close(runNextNode)
		}
		delete(GetSpiderProcessDataInstance().spiderData, taskId)
	}
ERR:
	conn.Close()
}
