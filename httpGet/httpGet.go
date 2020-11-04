package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func simpleResq() {
	resp, err := http.Get("https://sol.piclick.kr/similarSearch/2473/148_41670")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(data))
}

func withHeaderResq() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "https://sol.piclick.kr/similarSearch/2473/148_41670", nil)
	if err != nil {
		panic(err)
	}

	// 해더 추가
	req.Header.Add("User-Agent", "soyoung")

	// Client 객체에서 Request 실행
	client := &http.Client{}
	resq, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resq.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resq.Body)
	str := string(bytes)

	fmt.Println(str)
}

func main() {
	// http 패키지는 웹 관련 클라이언트 및 서버 기능을 제공
	// http.Get() 메서드는 쉽게 웹 페이지나 웹 데이터들 가져오는데 사용

	// 간편하게 호출 가능하나, 세밀한 컨트롤 불가
	fmt.Println("simpleResq----------")
	simpleResq()

	// 섬세한 컨트롤이 가능, 직접 Request객체 생성
	fmt.Println("withHeaderResq----------")
	withHeaderResq()
}
