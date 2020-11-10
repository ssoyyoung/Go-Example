package main

import "fmt"

func main() {
	// Go 에서는 다른 언어에서와 비슷하게
	// 산술연산자, 관계연산자, 논리연산자, Bitwise 연산자, 할당연산자, 포인터연산자 등을 지원

	// 산술 연산자 : 사칙연산자(+, -, *, /, % (Modulus))와 증감연산자(++, --)를 사용
	a := 1
	b := 2
	c := a + b
	d := a * b
	fmt.Println(a, b, c, d)

	// 관계 연산자 : 서로의 크기를 비교하거나 동일함을 체크하는데 사용
	e := a == b
	f := c < d
	fmt.Println(a == b)
	fmt.Println(c < d)

	// 논리 연산자 : AND, OR, NOT을 표현하는데 사용
	fmt.Println(e && f)
	fmt.Println(e || f)

	// Bitwise 연산자 : Bitwise 연산자는 비트단위 연산을 위해 사용되는데, 바이너리 AND, OR, XOR와 바이너리 쉬프트 연산자가 있다.

	// 할당연산자 : 값을 할당하는 = 연산자 외에 사칙연산, 비트연산을 축약한 +=, &=, <<= 같은 연산자들도 있다.

	// 포인터연산자 : & 혹은 * 을 사용하여 해당 변수의 주소를 얻어내거나 이를 반대로 Dereference 할 때 사용한다.
	//             : Go 는 비록 포인터연산자를 제공하지만 포인터 산술 즉 포인터에 더하고 빼는 기능은 제공하지 않는다.

	var k int = 10
	var p = &k
	fmt.Println(*p)
}
