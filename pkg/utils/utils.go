package utils

import (
	"encoding/json"
	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
	"net/http"
	"runtime"
)

func GetCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown", 0, "Unknown"
	}

	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return file, line, functionName
}
func SendResponse(msg string, w http.ResponseWriter, statusCode int) {
	response := models.Response{Message: msg}

	responseJSON, err := json.Marshal(response)
	if err != nil {

		resp := models.Response{Message: "Internal Server Error"}
		internalErrorJSON, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
