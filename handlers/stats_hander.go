package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

// StatsHandler ...
type StatsHandler struct {
	Stats  models.RollStats
	Logger httputils.Logger
}

// ServeHTTP ...
func (statsHandler StatsHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	statsCopy := statsHandler.Stats.CopyData()
	httputils.HandleJSON(writer, statsHandler.Logger, statsCopy)
}