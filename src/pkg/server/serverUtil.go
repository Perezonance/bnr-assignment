package server

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/util"
	"net/http"
)

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		util.ErrorLog("Failed to write response", err)
	}
	return
}