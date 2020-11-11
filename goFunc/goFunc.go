package main

import "fmt"

func funcTest1(msg string) {
	// 함수는 여러 문장을 묶어서 실행하는 코드 블럭의 단위
	// Go에서 함수는 func 키워드를 사용하여 정의
	//  함수 파라미터는 0개 이상 사용할 수 있다.
	// 각 파라미터는 파라미터명 뒤에 int, string 등의 파라미터 타입을 적어서 정의
	// 함수의 리턴 타입은 파라미터 괄호 ( ) 뒤에 적는다.
	fmt.Println("[func 1]", msg)

	msg = "func test change!"

	fmt.Println("[func 1]", msg)
}

func funcTest2(msg *string) {
	// Go에서 파라미터를 전달하는 방식은 아래와 같이 크게 2개로 나뉜다.
	// 1. Pass By Value
	//    > 위의 예제에서 msg의 값이 복사되어 함수에게 전달
	//    > 만약 파라미터 msg의 값이 함수에서 변경되더라도 호출함수인 main에서 msg의 값은 그대로이다.
	// 2. Pass By Reference
	//    > 아래 예제와 같이 msg 변수 앞에 & 부호를 붙이면 msg 변수의 주소를 표시한다.
	//	  > 흔히 포인터라 불리는 이 용법은, 함수에게 msg 변수의 값을 복사하지 않고, 변수의 주소를 전달한다.
	//	  > 피호출 함수는 *으로 파라미터가 포인터임을 표시한다.
	//	  > 이떄 msg는 문자열이 아닌, 문자열을 갖는 메모리 영역의 주소를 갖는다.
	//	  > msg 주소에 데이터를 쓰기 위해서는 *msg = "" 과 같이 앞에 *를 붙이는데 이를 흔히 Dereferencing 이라 한다.

	fmt.Println("[func 2]", msg) // 주소값이 프린트 된다.

	*msg = "func test change!"

	fmt.Println("[func 2]", msg) // 변수의 값은 바뀌여도 메모리 주소값은 동일하게 프린트 된다.
}

func funcTest3(msgList ...string) {
	// 함수에 고정된 수의 파라미터가 아닌, 다양한 숫자의 파라미터를 전달하고자 할때
	// 가변 파라미터를 나타내는 ... (3개의 마침표)을 사용한다.
	// 즉 문자열 가변 파라미터를 나타내기 위해서 ...string 과 같이 표현
	// 가변 파라미터를 갖는 함수를 호출할 때는 0개, 1개, 2개, 혹은 n개의 동일타입 파라미터를 전달할 수 있다
	// 가변 파라미터를 받아들이는 함수를 Variadic Function (가변인자함수)라고 부른다

	for _, s := range msgList {
		fmt.Print(s)
	}
	fmt.Println()
}

func funcTest4(nums ...int) int {
	// Go 프로그래밍 언어에서 함수는 리턴값이 없을 수도, 하나 일 수도, 또는 복수 개일 수도 있다.
	// Go 언어는 또한 Named Return Parameter 라는 기능을 제공
	// 이는 리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있는 기능

	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum is", sum)

	return sum

}

func funcTest5(nums ...int) (int, int) {
	// Go에서 복수 개의 값을 리턴하기 위해서는 해당 리턴 타입들을 괄호 ( ) 안에 적어 준다

	var count int
	var sum int
	for _, num := range nums {
		sum += num
		count++
	}
	fmt.Println("sum is", sum)
	fmt.Println("intput parameter count is", count)

	return sum, count
}

func funcTest6(nums ...int) (sum int, count int) {
	// Go에서 Named Return Parameter들에 리턴값들을 할당하여 리턴
	// 리턴되는 값들이 여러 개일 때, 코드 가독성을 높이는 장점
	// 리턴 파라미터와 그 타입을 함께 정의한것이다. (따라서 변수를 함수내에서 선언하지 않아도 된다.)
	// 그리고 함수내에서 이 리턴 파라미터에 결과값을 직접 할당하고 있다.
	// 실제 return 문에는 아무 값들을 리턴하지 않지만, 그래도 리턴되는 값이 있을 경우에는 빈 return 문을 반드시 써 주어야 한다

	for _, num := range nums {
		sum += num
		count++
	}
	fmt.Println("[Named Return] sum is", sum)
	fmt.Println("[Named Return] intput parameter count is", count)

	return
}

func main() {
	fmt.Println("func test 1 ==========")
	msg := "func test!"
	funcTest1(msg)
	fmt.Println("[main]", msg) // funcTest에서 msg의 값을 변경하더라도, main에서는 원래 값이 유지된다.

	fmt.Println("func test 2 ==========")
	funcTest2(&msg)
	fmt.Println("[main]", msg)

	fmt.Println("func test 3 ==========")
	funcTest3("a", "b", "c", "d")
	funcTest3("a", "b")

	fmt.Println("func test 4 ==========")
	funcTest4(1, 2, 3, 4, 5)

	fmt.Println("func test 5 ==========")
	funcTest5(1, 2, 3, 4, 5)

	fmt.Println("func test 6 ==========")
	funcTest6(1, 2, 3, 4, 5)
}
