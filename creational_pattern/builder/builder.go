/*
  @Time :       2020/7/14 4:20 下午
  @Author :     lyb
  @File :       builder
  @Software:    GoLand
  @Description:
*/
package builder

import "fmt"

// 产品角色
type Car struct {
	Brand string
	Type string
	Color string
}

func (c *Car) Drive() error {
	fmt.Printf("A %s %s %s car is running on the road!\n", c.Color, c.Type, c.Brand)
	return nil
}

// 建造者角色
type Builder interface {
	PaintColor(color string) Builder
	AddBrand(brand string) Builder
	SetType(t string) Builder
	Build() Car
}

// 具体的建造者
type ConcreteBuilder struct {
	ACar Car
}

func (c *ConcreteBuilder) PaintColor(color string) Builder {
	c.ACar.Color = color
	return c
}

func (c *ConcreteBuilder) AddBrand(brand string) Builder {
	c.ACar.Brand = brand
	return c
}

func (c *ConcreteBuilder) SetType(t string) Builder {
	c.ACar.Type = t
	return c
}

func (c *ConcreteBuilder) Build() Car {
	return c.ACar
}


// 定义导演者角色
type Director struct {
	Builder Builder
}



