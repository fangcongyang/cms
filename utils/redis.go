package utils

import (
	"cms/global"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

func Lock(key string, value string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	lockSuccess, err := global.GVA_REDIS.SetNX(key, value, 2*time.Hour).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return lockSuccess
}

func UnLock(key string) {
	global.GVA_REDIS.Del(key).Result()
}
