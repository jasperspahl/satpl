package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/internal/templates"
)

type HomepageHandler interface {
	Home(c *gin.Context)
	Config(c *gin.Context)
}

type homepageHandler struct {
	userService services.UserService
}

func NewHomepageHandler(userService services.UserService) HomepageHandler {
	return &homepageHandler{userService}
}

func (h *homepageHandler) Home(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(sessionUserId)
	if userId == nil {
		c.HTML(http.StatusOK, "", templates.Home())
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/config")
}

func (h *homepageHandler) Config(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(sessionUserId)
	if userId == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	user, err := h.userService.GetUserByID(c, userId.(int))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(http.StatusOK, "", templates.LoggedIn(user.DisplayName))
}
