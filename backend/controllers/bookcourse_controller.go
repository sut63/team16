package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/G16/app/ent/bookcourse"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/course"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/member"

	"github.com/gin-gonic/gin"
)

// BookcourseController defines the struct for the bookcourse controller
type BookcourseController struct {
	client *ent.Client
	router gin.IRouter
}

// Bookcourse defines the struct for the bookcourse controller
type Bookcourse struct {
	MEMBER   int
	EMPLOYEE int
	COURSE   int
	BOOKTIME string
}

// CreateBookcourse handles POST requests for adding bookcourse entities
// @Summary Create bookcourse
// @Description Create bookcourse
// @ID create-bookcourse
// @Accept   json
// @Produce  json
// @Param bookcourse body ent.Bookcourse true "Bookcourse entity"
// @Success 200 {object} ent.Bookcourse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookcourses [post]
func (ctl *BookcourseController) CreateBookcourse(c *gin.Context) {
	obj := Bookcourse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookcourse binding failed",
		})
		return
	}

	mem, err := ctl.client.Member.
		Query().
		Where(member.IDEQ(int(obj.MEMBER))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "member not found",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.EMPLOYEE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(obj.COURSE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.BOOKTIME)
	bc, err := ctl.client.Bookcourse.
		Create().
		SetMember(mem).
		SetCourse(co).
		SetEmployee(em).
		SetBOOKTIME(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bc)
}

// GetBookcourse handles GET requests to retrieve a bookcourse entity
// @Summary Get a bookcourse entity by ID
// @Description get bookcourse by ID
// @ID get-bookcourse
// @Produce  json
// @Param id path int true "Bookcourse ID"
// @Success 200 {object} Bookcourse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookcourses/{id} [get]
func (ctl *BookcourseController) GetBookcourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bc, err := ctl.client.Bookcourse.
		Query().
		WithMember().
		WithEmployee().
		Where(bookcourse.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bc)
}

// ListBookcourse handles request to get a list of bookcourse entities
// @Summary List bookcourse entities
// @Description list bookcourse entities
// @ID list-bookcourse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bookcourse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookcourses [get]
func (ctl *BookcourseController) ListBookcourse(c *gin.Context) {
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

	bookcourses, err := ctl.client.Bookcourse.
		Query().
		WithMember().
		WithEmployee().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bookcourses)
}

// NewBookcourseController creates and registers handles for the bookcourse controller
func NewBookcourseController(router gin.IRouter, client *ent.Client) *BookcourseController {
	pmc := &BookcourseController{
		client: client,
		router: router,
	}
	pmc.register()
	return pmc

}

// InitBookcourseController registers routes to the main engine
func (ctl *BookcourseController) register() {
	bookcourses := ctl.router.Group("/bookcourses")

	bookcourses.GET("", ctl.ListBookcourse)

	// CRUD
	bookcourses.POST("", ctl.CreateBookcourse)
	bookcourses.GET(":id", ctl.GetBookcourse)
}
