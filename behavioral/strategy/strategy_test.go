/*
  @Author :     lyb
  @File :       strategy_test
  @Description:
*/
package strategy

import "testing"

func TestStrategy(t *testing.T) {
	payment := NewPayment("Ada","",123,&Cash{})
	payment.Pay()

	payment.strategy = &Bank{}
	payment.Pay()
}
