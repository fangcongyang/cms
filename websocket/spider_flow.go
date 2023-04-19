package websocket

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func ParseSpiderFlowXml(data string) *MxGraphModel {
	v := new(MxGraphModel)

	err := xml.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return v
}

func BindFlow(mxCells []MxCell) *SpiderNode {
	var nodeMap = make(map[string]*SpiderNode)
	var root = new(SpiderNode)
	var firstNode = new(SpiderNode)
	var edgeMap = make(map[string]MxCell)
	for _, cell := range mxCells {
		var jsonProperty map[string]interface{}
		err := json.Unmarshal([]byte(cell.JsonProperty), &jsonProperty)
		if err != nil {
			fmt.Println("JsonToMapDemo err: ", cell.JsonProperty)
		}
		var node = new(SpiderNode)
		node.JsonProperty = jsonProperty
		node.NodeId = cell.Id
		node.NodeName = cell.Value
		nodeMap[cell.Id] = node
		if cell.Edge == "1" { //判断是否是连线
			edgeMap[cell.Id] = cell
		} else if jsonProperty != nil && jsonProperty["shape"] != "" {
			if "start" == jsonProperty["shape"] {
				root = node
			}
		}
		if "0" == cell.Id {
			firstNode = node
		}
	}
	for _, v := range edgeMap {
		var sourceNode = nodeMap[v.Source]
		var targetNode = nodeMap[v.Target]
		sourceNode.NextNodes.PushBack(targetNode)
	}
	firstNode.NextNodes.PushBack(root)
	return firstNode
}
