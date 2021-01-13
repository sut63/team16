package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/salary"
	"github.com/gin-gonic/gin"
)

// SalaryController defines the struct for the salary controller
type SalaryController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSalary handles POST requests for adding salary entities
// @Summary Create salary
// @Description Create salary
// @ID create-salary
// @Accept   json
// @Produce  json
// @Param salary body ent.Salary true "Salary entity"
// @Success 200 {object} ent.Salary
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys [post]
func (ctl *SalaryController) CreateSalary(c *gin.Context) {
	obj := ent.Salary{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "salary binding failed",
		})
		return
	}

	sr, err := ctl.client.Salary.
		Create().
		SetSALARY(obj.SALARY).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sr)
}

// GetSalary handles GET requests to retrieve a salary entity
// @Summary Get a salary entity by ID
// @Description get salary by ID
// @ID get-salary
// @Produce  json
// @Param id path int true "Salary ID"
// @Success 200 {object} ent.Salary
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys/{id} [get]
func (ctl *SalaryController) GetSalary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sr, err := ctl.client.Salary.
		Query().
		Where(salary.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sr)
}

// ListSalary handles request to get a list of salary entities
// @Summary List salary entities
// @Description list salary entities
// @ID list-salary
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Salary
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys [get]
func (ctl *SalaryController) ListSalary(c *gin.Context) {
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

	salarys, err := ctl.client.Salary.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, salarys)
}

// NewSalaryController creates and registers handles for the Salary controller
func NewSalaryController(router gin.IRouter, client *ent.Client) *SalaryController {
	src := &SalaryController{
		client: client,
		router: router,
	}
	src.register()
	return src
}

// InitSalaryController registers routes to the main engine
func (ctl *SalaryController) register() {
	salarys := ctl.router.Group("/salarys")

	salarys.GET("", ctl.ListSalary)

	// CRUD
	salarys.POST("", ctl.CreateSalary)
	salarys.GET(":id", ctl.GetSalary)

}
