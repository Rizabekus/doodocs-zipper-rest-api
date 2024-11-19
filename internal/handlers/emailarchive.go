package handlers

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/loggers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/utils"
)

func (handler *Handlers) EmailArchive(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "File is required")
		utils.SendResponse("File is required", w, http.StatusBadRequest)
		return
	}
	defer file.Close()

	allowedMIMETypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}
	mimeType := fileHeader.Header.Get("Content-Type")
	if !allowedMIMETypes[mimeType] {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Wrong MIME type", "Wrong MIME type")
		utils.SendResponse("Unsupported file type", w, http.StatusBadRequest)
		return
	}

	emails := r.FormValue("emails")
	if emails == "" {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "No emails", "No emails")
		utils.SendResponse("Emails are required", w, http.StatusBadRequest)
		return
	}
	emailList := strings.Split(emails, ",")
	for i := range emailList {
		emailList[i] = strings.TrimSpace(emailList[i])
	}

	var fileBuffer bytes.Buffer
	_, err = io.Copy(&fileBuffer, file)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Internal Server Error")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	err = utils.SendEmail(emailList, fileHeader.Filename, fileBuffer.Bytes())
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Internal Server Error")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "File successfully sent"}`))
	f, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(f, line+1, r.Method, r.URL.Path, http.StatusOK, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully sent file to emails")

}
