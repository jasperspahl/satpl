package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/internal/templates"
)

type PlaylistHandler interface {
	Create(c *gin.Context)
	GetPlaylists(c *gin.Context)
}

type playlistHandler struct {
	playlistService services.PlaylistService
}

func NewPlaylistHandler(playlistService services.PlaylistService) PlaylistHandler {
	return &playlistHandler{playlistService}
}

func (h *playlistHandler) GetPlaylists(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(sessionUserId)
	if userId == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	playlists, err := h.playlistService.GetPlaylist(c, userId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "", templates.Playlists(playlists))
}

type PlaylistCreateBody struct {
	Name   string `form:"name"`
	Public string `form:"public,omitempty"`
}

func (h *playlistHandler) Create(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(sessionUserId)
	if userId == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var playlistArgs PlaylistCreateBody
	err := c.Bind(&playlistArgs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": fmt.Sprintf("unable to bind body: %v", err)})
		return
	}
	playlist, err := h.playlistService.CreatePlaylist(c, userId.(int), playlistArgs.Name, playlistArgs.Public == "on")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "unable to create playlist"})
		return
	}
	c.HTML(http.StatusCreated, "", templates.Playlist(playlist))
}
