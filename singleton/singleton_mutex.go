package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

// getInstance 这里使用了两个判断和锁来防止并发创建多个实例
// 第一个判断 singleInstance == nil 保证在最开始时实例为空
// 第二个判断 singleInstance == nil 防止多个线程绕过了第一个singleInstance == nil的判断，防止后进来的等锁unlock后再创建实例
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
