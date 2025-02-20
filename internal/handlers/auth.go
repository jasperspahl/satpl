package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/internal/spotify"
	"github.com/jasperspahl/satpl/internal/utils"
	"golang.org/x/oauth2"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) AuthHandler {
	return &authHandler{
		userService: userService,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	state := utils.RandomString(16)
	session.Set(sessionState, state)
	if err := session.Save(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	url := spotify.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *authHandler) Callback(c *gin.Context) {
	state := c.Query("state")
	session := sessions.Default(c)
	saved_state := session.Get(sessionState)
	if state == "" || saved_state == nil || state != saved_state {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": "Invalid State"})
		return
	}
	code := c.Query("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": "No code in callback URL"})
		return
	}

	token, err := spotify.Exchange(c, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": "Token exchange failed"})
		return
	}

	client := spotify.Client(c, token)
	user, err := client.GetCurrentUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "couldn't get user information"})
		return
	}
	userID, err := h.userService.LoginUser(c, user, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "unable to login"})
		return
	}
	session.Set(sessionUserId, userID)
	session.Set(sessionAccessToken, token.AccessToken)
	err = session.Save()
	if err != nil {
		fmt.Printf("Failed to save session: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "unable to save session", "innerError": err})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *authHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear session"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
