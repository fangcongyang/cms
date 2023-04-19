package spider

import (
	"bufio"
	"cms/common/datatype"
	"cms/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Ltengxsw(i string) {
	proxys := Ip3306(10)
	c := colly.NewCollector(colly.DetectCharset())

	q, _ := queue.New(
		4, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	// Find and visit all links
	c.OnHTML("#views_con_1", func(e *colly.HTMLElement) {
		e.DOM.Find("td").Each(func(i int, selection *goquery.Selection) {
			td := selection.Text()
			if strings.HasPrefix(td, "《") {
				href, _ := selection.Find("a").Attr("href")
				err := q.AddURL(href)
				if err != nil {
					return
				}
			}
		})
		ltengxswFiction(q, proxys)
	})

	c.OnRequest(func(r *colly.Request) {
		proxyURL := fmt.Sprintf("%v", proxys.Dequeue())
		r.ProxyURL = proxyURL
		proxys.Enqueue(proxyURL)
	})

	err := c.Visit(fmt.Sprintf("http://www.ltengxsw.com/top/allvisit/%s.htm", i))
	if err != nil {
		return
	}
}

func ltengxswFiction(q *queue.Queue, proxys datatype.Queue) {
	c := colly.NewCollector(colly.DetectCharset())

	// Find and visit all links
	c.OnHTML("li[class='button2 white'] a", func(e *colly.HTMLElement) {
		href, _ := e.DOM.Attr("href")
		ltengxswFictionChapter(href, proxys)
	})

	c.OnRequest(func(r *colly.Request) {
		proxyURL := fmt.Sprintf("%v", proxys.Dequeue())
		r.ProxyURL = proxyURL
		proxys.Enqueue(proxyURL)
	})

	err := q.Run(c)
	if err != nil {
		return
	}

	c.Wait()
	// 小说下载完成释放锁
	utils.UnLock("fictionSpider")
}

func ltengxswFictionChapter(fictionChapterUrl string, proxys datatype.Queue) {
	c := colly.NewCollector(colly.DetectCharset())
	// Find and visit all links
	c.OnHTML("div[class=novel]", func(e *colly.HTMLElement) {
		novel_head := e.DOM.Find("div[class=novel_head]")
		fiction_name := novel_head.Find("h1").Text()
		sort := novel_head.Find("p")
		author := sort.Find("a").Text()
		sort_name := sort.Text()
		filePath := "D:/fiction/" + fiction_name + ".txt"
		is_exist, err := utils.PathExists(filePath)
		if is_exist {
			fmt.Println("文件已下载")
			return
		}
		f, err := os.Create(filePath)
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		//Flush将缓存的文件真正写入到文件中
		write.WriteString(strings.Split(sort_name, "：")[2] + "\n")
		write.WriteString("fangcy\n")
		write.WriteString(author + "\n")
		write.WriteString("fangcy\n")
		write.WriteString("\n")
		write.WriteString("fangcy\n")
		compile, _ := regexp.Compile("[第弟](.+)[章张]")
		e.DOM.Find("div[class=novel_list]").Find("a").Each(func(i int, selection *goquery.Selection) {
			chapter_name := selection.Text()
			if compile.MatchString(chapter_name) {
				chapter_name = strings.ReplaceAll(strings.Trim(chapter_name, " "), compile.FindString(chapter_name), "")
			}
			href, _ := selection.Attr("href")
			write.WriteString("第" + strconv.Itoa(i+1) + "章 " + chapter_name + "\n")
			write.WriteString("fang|c|y\n")
			ltengxswDownFictionChapter(href, write, proxys)
		})
		f.Close()
		// 关闭文件
		file.Close()
		// 下载完成移动文件
		os.Rename(filePath, strings.ReplaceAll(filePath, "fiction", "fictions"))
	})

	c.OnRequest(func(r *colly.Request) {
		proxyURL := fmt.Sprintf("%v", proxys.Dequeue())
		r.ProxyURL = proxyURL
		proxys.Enqueue(proxyURL)
	})

	c.Visit(fictionChapterUrl)
}

func ltengxswDownFictionChapter(href string, writer *bufio.Writer, proxys datatype.Queue) {
	c := colly.NewCollector(colly.DetectCharset())

	// Find and visit all links
	c.OnHTML("div[class=novel_content]", func(e *colly.HTMLElement) {
		writer.WriteString(e.Text)
		writer.WriteString("\nfangcy\n")
		writer.Flush()
		time.Sleep(2 * time.Second)
	})

	c.OnRequest(func(r *colly.Request) {
		proxyURL := fmt.Sprintf("%v", proxys.Dequeue())
		r.ProxyURL = proxyURL
		proxys.Enqueue(proxyURL)
	})

	err := c.Visit(href)
	if err != nil {
		return
	}
}
