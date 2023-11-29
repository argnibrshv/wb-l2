Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет числа от 1 до 8 включительно в произвольном порядке и после этого будет
бесконечно выводить число 0.
Каналы a и b будут закрыты, после того как в функциях asChan все значения будут переданы в каналы.
Каналы a и b передаются функции merge, в функции горутины, созданной в функции merge, в бесконечном
цикле, с помощью оператора select, читаются данные из каналов a и b, и передаются в канал c.
При попытке получения данных из закрытого канала будет немедленно получено нулевое значение канала.
Поэтому после закрытия каналов a и b, в бесконечном цикле с оператором select в канал c будут
передаваться нулевые значения каналов a и b.
Чтобы решить эту проблему необходимо получать от канала вторую переменную булевого типа, которая
указывает открыт ли канал. Закрытому каналу необходимо присвоить значение nil, чтобы отключить ветвь 
с чтением из этого канала в операторе select, так как чтении из канала равного nil блокируется, 
и выполенение этой ветви никогда не произойдет.
```