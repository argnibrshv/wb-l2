package main

import (
	"fmt"
	"os"

	// Импортируем модуль ntp.
	"github.com/beevik/ntp"
)

func main() {
	// С помощью метода Time модуля ntp, обращаемся к NTP серверу и в случае
	// успеха получаем текущее точное время.
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.2org")
	// Если при обращении к NTP серверу произойдет ошибка, значение err будет
	// не равно nil, в этом случаем текст ошибки выводится в поток ошибок
	// и программа завершит свою работу со статусом 1 с помощью метода
	// Exit из пакета os.
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
	// Печатаем полученное от NTP сервера точное время.
	fmt.Println(t)
}
