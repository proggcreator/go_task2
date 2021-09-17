package main

//приложение симплексной передачей (в одну сторону)

func main() {
	writeApp := &writeApp{}

	sendCommand := &sendCommand{
		app: *writeApp,
	}

	getCommand := &getCommand{
		app: *writeApp,
	}

	sendButton := &button{
		command: sendCommand,
	}
	sendButton.press()

	getButton := &button{
		command: getCommand,
	}
	getButton.press()

}
