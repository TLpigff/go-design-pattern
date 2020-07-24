/*
  @Author :     lyb
  @File :       prototype_test
  @Description:
*/
package prototype

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t	//取一个值
	fmt.Printf("t:%q %v,tc:%q %v\n",t,t,tc,tc)
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable  {
	tc := *t	//取一个值
	fmt.Printf("t:%q %v,tc:%q %v\n",t,t,tc,tc)
	return &tc
}

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{name:"type1"}
	manager.Set("t1",t1)
}


func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.Clone()

	fmt.Printf("t1:%p %v,t2:%p %v\n",t1,t1,t2,t2)

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
}