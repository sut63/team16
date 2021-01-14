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
	COURSE string
}

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	EMPLOYEEID      string
	EMPLOYEENAME    string
	EMPLOYEEADDRESS string
	IDCARDNUMBER    string
}

// Ages  defines the struct for the ages
type Ages struct {
	Age []Age
}

// Age  defines the struct for the age
type Age struct {
	AGE int
}

// Positions  defines the struct for the positions
type Positions struct {
	Position []Position
}

// Position  defines the struct for the position
type Position struct {
	POSITION string
}

// Salarys  defines the struct for the salarys
type Salarys struct {
	Salary []Salary
}

// Classifiers  defines the struct for the classifiers
//---------------------------------------------
type Classifiers struct {
	Classifier []Classifier
}

// Classifier  defines the struct for the classifier
type Classifier struct {
	EQUIPMENTCLASSIFIER string
}

// Equipmenttypes  defines the struct for the Equipmenttypes
//---------------------------------------------
type Equipmenttypes struct {
	Equipmenttype []Equipmenttype
}

// Equipmenttype  defines the struct for the Equipmenttype
type Equipmenttype struct {
	EQUIPMENTTYPE string
}

// Zones  defines the struct for the Zones
//---------------------------------------------
type Zones struct {
	Zone []Zone
}

// Zone  defines the struct for the Zone
type Zone struct {
	EQUIPMENTZONE string
}

// Promotiontypes  defines the struct for the Promotiontypes
//---------------------------------------------
type Promotiontypes struct {
	Promotiontype []Promotiontype
}

// Promotiontype  defines the struct for the Promotiontype
type Promotiontype struct {
	TYPE string
}

// Promotionamounts  defines the struct for the Promotionamounts
//---------------------------------------------
type Promotionamounts struct {
	Promotionamount []Promotionamount
}

// Promotionamount  defines the struct for the Promotionamount
type Promotionamount struct {
	AMOUNT int
}

// Members  defines the struct for the Members
//---------------------------------------------
type Members struct {
	Member []Member
}

// Member  defines the struct for the Member
type Member struct {
	MEMBERID	string
	MAMBERNAME	string
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
	controllers.NewPromotiontypeController(v1, client)
	controllers.NewSalaryController(v1, client)
	controllers.NewZoneController(v1, client)

	// Set Member Data
	members := Members{
		Member: []Member{
			Menber{"B6011","โทน"},
			Menber{"B6022","ที"},
			Menber{"B6033","ตัง"},
			Menber{"B6044","รัก"},
			Menber{"B6055","ปลื้ม"},
		},
	}

	for _, mb := range members.Membere {
		client.Member.
			Create().
			SetMEMBERID(mb.MEMBERID).
			SetMEMBERNAME(mb.MEMBERNAME).
			Save(context.Background())
	}

	// Set Promotiontype Data
	promotiontypes := Promotiontypes{
		Promotiontype: []Promotiontype{
			Promotiontype{"ส่วนลด"},
			Promotiontype{"ลุ้นโชค"},
			Promotiontype{"สะสมแต้ม"},
			Promotiontype{"การใช้งาน"},
			Promotiontype{"อายุสมาชิก"},
		},
	}

	for _, prt := range promotiontypes.Promotiontype {
		client.Promotiontype.
			Create().
			SetTYPE(prt.TYPE).
			Save(context.Background())
	}

	// Set Promotionamount Data
	promotionamounts := Promotionamounts{
		Promotionamount: []Promotionamount{
			Promotionamount{1},
			Promotionamount{3},
			Promotionamount{5},
			Promotionamount{7},
			Promotionamount{9},
		},
	}

	for _, prm := range promotionamounts.Promotionamount {
		client.Promotionamount.
			Create().
			SetAMOUNT(prm.AMOUNT).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"EM01", "Tee", "อุดรธานี", "1234567892581"},
			Employee{"EM02", "Tone", "พัทลุง", "1234567892582"},
			Employee{"EM03", "Tang", "ชลบุรี", "1234567892583"},
			Employee{"EM04", "Rak", "นครราชสีมา", "1234567892584"},
			Employee{"EM05", "Jo", "บุรีรัมย์", "1234567892585"},
		},
	}

	for _, em := range employee.Employee {
		client.Employee.
			Create().
			SetEMPLOYEEID(em.EMPLOYEEID).
			SetEMPLOYEENAME(em.EMPLOYEENAME).
			SetEMPLOYEEADDRESS(em.EMPLOYEEADDRESS).
			SetIDCARDNUMBER(em.IDCARDNUMBER).
			Save(context.Background())
	}

	// Set Ages Data
	ages := Ages{
		Age: []Age{
			Age{22},
			Age{23},
			Age{24},
			Age{25},
			Age{26},
			Age{27},
			Age{28},
			Age{29},
			Age{30},
		},
	}
	for _, ag := range age.Age {
		client.Age.
			Create().
			SetAGE(ag.age).
			Save(context.Background())
	}

	// Set Positions Data
	positions := Positions{
		Position: []Position{
			Position{"เจ้าหน้าที่รักษาความปลอดภัย"},
			Position{"พนักงานทำความสะอาด"},
			Position{"พนักงานเคาน์เตอร์"},
			Position{"ผู้จัดการ"},
			Position{"เทรนเนอร์"},
		},
	}
	for _, pst := range position.Position {
		client.Position.
			Create().
			SetPOSITION(pst.position).
			Save(context.Background())
	}

	// Set Salarys Data
	salarys := Salarys{
		Salary: []Salary{
			Salary{15000},
			Salary{17000},
			Salary{19000},
			Salary{21000},
			Salary{23000},
			Salary{25000},
			Salary{30000},
		},
	}

	for _, sr := range salary.Salary {
		client.Salary.
			Create().
			SetSALARY(sr.salary).
			Save(context.Background())
	}

	// Set Classifiers Data
	classifiers := Classifiers{
		Classifier: []Classifier{
			Classifier{"ชั้น"},
			Classifier{"อัน"},
			Classifier{"เครื่อง"},
		},
	}

	for _, ecf := range classifiers.Classifier {
		client.Classifier.
			Create().
			SetEQUIPMENTCLASSIFIER(ecf.EQUIPMENTCLASSIFIER).
			Save(context.Background())
	}

	// Set Equipmenttypes Data
	equipmenttypes := Equipmenttypes{
		Equipmenttype: []Equipmenttype{
			Equipmenttype{"เเบต"},
			Equipmenttype{"กีฑา"},
			Equipmenttype{"ว่ายน้ำ"},
		},
	}

	for _, et := range equipmenttypes.Equipmenttype {
		client.Equipmenttype.
			Create().
			SetEQUIPMENTTYPE(et.EQUIPMENTTYPE).
			Save(context.Background())
	}

	// Set Zones Data
	zones := Zones{
		Zone: []Zone{
			Zone{"A"},
			Zone{"B"},
			Zone{"C"},
			Zone{"D"},
			Zone{"E"},
		},
	}

	for _, zn := range zones.Zone {
		client.Zone.
			Create().
			SetEQUIPMENTZONE(zn.EQUIPMENTZONE).
			Save(context.Background())
	}

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
			SetCOURSE(co.COURSE).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
