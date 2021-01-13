package main

import (
	"context"
	"log"

	"github.com/G16/app/controllers"
	_ "github.com/G16/app/docs"
	"github.com/G16/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Courses  defines the struct for the courses
//---------------------------------------------
type Courses struct {
	Course []Course
}

// Course  defines the struct for the course
type Course struct {
	course string
}

// @title SUT SA Example API Patient
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

	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewAgeController(v1, client)
	controllers.NewBookcourseController(v1, client)
	controllers.NewClassifierController(v1, client)
	controllers.NewCoursecontroller(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewEquipmentController(v1, client)
	controllers.NewEquipmentrentalController(v1, client)
	controllers.NewEquipmenttypeController(v1, client)
	controllers.NewMemberController(v1, client)
	controllers.NewPaymentController(v1, client)
	controllers.NewPaymenttypeController(v1, client)
	controllers.NewPositionController(v1, client)
	controllers.NewPromotionController(v1, client)
	controllers.NewPromotionamountController(v1, client)
	controllers.NewPromotiontimeController(v1, client)
	controllers.NewPromotiontypeController(v1, client)
	controllers.NewSalaryController(v1, client)
	controllers.NewZoneController(v1, client)

	// Set Courses Data
	courses := Courses{
		Course: []Course{
			Course{"สนามแบดมินตัน"},
			Course{"สนามบาส"},
			Course{"สนามบอล"},
			Course{"สนามวอลเลย์บอล"},
			Course{"สนามตะกร้อ"},
		},
	}

	for _, co := range courses.Course {
		client.Course.
			Create().
			SetCOURSE(co.course).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
