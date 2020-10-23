package controllers

import (
	"context"
	"strconv"

	"github.com/Varissara/app/ent"
	"github.com/Varissara/app/ent/educationlevel"
	"github.com/gin-gonic/gin"
)

// EducationLevelController defines the struct for the educationlevel controller
type EducationLevelController struct {
	client *ent.Client
	router gin.IRouter
}

// GetEducationLevel handles GET requests to retrieve a educationlevel entity
// @Summary Get a educationlevel entity by ID
// @Description get educationlevel by ID
// @ID get-educationlevel
// @Produce  json
// @Param id path int true "EducationLevel ID"
// @Success 200 {object} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels/{id} [get]
func (ctl *EducationLevelController) GetEducationLevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	el, err := ctl.client.Educationlevel.
		Query().
		Where(educationlevel.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, el)
}

// ListEducationLevel handles request to get a list of educationlevel entities
// @Summary List educationlevel entities
// @Description list educationlevel entities
// @ID list-educationlevel
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels [get]
func (ctl *EducationLevelController) ListEducationLevel(c *gin.Context) {
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

	educationlevel, err := ctl.client.Educationlevel.
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

	c.JSON(200, educationlevel)
}

// NewEducationLevelController creates and registers handles for the educationlevel controller
func NewEducationLevelController(router gin.IRouter, client *ent.Client) *EducationLevelController {
	elc := &EducationLevelController{
		client: client,
		router: router,
	}

	elc.register()

	return elc

}

func (ctl *EducationLevelController) register() {
	educationlevel := ctl.router.Group("/educationlevels")

	educationlevel.GET(":id", ctl.GetEducationLevel)
	educationlevel.GET("", ctl.ListEducationLevel)

}
