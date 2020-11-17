package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Go 루틴
// Go 런타임이 관리하는 논리적 쓰레드이다.
// GO에서 go라는 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행
// goroutine은 비동기적으로 함수루틴을 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.
// goroutine은 OS쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위해 만든것으로, Go 런타임이 자체 관리
// 여러 goroutine들은 종종 하나의 OS 쓰레드 1개로 실행되곤 한다.
// 즉, Go루틴은 OS쓰레드와 1대 1로 대응되지 않고, Multiplexing을 훨씬 적은 OS쓰레드 사용
// 메모리 측면에서도, OS쓰레드가 1메가 바이트 스택을 갖는반면에
// goroutine은 훨씬 작은 몇 킬로바이트의 스택을 갖는다(동적증가)
// Go 런타임은 goroutine을 관리하면서, Go채널을 통해 Go루틴간의 통신을 쉽게할 수 있게 함

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func goroutineTest1() {
	// 함수를 동기적으로 실행
	say("Sync")
	// > 동기적 호출은 say 함수가 완전히 끝났을 때 다음 문장으로 이동한다.

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")
	// > go say() 비동기 호출은 별도의 Go루틴들에서 동작하면서, 메인루틴은 계속 다음 문장(여기서는 time.Sleep)을 실행한다

	// 3초 대기
	time.Sleep(time.Second * 3)
}

func goroutineTest2() {

	// 익명함수 Go루틴
	// Go 루틴은 익명함수에 대해서도 사용할 수 있다.
	// 즉, go 키워드 뒤에 익명함수를 정의하는 것, 이는 익명함수를 비동기적으로 실행한다.

	// WaitGroup 생성. 2개의 Goroutine을 기다린다.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 Goroutine
	go func() {
		defer wait.Done() // 함수가 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	fmt.Println("between")

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 함수가 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // Go루틴이 모두 끝날 때 까지 대기
	fmt.Println("go routine done")

	// 여기서 sync.WaitGroup을 사용하고 있는데, 이는 기본적으로 여러 Go루틴들이 끝날때까지 기다리는 역할
	// WaitGroup을 사용하기 위해서는 먼저 Add()메소드에 몇개의 Go루틴을 기다릴것인지 지정
	// 그후, 각 Go루틴에서 Done()메서드 호출.
	// 메인루틴에서는 Wait() 메서드를 사용하여 Go루틴이 끝나기를 기다린다.
}

func goroutineTest3() {
	// 다중 CPU 처리
	// Go에서는 디폴트로 1개의 CPU를 사용한다.
	// 즉, 여러개의 Go루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다(Concurrent 처리)
	// 만약 머신이 복수개의 CPU를 가진경우,
	// Go 프로그램을 다중 CPU에서 병렬처리(Parrallel 처리)하게 할 수 있는데,
	// 병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다
	// (여기서 CPU 수는 Logical CPU 수를 가리킨다).

	runtime.GOMAXPROCS(4)
	fmt.Println("4개의 cpu 사용")
}

func main() {

	goroutineTest3()

	goroutineTest1()

	goroutineTest2()

}

// Concurrency/Concurrent 와 Parallelism/Parallel 은 전혀 다른 개념이다.
// 프로그램에서 Concurrency는 독립적으로 실행되는 프로세스의 구성, 동시성은 한 번에 많은 것을 처리하는 것
// Parallelism은 계산을 동시에 실행하는 것이다. 병렬 처리는 한 번에 많은 작업을 수행하는 것
