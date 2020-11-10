package main

func main() {
	// 간단한 웹 서버를 만들경우 net/http 패키지를 사용하면 되지만, 다양한 기능을 요하는 경우 웹프레임워크를 사용하는 것이 좋다
	// Web Framework는 Request와 Handler를 매핑하는 Routing기능, Request 파라미터들을 Handler의 파라미터에 연결하는 바인딩 기능
	// Request의 상태를 유지하는 컨텍스트 기능, 핸들러에서 자주 쓰이는 공통된 기능을 제공하는 미들웨어 기능 등 다양한 기능이 있다.

	// 현재 사용되는 웹 프레임워크 : Revel, Beego, Martini, Gin, GoCraft, Traffic, Gorilla, Echo
}
