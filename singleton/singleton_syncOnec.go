package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleSync struct {
}

var singleInstanceSync *singleSync

// getInstanceSync 这种方式在项目中获取配置文件信息已经在使用了
// 需要注意的是，每一个sync.Once对应一个单例，不能重复使用
func getInstanceSync() *singleSync {
	if singleInstanceSync == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstanceSync = &singleSync{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstanceSync
}
