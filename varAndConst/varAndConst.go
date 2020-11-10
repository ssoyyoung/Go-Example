package main

import "fmt"

// 변수 테스트
func varTest() {
	// 변수 : 변수는 Go 키워드 var를 사용하여 선언.
	// var 변수명 변수타입
	var a int
	fmt.Println(a)

	// 변수 선언문에서 변수 초깃값 할당 가능
	var b float64 = 3.16
	fmt.Println(b)

	// 한번 선언된 변수는 그 뒤 문장에서 해당 타입의 값을 할당할 수 있다.
	a = 5
	b = 1.14

	fmt.Println(a, b)

	// 만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다.

	// 동일한 타입의 변수가 복수 개 있을 경우, 아래와 같이 사용 가능
	var d, e, f int

	// 변수를 선언만 한 경우, 초기값은 zero value를 기본적으로 할당한다.
	// 숫자형에는 0, bool타입은 false, string타입은 ""(빈문자열)이 할당된다.
	fmt.Println(d, e, f)

	// 복수 변수들이 선언된 상황에서 초기값을 지정 할 수 있다.
	// 초기값은 순서대로 변수에 할당
	var g, h, i int = 10, 11, 12

	fmt.Println(g, h, i)

	// Go에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 사용된다.
	var j = 1
	var k = "hello"

	fmt.Println(j, k)

	// 변수를 선언하는 다른 방식으로 짧은 선언문도 존재한다.
	// 즉 var i = 1 eotls i := 1 이라고 var 생략 가능
	// 이러한 표현은 func안에서만 사용할 수 있으며, 함수 밖에서는 var을 사용해야 한다.
	l := 316
	m := "hi"

	fmt.Println(l, m)
}

func constTest() {
	// 상수는 Go 키워드 const를 사용하여 선언한다.
	// const 상수명 상수타입 상수값
	const a int = 10
	const b string = "hi"

	fmt.Println(a, b)

	// Go에서는 할당되는 값을 보고 타입을 추론하는 기능을 자주 사용함
	const c = 316
	const d = "soyoung"
	fmt.Println(c, d)

	// Go에서는 여러 개의 상수를 묶어서 지정할 수 있다. 아래와 같이 괄호안에 상수를 나열
	const (
		e = "hello"
		f = "go"
		i = "!!"
	)
	fmt.Println(e, f, i)

	// 유용한 팁 : 상수값을 0부터 순차적으로 부여하기 위해 iota라는 identifier를 사용할 수 있다.
	// 이 경우, iota가 지정된 Apple에는 0이 할당, 나머지는 1씩 증가
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)
	fmt.Println(Apple, Grape, Orange)
}

func goKeyword() {
	// Go 프로그래밍 언어는 다음과 같은 25개의 예약 키워드들을 갖는다.
	// 이들 Go 키워드들은 변수명, 상수명, 함수명 등의 Identifier로 사용할 수 없다
	fmt.Println("break", "default", "func", "interface", "select")
	fmt.Println("case", "defer", "go", "map", "struct")
	fmt.Println("chan", "else", "goto", "package", "swtich")
	fmt.Println("const", "fallthrough", "if", "range", "type")
	fmt.Println("continue", "for", "import", "return", "var")
}

func main() {
	fmt.Println("varTest() ==================")
	varTest()

	fmt.Println("constTest() ==================")
	constTest()

	fmt.Println("goKeyword() ==================")
	goKeyword()
}
