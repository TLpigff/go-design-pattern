/*
  @Author :     lyb
  @File :       iterator_test
  @Description:
*/
package iterator

import "testing"

func TestIterator(t *testing.T) {
	var addregate Aggregate
	addregate = NewNumbers(1,10)

	IteratorPrint(addregate.Iterator())
}
