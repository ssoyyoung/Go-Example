package main

import "fmt"

func ifTest() {
	// 해당 조건이 맞으면 {} 블럭안의 내용을 실행
	// 반드시 조건 블럭 시작 { 괄호를 if문과 같은 라인에 두어야 한다.
	// if 문의 조건식은 반드시 Boolean식으로 표현되어야 한다.

	k := 1

	if k == 1 {
		fmt.Println("one")
	}

	// if 문은 else if, else문과 함께 사용할 수 있다.

	h := 2

	if h == 1 {
		fmt.Println("one")
	} else if h == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("not one or two")
	}

	// if 문에서 조건식을 사용하기 이전에 간단한 문장을 함께 실행 할 수 있다.
	// 여기서 사용된 변수는 if문 안에서만 사용 가능

	i := 3

	if val := i * 2; val > 4 {
		fmt.Println(val)
	}
}

func switchTest() {
	// 여러값을 비교해야 하는 경우, 다수의 조건식을 체크해야하는 경우 switch문을 사용

	var name string
	var category = 1

	switch category {
	case 1:
		name = "soyoung park"
	case 2:
		name = "soyoung"
	default:
		name = "other"
	}

	fmt.Println(name)

	// Swtich문에서 사용되는 용법들
	// 1. switch뒤에 expression이 없을 수 있음 : 이 경우 true로 생각하고 첫번째 case로 넘어감
	// 2. case 문에 expression을 쓸 수 있음 : Go는 case문에 복잡한 expression을 가질 수 있다.
	// 3. No default fall through : Go는 다음 case로 가지 않는다.
	// 4. Type switch :  Go는 그 변수의 Type에 따라 case로 분기할 수 있다
	// 5. switch 뒤에 조건 변수 혹은 Expression을 적지 않는다.
	//  >> 이 경우 각 case 조건문들을 순서대로 검사하여 조건에 맞는 경우 해당 case블럭을 실행하고 switch 문을 빠져나온다.

	score := 68
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}

	// Go는 마지막에 Go 컴파일러가 자동으로 각 case 블럭마다 마지막에 break문을 추가한다.

}

func main() {
	ifTest()
	switchTest()
}
