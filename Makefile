ocr.so ocr.h: capture.go image.go main.go upload.go
	go build -buildmode=c-shared -o ocr.so