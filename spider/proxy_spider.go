package spider

import (
	"cms/common/datatype"
	"cms/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
)

func Ip3306(ipNum uint) (proxys datatype.Queue) {
	q := datatype.Queue{}
	q.Init()
	pageNo := 1
	ParseIp("http://www.ip3366.net/?stype=1&page="+strconv.Itoa(pageNo), q)
	for {
		if q.Size() >= ipNum { //循环条件
			break // 跳出for循环,结束for循环
		}
		ParseIp("http://www.ip3366.net/?stype=1&page="+strconv.Itoa(pageNo), q)
		fmt.Println(q.Size())
		pageNo++
	}
	return q
}

func ParseIp(url string, q datatype.Queue) {
	c := colly.NewCollector(colly.DetectCharset())

	// Find and visit all links
	c.OnHTML("#list tbody tr", func(e *colly.HTMLElement) {
		ip := e.ChildTexts("td")
		if utils.Ping(ip[0], 32, 4) {
			q.Enqueue(ip[3] + ":" + ip[0] + ":" + ip[1])
		}
	})

	c.Visit(url)
}
