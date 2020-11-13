package main

import "fmt"

func mapTest1() {
	// Map은 키에 대응하는 값을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조
	// Go 언어는 Map 타입을 내장 :  "map[Key타입]Value타입" 과 같이 선언할 수 있다

	// 아래는 정수를 키로 하고 문자열을 값으로 하는 맵 변수인 idMap를 선언한 것
	// map은 reference 타입이므로 nil값을 가진다. 이를 nil map이라 한다.
	// Nil map에는 값을 넣을 수 없다.
	var idMap map[int]string // Nil map

	// ============ 초기화 방법 1
	// Nil map에는 어떤 데이터를 쓸 수 없다. 따라서 map을 초기화하기 위해 make()함수 사용
	// make()함수의 첫번째 파라미터로 "맵키워드[키타입]값타입" 으로 지정
	// 음 map이 make() 함수에 의해 초기화 되었을 때는, 아무 데이타가 없는 상태
	// 초기화 후에는 값을 쓸 수 있다.
	idMap = make(map[int]string) // 이떄는 아무 데이터가 존재하지 않는다.
	fmt.Println(idMap)
	idMap[1] = "test" // 데이타를 추가하기 위해서는 "map변수[키] = 값" 과 같이 해당 키에 그 값을 할당
	// 새로운 키에 값을 할당하면 새 해시 키-값 쌍이 추가
	// 만약 키 901의 값이 이미 존재했다면, 추가대신 값만 갱신한다.
	fmt.Println(idMap)

	// make() 함수 내부 로직
	// 해시테이블 자료구조를 메모리에 생성, 그 메모리를 가리키는 map value 리턴 (map value는 내부적으로 runtime.hmap 구조체를 가리키는 포인터이다)
	// 따라서 idMap 변수는 이 해시테이블을 가리키는 map을 가리키게 된다.

	// ============ 초기화 방법 2
	// 리터럴을 사용해 초기화 할 수 있다.
	// 리터럴 초기화는 "map[Key타입]Value타입 { key:value }" 와 같이 Map 타입 뒤 { } 괄호 안에 "키: 값" 들을 열거
	idMap2 := map[int]string{
		1: "test1",
		2: "test2",
		3: "test3",
	}

	fmt.Println(idMap2)

	// map 키 지우기
	delete(idMap2, 2)
	fmt.Println(idMap2)

	//  만약 map안에 찾는 키가 존재하지 않는다면 reference 타입인 경우 nil을 value 타입인 경우 zero를 리턴한다.

}

func mapTest2() {
	// map 키 체크
	// map을 사용하는 경우 종종 map안에 키가 존재하는지 체크
	// 2개의 리턴, 첫번째는 키에 상응하는 값, 두번째는 존재 여부 = map[키값]

	idMap3 := map[int]string{
		1: "test1",
		2: "test2",
		3: "test3",
	}

	val, exist := idMap3[2]
	fmt.Println(val, exist) // 값이 존재하면, 키에 상응하는 값과, true 가 리턴

	val2, exist2 := idMap3[4]
	fmt.Println(val2, exist2) // 값이 존재하지않으면, 빈값과, false 가 리턴

}

func mapTest3() {
	// for 을 활용한 map 열거
	idMap4 := map[int]string{
		1: "test1",
		2: "test2",
		3: "test3",
	}

	for key, val := range idMap4 {
		fmt.Println(key, val)
	}

}

func main() {
	mapTest1()
	mapTest2()
	mapTest3()
}
