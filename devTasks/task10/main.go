package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	host    string
	port    int
	timeout time.Duration
)

func readRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	defer cancel()
	scanner := bufio.NewScanner(conn)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break read")
			break OUTER
		default:
			if scanner.Err() != nil {
				fmt.Println(scanner.Err())
			}
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN from conn")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("from server: %s", text)
		}
	}

	log.Printf("finished read")
}

func writeRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	defer cancel()
	scanner := bufio.NewScanner(os.Stdin)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break write")
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			_, err := conn.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				log.Printf("error on write to server: %v\n", err)
				break OUTER
			}
		}
	}

	log.Printf("finished write")
}

func parseArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("should be host and port")
	}

	host := args[0]
	if host == "" {
		return errors.New("host should be not empty")
	}

	port, _ := strconv.Atoi(args[1])
	if port == 0 {
		return errors.New("port can't be zero")
	}

	return nil
}

func main() {

	flagtime := flag.String("timeout", "10s", "print only separator")

	flag.Parse()
	err := parseArgs(flag.Args())
	if err != nil {
		fmt.Println(err)
	}
	//задаем хост
	host = flag.Arg(0)
	//задаем номер порта
	port, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("error name port transform")
	}

	timeout, err = time.ParseDuration(*flagtime)
	if err != nil {
		fmt.Println("error time  transform")
	}

	//инициализируем параметры
	addr := fmt.Sprintf("%s:%v", host, port)
	fmt.Printf("trying %s, timeout: %s...\n", addr, timeout)

	dialer := &net.Dialer{Timeout: timeout}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//создание соединения
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	fmt.Printf("connected to %s\n", addr)

	//чтение
	go readRoutine(ctx, conn, cancel)
	//запись
	go writeRoutine(ctx, conn, cancel)

	//ждем сигнала завершения
	select {
	case <-ctx.Done():
		fmt.Println("shutdown signal received")
	}
	//закрыть соединение
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection closed")
}
