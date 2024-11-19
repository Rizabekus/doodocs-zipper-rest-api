package main

import "github.com/Rizabekus/doodocs-zipper-rest-api/internal/app"

func main() {
	app.Run()
}

// go build -tags netgo -ldflags '-s -w' -o app ./cmd/web
