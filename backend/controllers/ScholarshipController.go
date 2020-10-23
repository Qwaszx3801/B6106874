package controllers

import (
	"context"
	"strconv"

	"github.com/Varissara/app/ent"
	"github.com/Varissara/app/ent/educationlevel"
	"github.com/Varissara/app/ent/scholarship"
	"github.com/Varissara/app/ent/scholarshiptype"
	"github.com/Varissara/app/ent/semester"
	"github.com/Varissara/app/ent/studyplan"
	"github.com/Varissara/app/ent/user"
	"github.com/gin-gonic/gin"
)

// ScholarshipController defines the struct for the scholarship controller
type ScholarshipController struct {
	client *ent.Client
	router gin.IRouter
}

type Scholarship struct {
	ScholarshipName   string
	Organization      string
	UserID            int
	ScholarshipTypeID int
	EducationLevelID  int
	StudyPlanID       int
	SemesterID        int
}

// CreateScholarship handles POST requests for adding scholarship entities
// @Summary Create scholarship
// @Description Create scholarship
// @ID create-scholarship
// @Accept   json
// @Produce  json
// @Param scholarship body Scholarship true "Scholarship entity"
// @Success 200 {object} ent.Scholarship
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarships [post]
func (ctl *ScholarshipController) CreateScholarship(c *gin.Context) {
	obj := Scholarship{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "scholarship binding failed",
		})
		return
	}

	st, err := ctl.client.Scholarshiptype.
		Query().
		Where(scholarshiptype.IDEQ(int(obj.ScholarshipTypeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "scholarshiptypeid not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "userid not found",
		})
		return
	}

	el, err := ctl.client.Educationlevel.
		Query().
		Where(educationlevel.IDEQ(int(obj.EducationLevelID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "educationlevelid not found",
		})
		return
	}

	sp, err := ctl.client.Studyplan.
		Query().
		Where(studyplan.IDEQ(int(obj.StudyPlanID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "studyplanid not found",
		})
		return
	}

	se, err := ctl.client.Semester.
		Query().
		Where(semester.IDEQ(int(obj.SemesterID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "semesterid not found",
		})
		return
	}

	s, err := ctl.client.Scholarship.
		Create().
		SetSchlolarshipname(obj.ScholarshipName).
		SetOrganization(obj.Organization).
		SetUserID(u).
		SetScholarshiptypeID(st).
		SetEducationlevelID(el).
		SetStudyplanID(sp).
		SetSemesterID(se).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   s,
	})
}

// ListScholarship handles request to get a list of scholarship entities
// @Summary List scholarship entities
// @Description list scholarship entities
// @ID list-scholarship
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Scholarship
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarships [get]
func (ctl *ScholarshipController) ListScholarship(c *gin.Context) {
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

	scholarship, err := ctl.client.Scholarship.
		Query().
		WithUserID().
		WithScholarshiptypeID().
		WithEducationlevelID().
		WithStudyplanID().
		WithSemesterID().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, scholarship)
}

// GetScholarship handles GET requests to retrieve a scholarship entity
// @Summary Get a scholarship entity by ID
// @Description get scholarship by ID
// @ID get-scholarship
// @Produce  json
// @Param id path int true "Scholarship ID"
// @Success 200 {object} ent.Scholarship
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarships/{id} [get]
func (ctl *ScholarshipController) GetScholarship(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Scholarship.
		Query().
		Where(scholarship.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// NewScholarshipController creates and registers handles for the scholarship controller
func NewScholarshipController(router gin.IRouter, client *ent.Client) *ScholarshipController {
	sc := &ScholarshipController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

// InitScholarshipController registers routes to the main engine
func (ctl *ScholarshipController) register() {
	scholarships := ctl.router.Group("/scholarships")

	scholarships.POST("", ctl.CreateScholarship)
	scholarships.GET("", ctl.ListScholarship)
	scholarships.GET(":id", ctl.GetScholarship)

}
