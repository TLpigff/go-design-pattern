/*
  @Time :       2020/4/16 2:06 下午
  @Author :     lyb
  @File :       simple_factory_test
  @Software:    GoLand
  @Description:
*/
package simpleFactory

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	factory := new(ShapeFactory)
	circleImpl := factory.getShape(CircleTypeNum)
	if circleImpl.Draw() != "I will draw a Circle" {
		t.Fatal("circleImpl test failed")
	}

	squareImpl := factory.getShape(SquareTypeNum)
	if squareImpl.Draw() != "I will draw a Square" {
		t.Fatal("squareImpl test failed")
	}

	rectangleImpl := factory.getShape(RectangleTypeNum)
	if rectangleImpl.Draw() != "I will draw a Rectangle" {
		t.Fatal("rectangleImpl test failed")
	}

	unsupportedImpl := factory.getShape(UnsupportedTypeNum)
	if unsupportedImpl != nil {
		t.Fatal("rectangleImpl test failed")
	}
}
