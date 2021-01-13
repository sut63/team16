package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/age"
	"github.com/gin-gonic/gin"
)

// AgeController defines the struct for the age controller
type AgeController struct {
	client *ent.Client
	router gin.IRouter
}

// Age defines the struct for the age
type Age struct {
	AGE int
}

// CreateAge handles POST requests for adding age entities
// @Summary Create age
// @Description Create age
// @ID create-age
// @Accept   json
// @Produce  json
// @Param age body ent.Age true "Age entity"
// @Success 200 {object} ent.Age
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ages [post]
func (ctl *AgeController) CreateAge(c *gin.Context) {
	obj := ent.Age{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "age binding failed",
		})
		return
	}

	ag, err := ctl.client.Age.
		Create().
		SetAGE(obj.AGE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ag)
}

// GetAge handles GET requests to retrieve a age entity
// @Summary Get a age entity by ID
// @Description get age by ID
// @ID get-age
// @Produce  json
// @Param id path int true "Age ID"
// @Success 200 {object} ent.Age
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ages/{id} [get]
func (ctl *AgeController) GetAge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ag, err := ctl.client.Age.
		Query().
		Where(age.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ag)
}

// ListAge handles request to get a list of age entities
// @Summary List age entities
// @Description list age entities
// @ID list-age
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Age
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ages [get]
func (ctl *AgeController) ListAge(c *gin.Context) {
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

	ages, err := ctl.client.Age.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ages)
}

// NewAgeController creates and registers handles for the age controller
func NewAgeController(router gin.IRouter, client *ent.Client) *AgeController {
	agc := &AgeController{
		client: client,
		router: router,
	}
	agc.register()
	return agc
}

// InitAgeController registers routes to the main engine
func (ctl *AgeController) register() {
	ages := ctl.router.Group("/ages")

	ages.GET("", ctl.ListAge)

	// CRUD
	ages.POST("", ctl.CreateAge)
	ages.GET(":id", ctl.GetAge)

}
