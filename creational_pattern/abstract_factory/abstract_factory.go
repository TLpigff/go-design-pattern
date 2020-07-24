/*
  @Time :       2020/4/27 2:48 下午
  @Author :     lyb
  @File :       abstract_factory
  @Software:    GoLand
  @Description:
*/
package abstractFactory

import "fmt"

// 抽象工厂

type ShapeTypeNum int32
type ColorTypeNum int32
type FactoryTypeNum int32

const (
	UnsupportedShapeTypeNum ShapeTypeNum = 0
	CircleShapeTypeNum      ShapeTypeNum = 1
	SquareShapeTypeNum      ShapeTypeNum = 2
	RectangleShapeTypeNum   ShapeTypeNum = 3

	UnsupportedColorTypeNum ColorTypeNum = 0
	BlueColorTypeNum        ColorTypeNum = 1
	RedColorTypeNum         ColorTypeNum = 2
	GreenColorTypeNum       ColorTypeNum = 3

	UnsupportedFactoryTypeNum = 0
	ShapeFactoryTypeNum       = 1
	ColorFactoryTypeNum       = 2
)

type FactoryProducer struct{}

func (f *FactoryProducer) getFactory(factoryTyp FactoryTypeNum) AbstractFactory {
	switch factoryTyp {
	case ShapeFactoryTypeNum:
		return new(ShapeFactory)
	case ColorFactoryTypeNum:
		return new(ColorFactory)
	default:
		return nil
	}
}

type AbstractFactory interface {
	getShape(shapeTyp ShapeTypeNum) Shape
	getColor(colorTyp ColorTypeNum) Color
}

type ShapeFactory struct{}

func (a *ShapeFactory) getShape(shapeTyp ShapeTypeNum) Shape {
	switch shapeTyp {
	case CircleShapeTypeNum:
		return new(Circle)
	case SquareShapeTypeNum:
		return new(Square)
	case RectangleShapeTypeNum:
		return new(Rectangle)
	default:
		fmt.Println("unsupported shape type")
		return nil
	}
}

func (a *ShapeFactory) getColor(colorTyp ColorTypeNum) Color {
	return nil
}

type ColorFactory struct{}

func (a *ColorFactory) getShape(shapeTyp ShapeTypeNum) Shape {
	return nil
}

func (a *ColorFactory) getColor(colorTyp ColorTypeNum) Color {
	switch colorTyp {
	case BlueColorTypeNum:
		return new(Blue)
	case RedColorTypeNum:
		return new(Red)
	case GreenColorTypeNum:
		return new(Green)
	default:
		fmt.Println("unsupported color type")
		return nil
	}
}

// shape接口
type Shape interface {
	Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "I will draw a Circle"
}

type Square struct{}

func (s *Square) Draw() string {
	return "I will draw a Square"
}

type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return "I will draw a Rectangle"
}

// color接口
type Color interface {
	Fill() string
}

type Blue struct{}

func (c *Blue) Fill() string {
	return "I will fill in Blue"
}

type Red struct{}

func (s *Red) Fill() string {
	return "I will fill in Red"
}

type Green struct{}

func (r *Green) Fill() string {
	return "I will fill in Green"
}
