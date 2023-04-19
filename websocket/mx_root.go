package websocket

import (
	"container/list"
	"encoding/xml"
	"github.com/gocolly/colly/v2"
)

type MxGraphModel struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	MxRoot  MxRoot   `xml:"root"`
}

type MxRoot struct {
	MxCells []MxCell `xml:"mxCell"`
}

type MxCell struct {
	Id           string `xml:"id,attr"`
	Value        string `xml:"value,attr"`
	Style        string `xml:"style,attr"`
	Parent       string `xml:"parent,attr"`
	Vertex       string `xml:"vertex,attr"`
	Source       string `xml:"source,attr"`
	Target       string `xml:"target,attr"`
	Edge         string `xml:"edge,attr"`
	JsonProperty string `xml:"JsonProperty"`
}

type SpiderNode struct {
	NodeId       string
	NodeName     string
	JsonProperty map[string]interface{}
	NextNodes    list.List
}

func (spiderNode SpiderNode) Request(taskId string) {
	spiderProcessDataInstance := GetSpiderProcessDataInstance()
	c := colly.NewCollector(colly.DetectCharset())

	for i, parameterName := range spiderNode.JsonProperty["parameter-name"].([]string) {
		c.OnHTML(spiderNode.JsonProperty["parameter-node"].([]string)[i], func(e *colly.HTMLElement) {
			if spiderNode.JsonProperty["parameter-is-list"].([]bool)[i] {
				parameterValue, ok := spiderProcessDataInstance.spiderData[taskId].(map[string]interface{})[parameterName]
				if ok {
					println(parameterValue)
				}
			}
			e.DOM.Find(spiderNode.JsonProperty["parameter-sub-node"].(string))
		})
	}

	err := c.Visit(spiderNode.JsonProperty["url"].(string))
	if err != nil {
		return
	}
}

type SocketMessage struct {
	NodeId    string  `json:"nodeId"`
	EventType string  `json:"eventType"`
	Message   Message `json:"message"`
}

type Message struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Variables string `json:"variables"`
}
