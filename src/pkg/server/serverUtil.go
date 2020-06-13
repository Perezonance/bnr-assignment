package server

import "net/http"

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		//TODO: ERROR HANDLING
	}
}