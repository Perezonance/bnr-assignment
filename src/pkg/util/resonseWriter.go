package util

import "net/http"

func WriteRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		ErrorLog("Failed to write response", err)
	}
	return
}