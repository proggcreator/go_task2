package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//nc tcp 192.168.0.113 80
func ClientNetcat(arg []string) error {
	if len(arg) < 1 {
		return errors.New("Error: few netcat arguments")
	}
	// arg задает udp или tcp протокол
	c, err := net.Dial(arg[0], arg[1]+":"+arg[2])
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer c.Close()

	//отправка данных
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		//отправка данных
		fmt.Fprintf(c, text+"\n")
		//ответ от сервера
		//message, _ := bufio.NewReader(c).ReadString('\n')
		//fmt.Print("->: " + message)

		//закрыть клиент
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return nil
		}
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// ввод
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// обработка ввода, разбиваем по |
		smas := strings.Split(input, "|")
		var wg sync.WaitGroup
		wg.Add(len(smas))
		for i := 0; i <= len(smas); i++ {
			go func() {
				if err = execInput(smas[i]); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}

}
func execInput(input string) error {
	// удаление символа новой строки
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	// проверка команд
	switch args[0] {
	case "cd":
		// пустой путь cd
		if len(args) < 2 {
			return errors.New("path required")
		}
		//реализация cd
		return os.Chdir(args[1])

		//утилита
	case "nc":
		err := ClientNetcat(args[1:])
		if err != nil {
			return err
		}
		return nil

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	// поток вывода
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//выполнить команду cd/pwd/echo/kill/ps
	return cmd.Run()
}
