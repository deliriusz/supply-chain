package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
)

func (hdl *httpHandler) GetLoginChallenge(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login, err := hdl.loginService.GetLoginChallenge(&input); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"nonce": strconv.FormatInt(login.Nonce, 10)})
	}
}

func (hdl *httpHandler) Login(c *gin.Context) {
	var input model.LoginChallenge

	if err := checkLoginRequest(&input, c); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login, err := hdl.loginService.Login(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err)
		return
	} else {
		c.SetCookie(config.COOKIE_SESSIONID, login.SessionId, int(login.TTL), "/", "localhost", true, true)
	}
}

func (hdl *httpHandler) Logout(c *gin.Context) {
	cookie, _ := c.Cookie(config.COOKIE_SESSIONID)

	if err := hdl.loginService.Logout(&model.Login{
		SessionId: cookie,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(config.COOKIE_SESSIONID, "", -1, "/", "localhost", true, true)
}
