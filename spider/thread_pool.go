package spider

import "cms/common/datatype"

//这个是工作线程，处理具体的业务逻辑，将jobs中的任务取出，处理后将处理结果放置在results中。

func Worker(jobs <-chan string, proxys datatype.Queue, fun func(href string, proxys datatype.Queue)) {
	for j := range jobs {
		fun(j, proxys)
	}
}

func StartThreadPool(threadNum int, jobs <-chan string, proxys datatype.Queue, fun func(href string, proxys datatype.Queue)) {
	for w := 1; w <= threadNum; w++ {
		go Worker(jobs, proxys, fun)
	}
}
