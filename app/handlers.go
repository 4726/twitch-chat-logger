package app

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/4726/twitch-chat-logger/storage"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	store storage.Storage
}

func (h *Handlers) searchHandler(c *gin.Context) {
	var opts storage.QueryOptions
	c.BindQuery(&opts)
	messages, err := h.store.Query(opts)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}

func (h *Handlers) Close() error {
	return h.store.Close()
}

func parseDate(s string) (time.Time, bool) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 3 {
		return time.Time{}, false
	}
	yearS := tokens[0]
	year, err := strconv.Atoi(yearS)
	if err != nil {
		return time.Time{}, false
	}
	monthS := tokens[1]
	month, err := strconv.Atoi(monthS)
	if err != nil {
		return time.Time{}, false
	}
	dayS := tokens[2]
	day, err := strconv.Atoi(dayS)
	if err != nil {
		return time.Time{}, false
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), true
}
