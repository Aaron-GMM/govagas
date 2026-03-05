package openingHandler

import (
	"github.com/Aaron-GMM/govagas/config"
	"github.com/Aaron-GMM/govagas/service"
)

var logger *config.Logger

func init() {
	logger = config.GetLogger("handler")
}

type OpeningHandler struct {
	service service.OpeningService
}

func NewOpeningHandler(s service.OpeningService) *OpeningHandler {
	return &OpeningHandler{
		service: s,
	}
}
