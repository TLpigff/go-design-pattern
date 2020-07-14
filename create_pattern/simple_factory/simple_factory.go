/*
  @Time :       2020/4/16 1:55 下午
  @Author :     lyb
  @File :       simple_factory
  @Software:    GoLand
  @Description:
*/
package simpleFactory

import "fmt"

type TypeNum int32

const (
	UnsupportedTypeNum TypeNum = 0
	CircleTypeNum TypeNum = 1
	SquareTypeNum TypeNum = 2
	RectangleTypeNum TypeNum = 3
)

// shape接口
type Shape interface {
	Draw() string
}

type Circle struct {}

func (c *Circle) Draw()  string{
	return "I will draw a Circle"
}

type Square struct {}

func (s *Square) Draw()  string{
	return "I will draw a Square"
}

type Rectangle struct {}

func (r *Rectangle) Draw()  string{
	return "I will draw a Rectangle"
}

type ShapeFactory struct {}

func (c *ShapeFactory) getShape(typ TypeNum) Shape {
	switch typ {
	case CircleTypeNum:
		return new(Circle)
	case SquareTypeNum:
		return new(Square)
	case RectangleTypeNum:
		return new(Rectangle)
	default:
		fmt.Println("unsupported type")
		return nil
	}
}


