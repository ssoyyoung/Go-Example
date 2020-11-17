package main

import (
	"fmt"
	"os"
)

func main() {
	// 지연실행 defer
	// Go의 defer 키워드는 특정 문장 혹은 함수를 나중에(함수가 리턴하기 직전)에 실행하게 된다.
	// 아래 예제는 에러가 발생하더라도 항상 파일을 Close 할수 있게 한다.

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err.Error()) // 현재함수의 defer을 모두 실행한 뒤 즉시 리턴
	}

	defer f.Close() // 함수가 끝나기 직전에 실행된다.

	bytes := make([]byte, 1024)
	// fmt.Println(bytes) 길이가 1024인 0으로 초기화된 배열이 생성
	f.Read(bytes)
	println(len(bytes))
	println(string(bytes))

	defer fmt.Println("panic 종료")

	// panic 함수
	// 내장함수인 panic()함수는 현재 함수를 즉시 멈추고, 현재 함수의 defer들을 모두 실행한 후 즉시 리턴
	// 이러한 panic 모드는 실행 방식은 다시 상위함수에도 똑같이 적용, 계속 콜스택을 타고 올라가며 적용된다.
	// 마지막에는 프로그램이 에러를 내고 종료하게된다.
	f1, err1 := os.Open("test2.txt") // 잘못된 파일명
	if err1 != nil {
		fmt.Println(err1.Error())
		// 현재함수의 defer을 모두 실행한 뒤 즉시 리턴
	}

	defer fmt.Println("panic 종료 2")
	defer f1.Close()

	// recover 함수
	// Go의 내장 함수인 recover() 함수는 panic함수에 의한 패닉상태에서 다시 정상상태로 되돌리는 함수다.

	openFile("test3.txt")

}

func openFile(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f2, err2 := os.Open(fileName)

	if err2 != nil {
		panic(err2)
	}

	defer f2.Close()
}
