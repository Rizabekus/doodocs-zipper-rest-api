package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/custom_errors"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/loggers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/utils"
)

func (handler *Handlers) CreateArchive(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
		utils.SendResponse("Bad Request", w, http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files[]"]
	if len(files) == 0 {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "No files uploaded", "Failed to unmarshal JSON")
		utils.SendResponse("No files uploaded", w, http.StatusBadRequest)
		return
	}
	buf, err := handler.Service.ZipperService.CreateArchive(files)
	if err != nil {

		var status int
		if errors.Is(err, custom_errors.ErrWrongMIMEType) {

			status = http.StatusBadRequest
			file, line, _ := utils.GetCallerInfo()
			loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, status, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
			utils.SendResponse("Wrong type of file", w, status)
			return
		} else {
			status = http.StatusInternalServerError
			file, line, _ := utils.GetCallerInfo()
			loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, status, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
			utils.SendResponse("Internal Server Error", w, status)
			return
		}

	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buf.Bytes())
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	f, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(f, line+1, r.Method, r.URL.Path, http.StatusCreated, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully created archive file")

}
