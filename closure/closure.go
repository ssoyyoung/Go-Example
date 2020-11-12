package main

import "fmt"

// int 를 리턴하는 익명함수를 리턴한다.
// Go 언어에서 함수는 일급함수로써 다른 함수로부터 리턴되는 리턴값을 사용할 수 있다.
// 아래에서 익명함수가 그 함수 밖에 있는 변수 i를 참조한다.
// 익명함수 자체가 로컬 변수로 i 를 갖는 것이 아니기 때문에 외부 변수 i 가 상태를 계속 유지하
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Go 언어에서 함수는 Closure로서 사용될 수도 있다
	// Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 말한다.
	// 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.

	next := nextValue() // Closure 함수를 next라는 변수에 할당

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	anotherNext := nextValue() // 새로운 Closure 함수값을 생성
	println(anotherNext())
	println(anotherNext())

}
