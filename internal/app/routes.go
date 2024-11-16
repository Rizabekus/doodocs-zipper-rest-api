package app

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()
	r.HandleFunc("/zipper")
	r.HandleFunc("/zipper")
	r.HandleFunc("/zipper")
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Started the server")
	defer loggers.CloseLogFile()
	log.Fatal(http.ListenAndServe(os.Getenv("PORT")))
}
