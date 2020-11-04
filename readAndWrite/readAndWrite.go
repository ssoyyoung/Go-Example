package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

// readAndWrite func > version1
func readAndWrite() {
	fileRead, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer fileRead.Close()

	fmt.Println(fileRead)
	fmt.Println(reflect.TypeOf(fileRead)) // *os.File

	fileOpen, err := os.Create("test2.txt")
	if err != nil {
		panic(err)
	}
	defer fileOpen.Close()

	// slice를 생성하기 위해 make 내장함수 사용, 두번째 인자로 길이설정
	buff := make([]byte, 30)

	for {
		// file 읽기
		cnt, err := fileRead.Read(buff)
		if err == io.EOF {
			break // file의 끝까지 다읽었을경우 반복문 break
		}

		// byte to string
		str := string(buff[:cnt])
		// string to byte
		newBuff := []byte(str)

		fmt.Println(str, newBuff)

		// file 쓰기
		_, err = fileOpen.Write(buff[:cnt])
		if err != nil {
			panic(err) //  현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴
		}
	}
}

// readAndWrite func > version1
func readAndWrite2() {
	// file 읽기
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	// file 쓰기
	err = ioutil.WriteFile("test3.txt", bytes, 0644) //0644 : 소유자는 읽고 쓸수있다.
	if err != nil {
		panic(err)
	}

}

func main() {
	readAndWrite()
	readAndWrite2()
}

// GOPATH
// export GOPATH=$HOME/go
// export GOROOT=/usr/local/go
// export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
// export PATH=$PATH:/usr/local/go/bin
