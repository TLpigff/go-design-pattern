/*
  @Time :       2020/7/14 4:09 下午
  @Author :     lyb
  @File :       singleton_test.go
  @Software:    GoLand
  @Description:
*/
package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance := GetInstance()
	fmt.Println(instance.Payload)

	instance.Payload = "I am second"

	instance1 := GetInstance()
	fmt.Println(instance1.Payload)
}