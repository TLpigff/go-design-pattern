/*
  @Author :     lyb
  @File :       chain_of_responsibility_test
  @Description:
*/
package chain_of_responsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c1
	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 100000)
}
