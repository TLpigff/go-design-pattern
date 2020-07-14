/*
  @Time :       2020/7/14 11:20 上午
  @Author :     lyb
  @File :       abstract_factory_test.go
  @Software:    GoLand
  @Description:
*/
package abstractFactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	// new 一个抽象工厂生产者
	factoryProducer := new(FactoryProducer)

	// 获取shape抽象工厂
	shapeFactory := factoryProducer.getFactory(ShapeFactoryTypeNum)
	// 获取shape抽象工厂的Circle对象并调用它的Draw方法
	fmt.Println(shapeFactory.getShape(CircleShapeTypeNum).Draw())
	// 获取shape抽象工厂的Square对象并调用它的Draw方法
	fmt.Println(shapeFactory.getShape(SquareShapeTypeNum).Draw())
	// 获取shape抽象工厂的Rectangle对象并调用它的Draw方法
	fmt.Println(shapeFactory.getShape(RectangleShapeTypeNum).Draw())

	if shapeFactory.getColor(BlueColorTypeNum) != nil {
		t.Fatal("shapeFactory dose not support getColor")
	}


	// 获取color抽象工厂
	colorFactory := factoryProducer.getFactory(ColorFactoryTypeNum)
	// 获取shape抽象工厂的Circle对象并调用它的Draw方法
	fmt.Println(colorFactory.getColor(BlueColorTypeNum).Fill())
	// 获取shape抽象工厂的Square对象并调用它的Draw方法
	fmt.Println(colorFactory.getColor(RedColorTypeNum).Fill())
	// 获取shape抽象工厂的Rectangle对象并调用它的Draw方法
	fmt.Println(colorFactory.getColor(GreenColorTypeNum).Fill())

	if colorFactory.getShape(CircleShapeTypeNum) != nil {
		t.Fatal("colorFactory dose not support getShape")
	}
}