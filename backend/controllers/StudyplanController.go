package controllers

import (
	"context"
	"strconv"

	"github.com/Varissara/app/ent"
	"github.com/Varissara/app/ent/studyplan"
	"github.com/gin-gonic/gin"
)

// StudyPlanController defines the struct for the studyplan controller
type StudyPlanController struct {
	client *ent.Client
	router gin.IRouter
}

// GetStudyPlan handles GET requests to retrieve a studyplan entity
// @Summary Get a studyplan entity by ID
// @Description get studyplan by ID
// @ID get-studyplan
// @Produce  json
// @Param id path int true "StudyPlan ID"
// @Success 200 {object} ent.Studyplan
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /studyplans/{id} [get]
func (ctl *StudyPlanController) GetStudyPlan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	sp, err := ctl.client.Studyplan.
		Query().
		Where(studyplan.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sp)
}

// ListStudyPlan handles request to get a list of studyplan entities
// @Summary List studyplan entities
// @Description list studyplan entities
// @ID list-studyplan
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Studyplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /studyplans [get]
func (ctl *StudyPlanController) ListStudyPlan(c *gin.Context) {
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

	studyplan, err := ctl.client.Studyplan.
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

	c.JSON(200, studyplan)
}

// NewStudyPlanController creates and registers handles for the studyplan controller
func NewStudyPlanController(router gin.IRouter, client *ent.Client) *StudyPlanController {
	spc := &StudyPlanController{
		client: client,
		router: router,
	}

	spc.register()

	return spc

}

func (ctl *StudyPlanController) register() {
	studyplan := ctl.router.Group("/studyplans")

	studyplan.GET(":id", ctl.GetStudyPlan)
	studyplan.GET("", ctl.ListStudyPlan)

}
