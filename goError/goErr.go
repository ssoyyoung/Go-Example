package main

import (
	"fmt"
	"os"
)

// Go 에러 ==============
// Go는 내장타입으로 error라는 interface를 가진다.
// Go 에러는 이 ERROR 인터페이스를 통해서 주고 받게 되는데, 이 interface는 다음과 같은 하나의 메서드를 갖는다.
// 개발자는 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.
type error interface {
	Error() string
}

func main() {
	// Go 에러처리 ==============
	// Go 함수가 결과와 에러를 함께 리턴한다면, 이 에러가 nil인지 체크해서 에러가 없는지를 체크
	// os.Open() 함수는 func Open(name string) (file *File, err error) 과 같은 함수 원형을 갖는 것
	// 첫번째는 파일 포인터를, 두번째는 error 인터페이스를 리턴
	// 두번째 error를 체크해서 nil이면 에러가 없는것, nil이 아니면 err.Error()로 해당 에러 확인
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("[ERROR]", err)         // 해당 에러에 대한 내용이 출력
		fmt.Println("[ERROR]", err.Error()) // 해당 에러에 대한 내용이 출력
	}
}
