package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.FileServer() : 특정 폴더 안에 있는 정적 파일들을 웹서버에서 클라이언트로 그대로 전달할 때 사용
	// http.FileServer() : 해당 디렉토리 내의 모든 정적 리소스를 1대1 매핑하여 그대로 전달
	// 위 매서드는 ServeHTTP() 핸들러처럼 세밀한 제어는 불가능
	fmt.Println("Starting Go file server...")
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.Handle("/static", http.FileServer(http.Dir("www")))

	http.ListenAndServe(":5000", nil)
}
