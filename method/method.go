package main

import (
	"fmt"
	"strconv"
)

// ========== Method ==========
// Go 언어는 객체지향 프로그래밍(OOP)을 고유의 방식으로 지원
// 타 언어의 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리
// Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의

// Go 매서드는 특별한 형태의 func 함수이다.
// 매서드는 함수의 정의에서 func 키워드와 함수명 사이에 *를 사용하여 어떤 struct을 위한 매서드인지 표시한다.
// 흔히 receiver라고 불리는 이 부분은, 매서드가 속한 struct 타입과 struct 변수명을 지정한다.
// struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다

// Person struct
type Person struct {
	name string
	age  int
}

// Value receiver는 struct의 데이터를 복사(copy)하여 전달, 포인터 receiver는 struct의 포인터만을 전달

// Pointer receiver : 메서드 내의 필드값 변경이 그대로 호출자에서 반영
func (p *Person) introduce() string {
	fmt.Println(&p) // 포인터
	p.age++
	return "[name] " + p.name + " [age] " + strconv.Itoa(p.age)
}

// Value receiver :  매서드내에서 struct의 필드값이 변경되더라도, 호출자의 데이터는 유지
func (p Person) introduce2() string {
	fmt.Println(&p) // 값
	p.age++
	return "[name] " + p.name + " [age] " + strconv.Itoa(p.age)
}

func main() {
	p := Person{"soyoung", 27}
	res := p.introduce() // method 호출
	fmt.Println(res)

	res2 := p.introduce2() // method 호출
	fmt.Println(res2)
}
