package pattern

// Фасад (Facade) — это структурный паттерн, который предоставляет простой интерфейс, упрощающий
// использование сложной системы объектов, библиотек или фреймворков.

// Когда имеется сложная система, и необходимо упростить с ней работу, фасад позволит определить
// одну точку взаимодействия между клиентом и системой.
// Когда надо уменьшить количество зависимостей между клиентом и сложной системой,
// фасадные объекты позволяют отделить, изолировать компоненты системы от клиента и развивать
// и работать с ними независимо.
// Создание фасадов для компонентов каждой отдельной подсистемы позволит упростить взаимодействие
// между ними и повысить их независимость друг от друга.

// Минус паттерна фасад,структура фасада может разрастись до объекта-бога, большим количеством
// методов.

//  Если приложение взаимодействует с различными сторонними сервисами и библиотеками,
// фасад может упростить интеграцию, предоставляя чистый и удобный интерфейс
// для взаимодействия с ними.

// SubsystemA, SubsystemB, SubsystemC являются компонентами
// сложной подсистемы, с которыми должен взаимодействовать клиент.
type SubsystemA struct {
}

func (s SubsystemA) Amethod() {
	println("some logic in SubsystemA")
}

type SubsystemB struct {
}

func (s SubsystemB) Bmethod() {
	println("some logic in SubsystemB")
}

type SubsystemC struct {
}

func (s SubsystemC) CAmethod() {
	println("some logic in SubsystemC")
}

// Facade - непосредственно фасад, который предоставляет интерфейс
// клиенту для работы с компонентами.
type Facade struct {
	sA *SubsystemA
	sB *SubsystemB
	sC *SubsystemC
}

func (f Facade) FacadeMethodOne() {
	f.sA.Amethod()
	f.sB.Bmethod()
	f.sC.CAmethod()
}

func (f Facade) FacadeMethodTwo() {
	f.sA.Amethod()
	f.sC.CAmethod()
}
