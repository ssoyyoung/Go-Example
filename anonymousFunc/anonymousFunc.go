package main

import "fmt"

func anonymousTest() {
	// 함수명을 갖지 않는 함수를 익명함수(Anonymous Function)이라 부른다
	// 1. 익명함수는 그 함수 전체를 변수에 할당
	//   > 변수에 할당된 익명함수는 func 뒤에 함수명이 생략
	//   > 익명함수가 변수에 할당된 이후에는 변수명이 함수명과 같이 취급
	//   > "변수명(파라미터들)" 형식으로 함수를 호출
	// 2. 다른 함수의 파라미터에 직접 정의되어 사용

	sum := func(nums ...int) int { // 익명함수 정의
		sumValue := 0
		for _, num := range nums {
			sumValue += num
		}
		return sumValue
	}

	result := sum(1, 2, 3, 4, 5) // 익명함수 호출
	fmt.Println("result value is", result)
}

func calc(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func funcTest() {
	// 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급
	// 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다
	// 즉, 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용될 수 있다

	// 함수를 다른 함수의 파라미터로 전달하기 위해서는
	// 1. 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법
	// 2. 직접 다른 함수 호출 파라미터에 함수를 적는 방법

	// 변수에 함수 할당
	add := func(i, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(a int, b int) int { return a * b }, 10, 20)
	fmt.Println(r2)

}

// [TODO] : type문을 사용한 함수 원형 정의

func main() {
	fmt.Println("anonymousTest ===============")
	anonymousTest()

	fmt.Println("funcTest ===============")
	funcTest()

}
