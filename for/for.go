package main

import "fmt"

func forTest1() {
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func forTest2() {
	// Go에서 for 루프는 초기값과 증감식을 생략하고 조건식만 사용할 수 있다.
	// 이는 마치 for 루프가 다른 언어의 while루프와 같이 쓰이도록한다.

	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)
}

func forTest3() {
	// for 루프로 무한루프를 만들려면, 초기값;조건식;증감 모두를 생략하면된다.
	for {
		fmt.Println("Warning : Infinite loop")
	}
}

func forTest4() {
	// for range문은 컬렉션으로부터 한 요소씩 가져와 차례로 for블럭의 문장을 실행한다.
	// 다른언어의 foreach와 비슷한 용법

	// for 인덱스, 요소값 := range 컬렉션 : 왼쪽과 같이 for루프를 구성한다.
	// range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치 인덱스와 값을 for 키워드 다음의 2개의 변수에 각각 할당

	names := []string{"go", "python", "java"}

	for idx, name := range names {
		fmt.Println(idx, name)
	}
}

func forTest5() {
	// 경우에 따라 for 루프내에서 즉시 빠져나올 필요가 있는데, 이때 break문을 사용한다.
	// 만약 for 루프의 중간에서 나머지 문장을 실행하지 않고 for 루프의 시작부분으로 바로 가려면 continue 사용
	// 기타 임의의 문장으로 이동하기 위해서는 goto문 사용
	// goto 문은 for루프와 관련없이 사용 가능

	// break문은 for 루프 이외에 switch문이나 select문에서도 사용가능
	// 하지만 continue는 for 루프와 연관되어 사용된다.

	var a = 1

	for a < 15 {
		if a == 5 {
			a += a
			fmt.Println("continue")
			continue
		}
		a++
		if a > 10 {
			fmt.Println("break")
			break
		}
	}

	if a == 11 {
		fmt.Println("a == 11")
		goto END
	}
	fmt.Println("here!", a)

END:
	fmt.Println("go to END")
}

func forTest6() {
	// break 문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블"과 같이 사용하여 지정된 레이블로 이동할 수도 있다
	// break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데
	// 이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고
	// break문의 직속 for 루프 전체의 다음 문장을 실행하게 한다.

	i := 0

L1:
	for {
		if i == 0 {
			break L1
		}
	}

	println("OK")
}

func main() {
	// Go 프로그래밍 언어에서 반복문은 for 루프를 이용한다. Go는 반복문에 for 하나만 존재함
	// for 초기값;조건식;증감 {} : 초기값, 조건식, 증감식은 경우에 따라 생략가능

	forTest1()

	forTest2()

	forTest4()

	forTest5()

	forTest6()

}
