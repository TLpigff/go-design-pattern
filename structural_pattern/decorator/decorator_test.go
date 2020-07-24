/*
  @Author :     lyb
  @File :       decorator_test.go
  @Description:
*/
package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c,10)
	c = WrapMulDecorator(c,8)
	res := c.Calc()

	fmt.Println(res)
}