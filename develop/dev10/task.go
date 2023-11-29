package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

var timeout string

func init() {
	flag.StringVarP(&timeout, "timeout", "t", "10s", "timeout for connecting to the server")
}

func scanStdIn(stdinChan chan<- string, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = fmt.Sprintf("%s\n", text)
		stdinChan <- text
	}
	cancel()
}

func scanConnIn(conn net.Conn, connChan chan<- string, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		connChan <- text
	}
	cancel()
}

func writeData(ctx context.Context, conn net.Conn, connChan <-chan string, stdinChan <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case text := <-connChan:
			fmt.Println(text)
		case text := <-stdinChan:
			conn.Write([]byte(text))
		}
	}
}

func main() {
	flag.Parse()
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("wrong timeout value: %s", err)
	}
	addressArg := flag.Args()
	if len(addressArg) != 2 {
		log.Fatal("введите хост и порт для подключения в формате HOST PORT")
	}

	ctx, cancel := context.WithCancel(context.Background())
	connChan := make(chan string)
	stdinChan := make(chan string)

	var wg sync.WaitGroup

	address := fmt.Sprintf("%s:%s", addressArg[0], addressArg[1])

	conn, err := net.DialTimeout("tcp", address, timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go scanStdIn(stdinChan, cancel)
	go scanConnIn(conn, connChan, cancel)

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeData(ctx, conn, connChan, stdinChan)
	}()
	wg.Wait()
}
