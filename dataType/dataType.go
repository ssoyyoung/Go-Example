package main

import (
	"fmt"
	"reflect"
)

// 문자열 테스트
func stringTest() {
	// 문자열 리터럴은 `` 혹은 ""를 사용하여 표현 가능

	// Back Quote
	// ``로 둘러싸인 문자열은 별도 해석없이 Raw string 그대로를 갖는다.
	// 예를들어 문자열 안의 \n를 인식하지 않는다.
	// 복수 라인의 문자열을 표현할 때 자주 사용

	a := `abcd \n string test. 
	this is new line`
	fmt.Println(a)

	// 이중인용부호(" ")로 둘러 싸인 문자열은 Interpreted String Literal이라 부르는데,
	// 복수 라인에 걸쳐 쓸 수 없으며,
	// 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다.
	// 예를 들어, 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석된다.
	// 이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.

	b := "abcd \n string test. this is new line"
	fmt.Println(b)

	c := "abcd \n string test." +
		"this is new line"
	fmt.Println(c)
}

// 데이터 타입변환 테스트
func convertType() {
	// 하나의 데이타 타입에서 다른 데이타 타입으로 변환하기 위해서는 T(v) 와 같이 표현
	// 여기서 T는 변환하고자 하는 타입을 표시하고, v는 변환될 값(value)을 지정한 것이다.

	// 정수 100을 float로 변경하기 위하여 float32(100) 처럼 표현하고, 문자열을 바이트배열로 변경하기 위하여 []byte("ABC") 처럼 표현할 수 있다.

	var a int = 100
	fmt.Println(a, ">>", reflect.TypeOf(a))

	b := float64(a)
	fmt.Println(b, ">>", reflect.TypeOf(b))

	c := int8(a)
	fmt.Println(c, ">>", reflect.TypeOf(c))

	var d string = "abc"
	fmt.Println(d, ">>", reflect.TypeOf(d))

	e := []byte(d)
	fmt.Println(e, ">>", reflect.TypeOf(e))

	f := string(e)
	fmt.Println(f, ">>", reflect.TypeOf(f))
}

func main() {
	// Go의 데이터 타입
	// bool, string, int, float,
	// byte : unit8타입과 동일하며 바이트 코드에 사용
	// rune : int32와 동일하며 유니코드 코드포인터에 사용
	stringTest()

	convertType()

}
