package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/position"
	"github.com/gin-gonic/gin"
)

// PositionController defines the struct for the position controller
type PositionController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePosition handles POST requests for adding position entities
// @Summary Create position
// @Description Create position
// @ID create-position
// @Accept   json
// @Produce  json
// @Param position body ent.Position true "Position entity"
// @Success 200 {object} ent.Position
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positions [post]
func (ctl *PositionController) CreatePosition(c *gin.Context) {
	obj := ent.Position{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "position binding failed",
		})
		return
	}

	pst, err := ctl.client.Position.
		Create().
		SetPOSITION(obj.POSITION).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pst)
}

// GetPosition handles GET requests to retrieve a position entity
// @Summary Get a position entity by ID
// @Description get position by ID
// @ID get-position
// @Produce  json
// @Param id path int true "Position ID"
// @Success 200 {object} ent.Position
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positions/{id} [get]
func (ctl *PositionController) GetPosition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pst, err := ctl.client.Position.
		Query().
		Where(position.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pst)
}

// ListPosition handles request to get a list of position entities
// @Summary List position entities
// @Description list position entities
// @ID list-position
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Position
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positions [get]
func (ctl *PositionController) ListPosition(c *gin.Context) {
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

	positions, err := ctl.client.Position.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, positions)
}

// NewPositionController creates and registers handles for the Position controller
func NewPositionController(router gin.IRouter, client *ent.Client) *PositionController {
	pstc := &PositionController{
		client: client,
		router: router,
	}
	pstc.register()
	return pstc
}

// InitPositionController registers routes to the main engine
func (ctl *PositionController) register() {
	positions := ctl.router.Group("/positions")

	positions.GET("", ctl.ListPosition)

	// CRUD
	positions.POST("", ctl.CreatePosition)
	positions.GET(":id", ctl.GetPosition)

}
