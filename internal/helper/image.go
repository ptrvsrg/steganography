package helper

import (
	"mime"
	"net/http"
)

var (
	pngContentType  = mime.TypeByExtension(".png")
	jpgContentType  = mime.TypeByExtension(".jpg")
	jpegContentType = mime.TypeByExtension(".jpeg")
)

func GetContentType(content []byte) string {
	return http.DetectContentType(content)
}

func IsPngContentType(contentType string) bool {
	return contentType == pngContentType
}

func IsJpgContentType(contentType string) bool {
	return contentType == jpgContentType
}

func IsJpegContentType(contentType string) bool {
	return contentType == jpegContentType
}
