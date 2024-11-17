package handlers

import "net/http"

func (handler *Handlers) GetArchiveInfo(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)

	if err := r.ParseMultipartForm(10 << 20); err != nil { // limit: 10MB
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()
}
