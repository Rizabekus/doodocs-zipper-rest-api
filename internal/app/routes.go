package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/handlers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/loggers"
	"github.com/Rizabekus/doodocs-zipper-rest-api/pkg/utils"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()
	r.HandleFunc("/api/archive/information", h.GetArchiveInfo).Methods("POST")
	r.HandleFunc("/api/archive/files", h.CreateArchive).Methods("POST")
	r.HandleFunc("/api/archive/mail", h.EmailArchive).Methods("POST")
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Started the server")
	defer loggers.CloseLogFile()
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}

// Расписать роуты
// Тесты
// Сваггер
// Запустить на рендер ком
// Записать видео
