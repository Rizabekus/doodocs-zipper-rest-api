package handlers

import "github.com/Rizabekus/doodocs-zipper-rest-api/internal/services"

type Handlers struct {
	Service *services.Services
}

func HandlersInstance(services *services.Services) *Handlers {
	return &Handlers{Service: services}
}
