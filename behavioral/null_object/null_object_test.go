/*
  @Author :     lyb
  @File :       null_object_test
  @Description:
*/
package null_object

import (
	"fmt"
	"testing"
)

func TestNullObjectPattern(t *testing.T) {
	customerFactory := new(CustomerFactory)
	customerFactory.Names = []string{"Rob", "Joe", "Julie"}

	customer1 := customerFactory.GetCustomer("Rob")
	customer2 := customerFactory.GetCustomer("Bob")
	customer3 := customerFactory.GetCustomer("Julie")
	customer4 := customerFactory.GetCustomer("Laura")

	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
}

