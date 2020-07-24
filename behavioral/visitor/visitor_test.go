/*
  @Author :     lyb
  @File :       visitor_test
  @Description:
*/
package visitor

import "testing"

func TestVisitor(t *testing.T) {
	//c := &CustomerCol{}
	//c.Add(NewEnterpriseCustomer("A company"))
	//c.Add(NewEnterpriseCustomer("B company"))
	//c.Add(NewIndividualCustomer("bob"))
	//c.Accept(&ServiceRequestVisitor{})

	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&AnalysisVisitor{})
}
