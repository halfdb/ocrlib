package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

var api = "http://localhost:8080/file"

func recogniseAndSet(fileBuffer *bytes.Buffer) {
	bodyBuffer := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyBuffer)

	fileWriter, err := writer.CreateFormFile("file", "img.png")
	if err != nil {
		return
	}
	_, err = io.Copy(fileWriter, fileBuffer)
	if err != nil {
		return
	}

	writer.Close()
	resp, err := http.Post(api, writer.FormDataContentType(), bodyBuffer)
	if err != nil {
		fmt.Println("error when sending request")
		return
	}
	defer resp.Body.Close()

	var result struct {
		Result string `json:"result"`
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read response body")
		return
	}
	if err := json.Unmarshal(all, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("cannot unmarshal JSON")
		return
	}
	setResult(result.Result)
}
