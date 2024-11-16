package app

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	loggers.InitLoggers()
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Loaded the configuration data from .env")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	loggers.InfoLog(file, line+7, "Successfully connected to database")

	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Force(1); err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			loggers.InfoLog(file, line+21, "No new migrations to apply.")
		} else {
			log.Fatal(err)
		}
	} else {
		version, _, err := m.Version()
		if err != nil {
			log.Fatal(err)
		}
		loggers.InfoLog(file, line, fmt.Sprintf("Successfully applied migrations. Current version: %d", version))
	}
	loggers.InfoLog(file, line+32, "Successfully applied migrations")

	storage := storage.StorageInstance(db)
	service := services.ServiceInstance(storage)
	handler := handlers.HandlersInstance(service)

	Routes(handler)

	defer db.Close()
}
