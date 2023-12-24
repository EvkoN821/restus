package handlers

import (
	"github.com/IlyaZayats/restus/internal/requests"
	"github.com/IlyaZayats/restus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FoodHandlers struct {
	svc    *services.FoodService
	engine *gin.Engine
}

func NewFoodHandlers(engine *gin.Engine, svc *services.FoodService) (*FoodHandlers, error) {
	h := &FoodHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *FoodHandlers) initRoute() {
	h.engine.GET("/food", h.GetFoods)           //
	h.engine.POST("/food/delete", h.DeleteFood) //
	h.engine.PUT("/food", h.InsertFood)         //
	h.engine.POST("/food", h.UpdateFood)        //
}

func (h *FoodHandlers) GetFoods(c *gin.Context) {
	//req, ok := GetRequest[requests.GetFoodsRequest](c)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "get foods request error", "text": ok})
	//	return
	//}
	foods, err := h.svc.GetFoods()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get foods error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": foods})
}

func (h *FoodHandlers) DeleteFood(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteFoodRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete food request error", "text": ok})
		return
	}

	if err := h.svc.DeleteFood(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete food error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *FoodHandlers) InsertFood(c *gin.Context) {

	req, ok := GetRequest[requests.InsertFoodRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert food request error", "text": ok})
		return
	}

	if err := h.svc.InsertFood(req.Name, req.Info, req.CourseId, req.Weight, req.Price, req.Calories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert food error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *FoodHandlers) UpdateFood(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateFoodRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update food request error", "text": ok})
		return
	}

	if err := h.svc.UpdateFood(req.Name, req.Info, req.Id, req.Weight, req.Price, req.Calories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update food error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
