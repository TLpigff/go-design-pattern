/*
  @Author :     lyb
  @File :       command
  @Description:
*/
package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb:mb}
}


type RebootCommand struct {
	mb *MotherBoard
}

func (r *RebootCommand) Execute() {
	r.mb.Reboot()
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb:mb}
}


type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1,button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}