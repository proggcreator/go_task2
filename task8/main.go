package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// ввод
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// обработка ввода
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// если команда без оргумента
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// удаление символа новой строки
	input = strings.TrimSuffix(input, "\n")

	// разделить на слова
	args := strings.Split(input, " ")

	// проверка команд
	switch args[0] {
	case "cd":
		// пустой путь cd
		if len(args) < 2 {
			return ErrNoPath
		}
		//реализация cd
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// команда отдельно
	cmd := exec.Command(args[0], args[1:]...)

	// поток вывода
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//выполнить команду
	return cmd.Run()
}
