/*
  @Author :     lyb
  @File :       flyweight_test
  @Description:
*/
package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	//viewer := NewImageViewer("image1.png")
	//viewer.Display()

	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}