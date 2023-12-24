package handlers

import (
	"github.com/IlyaZayats/restus/internal/requests"
	"github.com/IlyaZayats/restus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RestaurantHandlers struct {
	svc    *services.RestaurantService
	engine *gin.Engine
}

func NewRestaurantHandlers(engine *gin.Engine, svc *services.RestaurantService) (*RestaurantHandlers, error) {
	h := &RestaurantHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *RestaurantHandlers) initRoute() {
	h.engine.GET("/restaurant", h.GetRestaurants)
	h.engine.POST("/restaurant/delete", h.DeleteRestaurant)
	h.engine.PUT("/restaurant", h.InsertRestaurant)
	h.engine.POST("/restaurant", h.UpdateRestaurant)
}

func (h *RestaurantHandlers) GetRestaurants(c *gin.Context) {
	restaurants, err := h.svc.GetRestaurants()
	logrus.Debug(restaurants)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get restaurants error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": restaurants})
}

func (h *RestaurantHandlers) DeleteRestaurant(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteRestaurantRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete restaurant request error", "text": ok})
		return
	}

	if err := h.svc.DeleteRestaurant(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete restaurant error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *RestaurantHandlers) InsertRestaurant(c *gin.Context) {

	req, ok := GetRequest[requests.InsertRestaurantRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert restaurant request error", "text": ok})
		return
	}

	if err := h.svc.InsertRestaurant(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert restaurant error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *RestaurantHandlers) UpdateRestaurant(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateRestaurantRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update restaurant request error", "text": ok})
		return
	}

	if err := h.svc.UpdateRestaurant(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update restaurant error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
