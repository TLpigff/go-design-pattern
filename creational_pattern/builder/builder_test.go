/*
  @Author :     lyb
  @File :       builder_test
  @Description:
*/
package builder

import "testing"

func TestBuilder(t *testing.T) {
	dr := Director{Builder:&ConcreteBuilder{}}
	adCar := dr.Builder.SetType("SUV").AddBrand("奥迪").PaintColor("white").Build()
	adCar.Drive()


	bwCar := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bwCar.Drive()
}
