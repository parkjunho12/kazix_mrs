package webservice

import (
	"sync"

	"github.com/kazix-mrs/factory"
)

type (
	WebHandler struct {
		Fac   *factory.Factory
		mutex sync.Mutex
	}
)

func NewWebHandler(fac *factory.Factory) *WebHandler {
	return &WebHandler{Fac: fac}
}
