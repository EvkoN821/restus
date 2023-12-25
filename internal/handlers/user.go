package handlers

import (
	"github.com/IlyaZayats/restus/internal/requests"
	"github.com/IlyaZayats/restus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandlers struct {
	svc    *services.UserService
	engine *gin.Engine
}

func NewUserHandlers(engine *gin.Engine, svc *services.UserService) (*UserHandlers, error) {
	h := &UserHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *UserHandlers) initRoute() {
	h.engine.POST("/user", h.Login) //
}

func (h *UserHandlers) Login(c *gin.Context) {
	req, ok := GetRequest[requests.LoginRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login request error", "text": ok})
		return
	}
	user, err := h.svc.Login(req.Login, req.Pwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get users error", "text": err.Error()})
		return
	}
	logrus.Debug(user)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": user})
}
