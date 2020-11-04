package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: cfg,
	}
}

// simplePost : Plain text post
func simplePost() {
	reqBody := bytes.NewBufferString("Post plain text")
	fmt.Println(reqBody)

	resq, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resq.Body.Close()

	//Response 체크
	resqBody, err := ioutil.ReadAll(resq.Body)
	if err == nil {
		str := string(resqBody)
		fmt.Println(str)
	}
}

// formValuePost : formvalue post
func formValuePost() {
	// Form 데이터를 보내는데 유용한 매서드
	//resq, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	resq, err := http.PostForm("https://seoulbitz.sparker.kr:1323/getNear/shop", url.Values{"subway": {"강남"}})
	if err != nil {
		panic(err)
	}

	defer resq.Body.Close()

	// Response 체크
	resqBody, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resqBody))

}

// Person struct
type Person struct {
	Name string
	Age  int
}

// jsonDataPost : json data post
func jsonDataPost() {
	person := Person{"Bibiane", 27}
	pbytes, err := json.Marshal(person) // 논리적 구조체를 로우 바이트로 변경 = 마샬링/인코딩
	// json.Marshal : Go에서 모든 타입을 받을 수 있는 interface{} 타입을 인자로 받고 바이트 슬라이스를 반환
	fmt.Println(pbytes)

	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(pbytes) // inputParam:[]bytes, outputParam: *Buffer
	fmt.Println(buff)

	// json을 보내는 경우 두번째 파라미터에 application/json을 적고
	// 세번째 파라미터에 JSON으로 인코딩된 데이터 전달
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	// Response 체크
	resqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resqBody))
}

//xmlDataPost : xml data post
func xmlDataPost() {
	person := Person{"Bibiane", 27}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/xml", buff)

	// Response 체크
	resqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resqBody))
}

// useHardControlPost func
func useHardControlPost() {
	person := Person{"soyoung", 27}
	pbytes, _ := xml.Marshal(person) // struct to byte arrays
	buff := bytes.NewBuffer(pbytes)

	// Request 객체 생성
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	// Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/xml")

	// Client 객체에서 Requests 실행
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크
	resqBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(resqBody)
		fmt.Println(str)
	}
}

func main() {
	// http 패키지는 웹 관련 클라이언트 및 서버 기능을 제공
	// http.Post() 메서드는 웹서버로 간단히 데이터를 POST 하는데 사용
	simplePost()
	formValuePost()
	jsonDataPost()
	xmlDataPost()
	useHardControlPost()
}
