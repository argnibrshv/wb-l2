Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет:
"error"
Функция test() возвращает значение типа интерфейса, возвращаемое значение равно nil, но переменная типа интерфейса будет равна nil, только когда nil будет равны и значение переменной и динамический тип интерфейса. 

```
