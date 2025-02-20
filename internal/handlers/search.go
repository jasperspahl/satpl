package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/internal/templates"
)

type SearchHandler interface {
	Search(c *gin.Context)
}

type searchHandler struct {
	searchService services.SearchService
}

func NewSearchHandler(searchService services.SearchService) SearchHandler {
	return &searchHandler{searchService}
}

func (h *searchHandler) Search(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(sessionUserId)
	if userID == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	query := c.Query("q")
	if query == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	artists, err := h.searchService.Search(c, userID.(int), query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "", templates.Artists(artists))
}
