/*
  @Time :       2020/7/14 4:05 下午
  @Author :     lyb
  @File :       singleton
  @Software:    GoLand
  @Description:
*/
package singleton

import "sync"

type Singleton struct {
	Payload string
}

var singleton *Singleton
var once sync.Once

// sync.Once是go内部自己实现的专门用于执行单次的，内部有加锁机制
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			Payload:"I am once",
		}
	})
	return singleton
}
