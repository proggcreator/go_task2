package main

type command interface {
	execute()
}
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
