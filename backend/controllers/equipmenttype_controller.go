package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/equipmenttype"
	"github.com/gin-gonic/gin"
)

// EquipmenttypeController defines the struct for the equipmenttype controller
type EquipmenttypeController struct {
	client *ent.Client
	router gin.IRouter
}

// Equipmenttype defines the struct for the equipmenttype controller
type Equipmenttype struct {
	EQUIPMENTTYPE string
}

// CreateEquipmenttype handles POST requests for adding equipmenttype entities
// @Summary Create equipmenttype
// @Description Create equipmenttype
// @ID create-equipmenttype
// @Accept   json
// @Produce  json
// @Param equipmenttype body ent.Equipmenttype true "Equipmenttype entity"
// @Success 200 {object} ent.Equipmenttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmenttypes [post]
func (ctl *EquipmenttypeController) CreateEquipmenttype(c *gin.Context) {
	obj := ent.Equipmenttype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Equipmenttype binding failed",
		})
		return
	}

	et, err := ctl.client.Equipmenttype.
		Create().
		SetEQUIPMENTTYPE(obj.EQUIPMENTTYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, et)
}

// GetEquipmenttype handles GET requests to retrieve a equipmenttype entity
// @Summary Get a equipmenttype entity by ID
// @Description get equipmenttype by ID
// @ID get-equipmenttype
// @Produce  json
// @Param id path int true "Equipmenttype ID"
// @Success 200 {object} ent.Equipmenttype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmenttypes/{id} [get]
func (ctl *EquipmenttypeController) GetEquipmenttype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	et, err := ctl.client.Equipmenttype.
		Query().
		Where(equipmenttype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, et)
}

// ListEquipmenttype handles request to get a list of equipmenttype entities
// @Summary List equipmenttype entities
// @Description list equipmenttype entities
// @ID list-equipmenttype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Equipmenttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmenttypes [get]
func (ctl *EquipmenttypeController) ListEquipmenttype(c *gin.Context) {
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

	equipmenttypes, err := ctl.client.Equipmenttype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, equipmenttypes)
}

// NewEquipmenttypeController creates and registers handles for the equipmenttype controller
func NewEquipmenttypeController(router gin.IRouter, client *ent.Client) *EquipmenttypeController {
	etc := &EquipmenttypeController{
		client: client,
		router: router,
	}
	etc.register()
	return etc
}

// InitEquipmenttypeController registers routes to the main engine
func (ctl *EquipmenttypeController) register() {
	equipmenttypes := ctl.router.Group("/equipmenttypes")

	equipmenttypes.GET("", ctl.ListEquipmenttype)

	// CRUD
	equipmenttypes.POST("", ctl.CreateEquipmenttype)
	equipmenttypes.GET(":id", ctl.GetEquipmenttype)
}
