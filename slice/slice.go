package main

import "fmt"

func sliceTest1() {
	// 슬라이스는 배열과 달리 고정된 크기를 미리 지정하지 않을 수 있다.
	// 차후 그 크기를 동적으로 변경할 수 있고, 부분 배열 발췌가 가능하다.
	// 배열과 달리 초기 선언시 크기를 지정하지 않는다.
	var a []int // slice type
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)
}

func sliceTest2() {
	// Go에서 슬라이스를 생성하는 또 다른 방법 : Go의 내장함수인 make() 사용
	// make()함수로 슬라이스 생성시, 개발자가 슬라이스의 길이와 용량을 임의로 지정할 수 있다.
	// make() 함수의 첫번째 파라미터에 생성할 슬라이스 타입 지정, 두번쨰는 Length(슬라이스의 길이),  세번째는 Capacity (내부 배열의 최대 길이)를 지정
	// 위와 같이 지정시, 모든 요소가 Zero인 슬라이스를 만들게 된다.
	// 이떄 세번째 파라미터를 생략하면 Capacity는 Length와 같은 값을 갖는다
	// 그리고 슬라이스의 길이 및 용량은 내장함수 len(), cap()을 써서 확인할 수 있다
	b := make([]int, 5, 10)
	fmt.Println(len(b), cap(b))
}

func sliceTest3() {
	// 슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0 인 슬라이스를 만드는데, 이를 Nil Slice 라 하고, nil 과 비교하면 참을 리턴한다.
	var c []int

	if c == nil {
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(c), cap(c))

	c = append(c, 1)
	fmt.Println(len(c), cap(c))
}

func sliceTest4() {
	// 부분 슬라이스
	// 슬라이스에서 일부를 발췌하여 부분 슬라이스를 만들 수 있다
	// 부분 슬라이스는 "슬라이스[처음인덱스:마지막인덱스]" 형식으로 만든다.
	// 마지막인덱스는 원하는 인덱스+1 을 사용한다. 즉, 처음인덱스는 Inclusive 이며, 마지막인덱스는 Exclusive이다
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(d[2:5])
}

func sliceTest5() {
	// 슬라이스 추가, 병합(append), 복사(copy)
	// 배열은 고정된 크기로 그 크기 이상의 데이타를 임의로 추가할 수 없지만, 슬라이스는 자유롭게 새로운 요소를 추가할 수 있다
	// 슬라이스에 새로운 요소를 추가하기 위해서는 Go 내장함수인 append()를 사용한다. append()의 첫 파라미터는 슬라이스 객체이고, 두번째는 추가할 요소의 값이다.
	// 여러 개의 요소 값들을 한꺼번에 추가하기 위해서는 append() 두번째 파라미터 뒤에 계속하여 값을 추가할 수 있다.
	e := []int{1, 2, 3, 4}

	e = append(e, 5, 6, 7, 8)
	fmt.Println(e)
}

func sliceTest6() {
	// 내장함수 append()가 슬라이스에 데이터를 추가할 떄, 내부적으로 발생하는 상황
	//   1> 슬라이스 용량(capacity)이 아직 남아 있는지 체크,
	//   2> 남아 있는 경우는 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가
	//   3> 용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array를 생성
	//   4> 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당
	sliceF := make([]int, 0, 3) // len=0, cap=3인 슬라이스

	for i := 1; i <= 15; i++ {
		sliceF = append(sliceF, i)
		fmt.Printf("[%d] len:%d, cap:%d \n", i, len(sliceF), cap(sliceF))
	}

	fmt.Println(sliceF)
}

func sliceTest7() {
	// 한 슬라이스를 다른 슬라이스 뒤에 병합 : append() 사용
	// append()함수에서는 2개의 슬라이스를 파라미터로 갖는데, 처음 슬라이스 뒤에 두번째 파라미터의 슬라이스를 추가하게 된다
	// 여기서 한가지 주의할 것은 두번째 슬라이스 뒤에 ... 을 붙인다는 것인데, 이 ellipsis(...)는 해당 슬라이스의 컬렉션을 표현하는 것으로 두번째 슬라이스의 모든 요소들의 집합을 나타낸다

	g := []int{1, 2, 3, 4}
	h := []int{5, 6, 7, 8}

	i := append(g, h...)
	fmt.Println(i)
}

func sliceTest8() {
	// Go 슬라이스는 내장함수 copy()를 사용하여 한 슬라이스를 다른 슬라이스로 복사할 수도 있다.

	source := []int{0, 1, 2}
	fmt.Println(source)
	fmt.Println("source >>", len(source), cap(source))
	target := make([]int, len(source), cap(source)*2)

	copy(target, source)

	fmt.Println(target)
	fmt.Println("target >>", len(target), cap(target))
}

func sliceTest9() {
	// 슬라이스는 내부적으로 사용하는 배열의 부분 영역인 세그먼트에 대한 메타 정보를 가지고 있다
	// 슬라이스는 크게 3개의 필드로 구성
	//  1. 첫째 필드는 내부적으로 사용하는 배열에 대한 포인터 정보
	//  2. 두번째는 세그먼트의 길이
	//  3. 세번째는 세그먼트의 최대 용량(Capacity)

	// 처음 슬라이스가 생성될 때, 만약 길이와 용량이 지정되었다면, 내부적으로 용량(Capacity)만큼의 배열을 생성
	// 슬라이스 첫번째 필드에 그 배열의 처음 메모리 위치를 지정
	// 두번째 길이 필드는 지정된 (첫 배열요소로부터의) 길이를 갖게되고, 세번째 용량 필드는 전체 배열의 크기를 갖는다.

}

func main() {
	sliceTest1()

	sliceTest2()

	sliceTest3()

	sliceTest4()

	sliceTest5()

	sliceTest6()

	sliceTest7()

	sliceTest8()

	sliceTest9()
}
