package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/G16/app/ent/classifier"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/equipmenttype"
	"github.com/G16/app/ent/zone"

	"github.com/gin-gonic/gin"
)

// EquipmentController defines the struct for the equipment controller
type EquipmentController struct {
	client *ent.Client
	router gin.IRouter
}

// Equipment defines the struct for the equipment controller
type Equipment struct {
	EQUIPMENTCLASSIFIER int
	EQUIPMENTTYPE       int
	EQUIPMENTZONE       int
	EMPLOYEE            int
	EQUIPMENTNO         int
	EQUIPMENTNAME       string
	EQUIPMENTAMOUNT     int
	EQUIPMENTDATE       string
	EQUIPMENTDETAIL     string
}

// CreateEquipment handles POST requests for adding equipment entities
// @Summary Create equipment
// @Description Create equipment
// @ID create-equipment
// @Accept   json
// @Produce  json
// @Param equipment body ent.Equipment true "Equipment entity"
// @Success 200 {object} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments [post]
func (ctl *EquipmentController) CreateEquipment(c *gin.Context) {
	obj := Equipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "equipment binding failed",
		})
		return
	}

	cf, err := ctl.client.Classifier.
		Query().
		Where(classifier.IDEQ(int(obj.EQUIPMENTCLASSIFIER))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Classifier not found",
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

	zn, err := ctl.client.Zone.
		Query().
		Where(zone.IDEQ(int(obj.EQUIPMENTZONE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "zone not found",
		})
		return
	}

	et, err := ctl.client.Equipmenttype.
		Query().
		Where(equipmenttype.IDEQ(int(obj.EQUIPMENTTYPE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "equipmenttype not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.EQUIPMENTDATE)
	eq, err := ctl.client.Equipment.
		Create().
		SetClassifier(cf).
		SetEquipmenttype(et).
		SetEmployee(em).
		SetZone(zn).
		SetEQUIPMENTDETAIL(obj.EQUIPMENTDETAIL).
		SetEQUIPMENTNAME(obj.EQUIPMENTNAME).
		SetEQUIPMENTAMOUNT(obj.EQUIPMENTAMOUNT).
		SetEQUIPMENTDATE(time).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"error":  eq,
	})
}

// GetEquipment handles GET requests to retrieve a equipment entity
// @Summary Get a equipment entity by ID
// @Description get equipment by ID
// @ID get-equipment
// @Produce  json
// @Param id path int true "Equipment ID"
// @Success 200 {object} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments/{id} [get]
func (ctl *EquipmentController) GetEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	eq, err := ctl.client.Equipment.
		Query().
		WithEmployee().
		Where(equipment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, eq)
}

// GetEquipmentbyEmployee handles GET requests to retrieve a GetEquipmentbyEmployee entity
// @Summary Get a GetEquipmentbyEmployee entity by ID
// @Description get GetEquipmentbyEmployee by ID
// @ID get-GetEquipmentbyEmployee
// @Produce  json
// @Param id path int true "GetEquipmentbyEmployee ID"
// @Success 200 {array} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmentbyemployees/{id} [get]
func (ctl *EquipmentController) GetEquipmentbyEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Equipment.
		Query().
		WithClassifier().
		WithEquipmenttype().
		WithEmployee().
		WithZone().
		Where(equipment.HasEmployeeWith(employee.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListEquipment handles request to get a list of equipment entities
// @Summary List equipment entities
// @Description list equipment entities
// @ID list-equipment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments [get]
func (ctl *EquipmentController) ListEquipment(c *gin.Context) {
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

	equipments, err := ctl.client.Equipment.
		Query().
		WithClassifier().
		WithZone().
		WithEquipmenttype().
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

	c.JSON(200, equipments)
}

// NewEquipmentController creates and registers handles for the equipment controller
func NewEquipmentController(router gin.IRouter, client *ent.Client) *EquipmentController {
	eqc := &EquipmentController{
		client: client,
		router: router,
	}
	eqc.register()
	return eqc

}

// InitEquipmentController registers routes to the main engine
func (ctl *EquipmentController) register() {
	equipments := ctl.router.Group("/equipments")
	equipmentss := ctl.router.Group("/equipmentbyemployees")

	equipments.GET("", ctl.ListEquipment)
	equipmentss.GET(":id", ctl.GetEquipmentbyEmployee)

	// CRUD
	equipments.POST("", ctl.CreateEquipment)
	equipments.GET(":id", ctl.GetEquipment)
}
