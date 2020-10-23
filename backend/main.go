package main

import (
	"context"
	"log"

	"github.com/Varissara/app/controllers"
	"github.com/Varissara/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Username  string
	Useremail string
}

type Educationlevels struct {
	Educationlevel []EducationLevel
}

type EducationLevel struct {
	Educationlevelname string
}

type ScholarshipTypes struct {
	ScholarshipType []Scholarshiptype
}

type Scholarshiptype struct {
	Scholarshiptypename string
}

type Semesters struct {
	Semester []Semester
}

type Semester struct {
	Semestername string
}

type Studyplans struct {
	Studyplan []StudyPlan
}

type StudyPlan struct {
	Studyplanname string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewScholarshipTypeController(v1, client)
	controllers.NewSemesterController(v1, client)
	controllers.NewStudyPlanController(v1, client)
	controllers.NewEducationLevelController(v1, client)
	controllers.NewScholarshipController(v1, client)

	// Set User Data
	users := Users{
		User: []User{
			User{"Varissara Tundiloktrakul", "varissara@gmail.com"},
			User{"kitsada Jaitahan", "kitsada@gmail.com"},
			User{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUseremail(u.Username).
			SetUsername(u.Username).
			Save(context.Background())
	}

	// Set Semester Data
	semesters := []string{"1 / 2563", "2 / 2563","3 / 2563"}
	for _, se := range semesters {
		client.Semester.
			Create().
			SetSemestername(se).
			Save(context.Background())
	}

	// Set ScholarshipType Data
	scholarshiptypes := []string{"ทุนต่อเนื่อง", "ทุนไม่ต่อเนื่อง"}
	for _, st := range scholarshiptypes {
		client.Scholarshiptype.
			Create().
			SetScholarshiptypename(st).
			Save(context.Background())
	}

	// Set EducationLevel Data
	educationlevels := []string{"มัธยมศึกษาตอนปลาย", "ปริญญาตรี"}
	for _, el := range educationlevels {
		client.Educationlevel.
			Create().
			SetEducationlevelname(el).
			Save(context.Background())
	}

	// Set StudyPlan Data
	studyplans := []string{"วิทย์-คณิต", "วิศวกรรมคอมพิวเตอร์"}
	for _, sp := range studyplans {
		client.Studyplan.
			Create().
			SetStudyplanname(sp).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
