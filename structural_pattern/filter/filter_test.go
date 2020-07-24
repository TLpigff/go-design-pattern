/*
  @Author :     lyb
  @File :       filter_test
  @Description:
*/
package filter

import (
	"fmt"
	"log"
	"testing"
)

func TestFilter(t *testing.T) {
	// 字符串分割，再转成int，再求和
	split := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()

	// 这个地方一定要顺序一致
	p := NewPipeline("p1",split,converter,sum)
	ret,err := p.Process("4,5,6")
	if err != nil {
		log.Fatal(err)
	}

	if ret != 15 {
		log.Fatalf("the expected is 15,but the actual is %d",ret)
	}
	fmt.Println(ret)
}
