package main

import (
	"encoding/json"
	"fmt"
)

func defineMap() {
	// Map은 키에 대응하는 값을 신속히 찾는 해시테이블을 구현한 자료형
	// 아래와 같이 map을 선언하면 nil값을 가지며, 이를 nil map이라고 한다.
	// nil Map에는 어떤 데이터를 쓸수없다. 따라서 make()함수를 사용하여 초기화 해야함

	var test map[string]int // 선언

	test = make(map[string]int) // 초기화
	test["ssoyyoung"] = 27

	fmt.Println(test)
	// 여기서 make 함수는 해시테이블 자료구조를 메모리에 생성하고, 그 메모리를 가리티는 map value를 리턴
	// 따라서 test 변수는 이 해시테이블을 가리키는 map을 가리키게 된다.

	// Map은 make()함수를 사용하여 초기화하는 방법과, 리터럴을 사용하는 방법 두가지가 존재

	test2 := map[string]int{
		"soyoung": 27,
	}
	fmt.Println(test2)
	// 위와 같이 사용하는 것이 리터럴 초기화 : map[key타입]value타입{키:값}
}

func jsonEncoding() {
	// JSON 인코딩 : Go의 데이터 타입을 json으로 변환(인코딩)하기 위해서는 encoding/json패키지의 Marshal을 사용
	// Go의 데이터를 json.Marshal()의 파라미터로 전달 시,Json으로 인코딩 된 바이트 배열을 리턴
	// []bytes를 다시 문자열로 변환해야한다면, string([]bytes)를 해준다.
	// 여기서 유의할 점은 Json의 Key는 문자열이여야 한다. map인 경우 Key가 string인 map만 지원

	// 리터럴 초기화를 사용하여 map 선언
	// Go 데이터 타입
	test := map[string]int{
		"Soyoung": 27,
	}

	// JSON 인코딩
	jsonBytes, _ := json.Marshal(test)

	// JSON 바이트
	fmt.Println(jsonBytes)

	// JSON 바이트를 문자열로 변경
	fmt.Println(string(jsonBytes))
}

func jsonDecoding() {
	// Json으로 인코딩된 데이터를 다시 디코딩하기 위해서는 encoding/json의 Unmarshal()함수 사용
	// Unmarshal()의 첫번째 파라미터는 json 데이터, 두번째 파라미터는 출력할 구조체 혹은 포인터를 지정
	// 리턴은 에러 객체, Unmarshal()함수 실행 시 두번째 파라미터에 원래 데이터가 복원된다.

	// =============================================================== V1
	// 테스트용 json 데이터 생성
	jsonBytes, _ := json.Marshal(map[string]int{"test": 123})

	// Json 디코딩 : Map 버전
	var returnVal map[string]int
	returnVal = make(map[string]int)

	err := json.Unmarshal(jsonBytes, &returnVal) // 두번째 인자를 넣을 때 &를 꼭 붙여줘야한다.
	if err != nil {
		panic(err)
	}

	fmt.Println(returnVal)

	// =============================================================== V2

	// 테스트용 struct 생성
	type user struct {
		Name string
		Age  int
	}

	// 테스트용 json 데이터 생성
	jsonBytes2, _ := json.Marshal(user{"testUser", 27})

	// Json 디코딩 : 구조체 버전
	var us user

	err = json.Unmarshal(jsonBytes2, &us)
	if err != nil {
		panic(err)
	}

	fmt.Println(us)
	fmt.Println(us.Name, us.Age)

}

func main() {
	// Map 테스트
	defineMap()

	// Json 인코딩
	jsonEncoding()

	// Json 디코딩
	jsonDecoding()

}
