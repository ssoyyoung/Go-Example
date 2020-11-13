package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

func structTest() {
	// struct 정의하기 위해서는 Custom Type을 정의하는데 type문을 사용
	p := Person{}
	p.name = "soyoung"
	p.age = 27

	fmt.Println(p)
}

func structTest2() {
	// 선언된 struct타입으로부터 객체를 생성하는 법
	// 1. 빈 객체를 먼저 할당하고, 나중에 필드에 채워넣는 방법
	p := Person{}
	p.name = "soyoung"
	p.age = 27

	// 2. struct 객체를 생성할 때, 초기값을 함께 할당하는 방법
	// 생략된 필드들은 Zero value 를 갖는다.
	//    > Zero value: 정수인 경우 0, float인 경우 0.0, string인 경우 "", 포인터인 경우 nil 등
	p1 := Person{"soso", 27}
	fmt.Println(p1)

	p2 := Person{
		name: "sososo",
		age:  27,
	}
	fmt.Println(p2)

	// 3. Go의 내장함수 new()사용
	// new()를 사용하면, 모든 필드를 Zero value로 초기화하고 person 객체의 포인터(*person)를 리턴.
	// 객체 포인터인 경우에도 필드 엑세스 시 . (dot)을 사용, 이 때 포인터는 자동으로 Dereference 된다
	p3 := new(Person)

	p3.name = "soyoung"
	fmt.Println(p3)
	fmt.Println(&p3)

	// Go에서 struct는 기본적으로 mutable 개체이다
	// 따라서 필드값이 변화할 경우 (별도로 새 개체를 만들지 않고) 해당 개체 메모리에서 직접 변경
	p3.name = "soyoung123"
	fmt.Println(p3)
	fmt.Println(&p3) // 필드의 값이 변화해도 메모리 주소는 동일

	// struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by Value에 따라 객체를 복사해서 전달하게 된다.
	// 그래서 만약 Pass by Reference로 struct를 전달하고자 한다면, struct의 포인터를 전달해야 한다.
	p4 := new(Person)
	p4.name = "soyoung"
	fmt.Println(&p4)
	passByRTest(p4)
	fmt.Println(p4)
}

func passByRTest(p *Person) {
	fmt.Println(&p)
	p.name = "change!"
}

type dict struct {
	data    map[int]string
	comment string
}

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func structTest3() {
	// 떄론 구조체 필드가 사용 전에 초기화되어야 하는 경우가 있다.
	// struct 의 필드가 map 타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct 사용자가 매번 map을 초기화 해야 된다는 것을 기억할 필요가 없다
	// 이러한 목적을 위해 생성자 함수를 사용할 수 있다
	// 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다
	dic := newDict()
	dic.data[1] = "test"
	dic.comment = "생성자 함수 테스트"

	fmt.Println(dic)

	// 생성자 함수를 사용하지 않으면, 아래와 같이 map 타입의 필드에 값을 넣기전에 초기화해야한다.
	dicT := new(dict)
	dicT.data = map[int]string{}
	dicT.data[1] = "test2"
	dicT.comment = "생성자 함수 테스트2"
	fmt.Println(dicT)

}

func main() {
	// Go에서 struct는 Custom Data Type을 표현하는데 사용
	// Go의 struct는 필드들의 집합체이며 필드들의 컨테이너
	// Go에서 struct는 필드 데이타만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다

	// Go 언어는 객체지향 프로그래밍을 고유의 방식으로 지원한다
	// 즉, Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다.
	// Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의

	structTest()
	structTest2()
	structTest3()

}
