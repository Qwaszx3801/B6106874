package controllers

import (
	"context"
	"strconv"

	"github.com/Varissara/app/ent"
	"github.com/Varissara/app/ent/scholarshiptype"
	"github.com/gin-gonic/gin"
)

// ScholarshipTypeController defines the struct for the scholarshiptype controller
type ScholarshipTypeController struct {
	client *ent.Client
	router gin.IRouter
}



// GetScholarshipType handles GET requests to retrieve a scholarshiptype entity
// @Summary Get a scholarshiptype entity by ID
// @Description get scholarshiptype by ID
// @ID get-scholarshiptype
// @Produce  json
// @Param id path int true "ScholarshipType ID"
// @Success 200 {object} ent.Scholarshiptype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiptypes/{id} [get]
func (ctl *ScholarshipTypeController) GetScholarshipType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	st, err := ctl.client.Scholarshiptype.
		Query().
		Where(scholarshiptype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListScholarshipType handles request to get a list of scholarshiptype entities
// @Summary List scholarshiptype entities
// @Description list scholarshiptype entities
// @ID list-scholarshiptype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Scholarshiptype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiptypes [get]
func (ctl *ScholarshipTypeController) ListScholarshipType(c *gin.Context) {
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

	scholarshiptype, err := ctl.client.Scholarshiptype.
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

	c.JSON(200, scholarshiptype)
}

// NewScholarshipTypeController creates and registers handles for the scholarshiptype controller
func NewScholarshipTypeController(router gin.IRouter, client *ent.Client) *ScholarshipTypeController {
	stc := &ScholarshipTypeController{
		client: client,
		router: router,
	}

	stc.register()

	return stc

}

func (ctl *ScholarshipTypeController) register() {
	scholarshiptype := ctl.router.Group("/scholarshiptypes")

	scholarshiptype.GET(":id", ctl.GetScholarshipType)
	scholarshiptype.GET("", ctl.ListScholarshipType)

}
