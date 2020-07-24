/*
  @Author :     lyb
  @File :       state_test
  @Description:
*/
package state

import "testing"

func TestState(t *testing.T) {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()	// 输出
		ctx.Next()	// 改变状态
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
}
