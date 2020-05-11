package app

type Handlers struct {
	store storage.Storage
}

func (h *Handlers) searchHandler(c *gin.Context) {
	channel := c.Query("channel")
	term := c.Query("term")
	user := c.Query("user")
	dateQuery := c.Query("date")
	date, _ := parseDate(dateQuery)
	messages, err := h.store.Query(channel, term, user, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error"
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}

func parseDate(s string) (time.Time, bool) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 3 {
		return time.Time{}, false
	}
	year := tokens[0]
	month := tokens[1]
	day := tokens[2]

	return time.Date(year, month, dat, 0, 0, 0, 0, time.UTC), true
}