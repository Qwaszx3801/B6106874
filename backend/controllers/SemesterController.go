package controllers

import (
	"context"
	"strconv"

	"github.com/Varissara/app/ent"
	"github.com/Varissara/app/ent/semester"
	"github.com/gin-gonic/gin"
)

// SemesterController defines the struct for the semester controller
type SemesterController struct {
	client *ent.Client
	router gin.IRouter
}

// GetSemester handles GET requests to retrieve a semester entity
// @Summary Get a semester entity by ID
// @Description get semester by ID
// @ID get-semester
// @Produce  json
// @Param id path int true "Semester ID"
// @Success 200 {object} ent.Semester
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /semesters/{id} [get]
func (ctl *SemesterController) GetSemester(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	se, err := ctl.client.Semester.
		Query().
		Where(semester.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, se)
}

// ListSemester handles request to get a list of semester entities
// @Summary List semester entities
// @Description list semester entities
// @ID list-semester
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Semester
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /semesters [get]
func (ctl *SemesterController) ListSemester(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	semester, err := ctl.client.Semester.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, semester)
}

// NewSemesterController creates and registers handles for the semester controller
func NewSemesterController(router gin.IRouter, client *ent.Client) *SemesterController {
	sec := &SemesterController{
		client: client,
		router: router,
	}

	sec.register()

	return sec

}

func (ctl *SemesterController) register() {
	semester := ctl.router.Group("/semesters")

	semester.GET(":id", ctl.GetSemester)
	semester.GET("", ctl.ListSemester)

}
