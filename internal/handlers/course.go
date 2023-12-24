package handlers

import (
	"github.com/IlyaZayats/restus/internal/requests"
	"github.com/IlyaZayats/restus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseHandlers struct {
	svc    *services.CourseService
	engine *gin.Engine
}

func NewCourseHandlers(engine *gin.Engine, svc *services.CourseService) (*CourseHandlers, error) {
	h := &CourseHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *CourseHandlers) initRoute() {
	h.engine.GET("/course", h.GetCourses)           //
	h.engine.POST("/course/delete", h.DeleteCourse) //
	h.engine.PUT("/course", h.InsertCourse)         //
	h.engine.POST("/course", h.UpdateCourse)        //
}

func (h *CourseHandlers) GetCourses(c *gin.Context) {
	//req, ok := GetRequest[requests.GetCoursesRequest](c)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "get courses request error", "text": ok})
	//	return
	//}

	courses, err := h.svc.GetCourses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get courses error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": courses})
}

func (h *CourseHandlers) DeleteCourse(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteCourseRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete course request error", "text": ok})
		return
	}

	if err := h.svc.DeleteCourse(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete course error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *CourseHandlers) InsertCourse(c *gin.Context) {

	req, ok := GetRequest[requests.InsertCourseRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert course request error", "text": ok})
		return
	}

	if err := h.svc.InsertCourse(req.RestaurantId, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert course error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *CourseHandlers) UpdateCourse(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateCourseRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update course request error", "text": ok})
		return
	}

	if err := h.svc.UpdateCourse(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update course error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
