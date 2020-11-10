package main

import (
	"io"
	"log"
	"os"
)

var myLogger *log.Logger

func loggerT() {
	// 새로운 로거를 만들기 위해 log.New()함수 사용
	// log.New()는 3개의 파라미터가 존재, 첫번째는 io.Writer 인터페이스를 지원하는 타입으로 콘솔출력(os.Stdout), 표준에러(os.Stderr), 파일포인터 등
	// 두번째 파라미터는 로그 출력의 가장 처음에 적는 Prefix로서 프로그램명, 카테고리 등을 기재
	// 세번째 파라미터는 로그 플래그로 표준플래그(log.LstdFlags), 날짜플래그(log.Ldate), 시간플래그(log.Ltime), 파일위치플래그(log.Lshortfile, log.Llongfile) 등을 | (OR 연산자)로 묶어 지정

	myLogger = log.New(os.Stdout, "INFO: ", log.Lshortfile)

	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Printf("run func")
}

func multiLoggerT() {
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	run2()
	log.Println("End of multiLogger Program")
}

func run2() {
	log.Printf("multiLogger run func")
}

func main() {
	// Go에서 로그를 사용하기 위해 Go의 표준 패키지인 log를 이용할 수 있다.
	// log 패키지의 Print*() 같은 함수를 직접 사용하는 경우, 디폴트 출력인 Stdout에 날짜/시간과 함께 로그가 출력된다.

	log.Println("Logging")

	log.SetFlags(0) // 날짜, 시간을 없애고자 하는 경우
	log.Println("Logging")

	// 로그 패키지는 여러가지 로거를 지원하기 위해 Logger 라는 인터페이스/타입을 제공
	// 위에서는 별도의 Logger를 생성하지 않았기 때문에 표준 Logger를 사용한것

	loggerT()
	multiLoggerT()
}
