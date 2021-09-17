package main

type sendCommand struct {
	app writeApp
}

func (c *sendCommand) execute() {
	c.app.send()
}

type getCommand struct {
	app writeApp
}

func (c *getCommand) execute() {
	c.app.get()
}
