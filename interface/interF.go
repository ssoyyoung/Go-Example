package main

import (
	"fmt"
	"math"
	"reflect"
)

// GO 인터페이스
// struct : 필드의 집합체
// interface : 매서드의 집합체

// interface는 type이 구현해야하는 메서드 원형(prototype)들을 정의한다.
// 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다.

// Shape interface
type Shape interface {
	// interface는 type문을 사용하여 정의
	area() float64
	perimeter() float64
}

// 인터페이스를 구현하기 위해서 해당 타입이 그 인터페이스의 메서드들을 모두 구현하면 된다.
// 위의 Shape 인터페이스를 구현하기 위해서는 area(), perimeter() 2개의 메서드만 구현하면 된다

//Rect 정의
type Rect struct {
	width, height float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 정의
type Circle struct {
	radius float64
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		println(a)
	}
}

// 빈 인터페이스는 모든 타입을 받아드린다.
func printIt(i interface{}) {
	fmt.Println(i)
}

func main() {
	// 인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우를 들 수 있다
	// 함수 파라미터가 interface 인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 모두 입력 파라미터로 사용될 수 있다는 것을 의미

	r := Rect{10, 20}
	c := Circle{10}

	showArea(r, c)

	// 인터페이스 타입
	// 흔히 빈 인터페이스(empty interface)를 자주 접하게 되는데, 흔히 인터페이스 타입(interface type)으로도 불리운다
	// 빈 interface는 interface{} 와 같이 표현한다.

	// Empty interface는 메서드를 전혀 갖지 않는 빈 인터페이스로서
	// Go의 모든 Type은 적어도 0개의 메서드를 구현하므로, 흔히 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를 사용
	// 즉, 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너 = Dynamic Type

	var x interface{}

	x = 1
	x = "Soyoung"

	printIt(x)

	// Type Assertion
	var y interface{}
	fmt.Println("Before Type check..", y, reflect.TypeOf(y)) // nil, nil

	y = 1
	fmt.Println("Before Type check..", y, reflect.TypeOf(y))
	res := y.(int) // 값이 nil이거나, 타입이 맞지 않을 경우 panic 발생
	// x.(T)로 표현했을 때, 이는 x가 nil이 아니며,
	// x는 T 타입에 속한다는 점을 확인(assert)하는 것
	// 타입이 맞으면, 해당 타입의 값을 리턴
	fmt.Println("After Type check..", res, reflect.TypeOf(res))
}
