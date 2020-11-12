package main

import "fmt"

func main() {
	// 배열은 연속적인 메모리 공간에 동일한 타입의 데이터를 순서적으로 저장하는 자료구조
	// Go에서 배열의 첫번째 요소는 0번, 다음은 1,2,3 순서로 인덱스를 정한다. (Zero based array)

	// 배열의 선언 : var 변수명 [배열크기] 데이터 타입
	// Go에서 배열의 크기는 Type을 구성하는 3요소이다. 즉 [3]int와 [5]int는 다른 타입으로 인식된다.
	// 배열이 선언된 후에 각 배열의 요소를 인덱스를 사용하여 읽고 쓰기 가능

	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println(a)

	// 배열의 초기화
	// 배열을 정의할 떄, 초기값 설정할 수 있다.
	// var 변수명 [배열크기] 데이터 타입 {초기값}
	// 초기화 과정에서 ...을 사용하여 배열크기를 생략하면, 자동으로 초기화 요소 숫자만큼 배열크기가 정해진다.

	var b = [4]int{1, 2, 3, 4}

	fmt.Println(b)

	var c = [...]int{1, 2}
	fmt.Println(c)

	// 다배열 배열
	// Go 언어는 다차원 배열을 지원한다.
	// 다차원 배열은 배열크기 부분을 복수 개로 설정하여 선언한다
	// 3 x 4 x 5 차원 정수 배열

	var multiArray [2][3][4]int //정의
	multiArray[0][0][0] = 8

	fmt.Println(multiArray)

	// 다차원 배열의 초기화
	// 다차원 배열의 초기화는 단차원 배열의 초기화와 비슷
	// 다차원이므로 배열 초기값 안에 다시 배열값을 넣는 형태를 띤다

	var d = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(d)

	// Go 배열은 고정된 배열크기 안에 동일한 타입의 데이타를 연속적으로 저장하지만,
	// 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능을 가지고 있지 않다

}
