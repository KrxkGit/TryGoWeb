package controller

import (
	"fmt"
	"io"
	"net/http"
)

func registerHandleUploadRoute() {
	http.HandleFunc("/upload", handleUpload)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

	file, _, err = r.FormFile("uploaded") /*自动返回第一个文件*/
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

	// 表单流式处理
	/*mulReader, err := r.MultipartReader()
	if err == nil {
		mulReader.NextPart()

	}*/
}
