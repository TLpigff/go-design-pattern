/*
  @Author :     lyb
  @File :       memento_test
  @Description:
*/
package memento

import "testing"

func TestMemento(t *testing.T) {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()

	progress := game.Save()

	game.Play(-2,-3)
	game.Status()

	game.Load(progress)
	game.Status()
}