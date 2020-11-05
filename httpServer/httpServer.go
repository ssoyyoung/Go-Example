package main

import "net/http"

// http.Handler 인터페이스를 갖는 struct 정의
type testHandler struct {
	http.Handler
}

// 위에서 정의한 struct의 매서드 ServeHTTP() 구현
// ServeHTTP() : HTTP Response에 데이터를 쓰기위한 Writer와 HTTP Request 입력데이터를 파라미터로 갖는다
func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

func main() {
	// net/http패키지는 웹 관련 서버/클라이언트 기능을 제공
	// HTTP 서버를 만들기 위한 중요 매서드는 ListenAndServe(), Handle(), HandleFunc() 가 있다.
	// ListenAndServe() : 지정된 포트에 웹 서버를 열고 클라이언트의 Requests를 받아들려 새 Go루틴에 작업을 할당
	// Handle(), HandleFunc() : 요청된 Request Path에 어떤 Request 핸들러를 사용할 지 지정하는 라우팅 역할

	// ============= HandleFunc 사용 =============
	// /hello path에 대한 익명함수 실행
	// http.ResponseWriter는 HTTP Response에 무엇인가 쓸 수 있게 함
	// http.Request는 입력된 Request 요청을 검토
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// ============= Handle 사용 =============
	// Http handler를 정의하는 또 다른 방식
	// 첫번째 파라미터로 URL/URL패턴, 두번째 파라미터로 http.Handler 인터페이스를 갖는 객체
	// Handle()의 두번째 파라미터는 testHandler 객체를 new()함수로 생성하여 전달
	http.Handle("/", new(testHandler))

	// http.ListenAndServe의 첫번째 파라미터는 포트지정, 두번째 파라미터는 어떤 ServeMux를 사용할지
	// nil인경우 DefaultServeMuxfmf를 사용
	// DefaultServeMux를 사용하는 경우, Handle() 혹은 HandleFunc()을 사용하여 라우팅 패턴을 추가
	http.ListenAndServe(":5000", nil)
}
