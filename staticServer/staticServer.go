package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	// Static 파일 Handler
	// HTML, 이미지, js 등 정적 파일이 요청되었을 때, 웹서버에서 해당 Static 파일을 적절한 헤더와 함께 전달하는 역할
	// 파일 내용 전달시, 디폴트로 text/plain으로 전송되므로, 이를 막기 위해 파일 확장자에 따른 Content-Type을 헤더에 추가해야한다.

	http.Handle("/", new(staticHandler))

	http.ListenAndServe(":5000", nil)

}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "www" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)

	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string

	fileExt := filepath.Ext(localPath)

	switch fileExt {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
