package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/equipmentrental"
	"github.com/G16/app/ent/equipmenttype"
	"github.com/G16/app/ent/member"
	"github.com/gin-gonic/gin"
)

// EquipmentrentalController defines the struct for the equipmentrental controller
type EquipmentrentalController struct {
	client *ent.Client
	router gin.IRouter
}

// Equipmentrental defines the struct for the equipmentrental controller
type Equipmentrental struct {
	MEMBER        int
	EQUIPMENT     int
	EQUIPMENTTYPE int
	EMPLOYEE      int
	RENTALAMOUNT  int
	RENTALDATE    string
	RETURNDATE    string
}

// CreateEquipmentrental handles POST requests for adding equipmentrental entities
// @Summary Create equipmentrental
// @Description Create equipmentrental
// @ID create-equipmentrental
// @Accept   json
// @Produce  json
// @Param equipmentrental body ent.Equipmentrental true "Equipmentrental entity"
// @Success 200 {object} ent.Equipmentrental
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmentrentals [post]
func (ctl *EquipmentrentalController) CreateEquipmentrental(c *gin.Context) {
	obj := Equipmentrental{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Equipmentrental binding failed",
		})
		return
	}

	mb, err := ctl.client.Member.
		Query().
		Where(member.IDEQ(int(obj.MEMBER))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "member not found",
		})
		return
	}

	eq, err := ctl.client.Equipment.
		Query().
		Where(equipment.IDEQ(int(obj.EQUIPMENT))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "equipment not found",
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

	emp, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.EMPLOYEE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	time1, err := time.Parse(time.RFC3339, obj.RENTALDATE)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "วันที่ยืมอุปกรณ์ไม่ถูกต้อง",
		})
		return
	}

	timere, err := time.Parse(time.RFC3339, obj.RETURNDATE)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "วันที่ครบกำหนดยืมอุปกรณ์ไม่ถูกต้อง",
		})
		return
	}

	if obj.RENTALAMOUNT <= 0 {
		c.JSON(400, gin.H{
			"error": "จำนวนยืมอุปกรณ์ไม่ถูกต้อง",
		})
		return
	}

	er, err := ctl.client.Equipmentrental.
		Create().
		SetMember(mb).
		SetEquipment(eq).
		SetEquipmenttype(et).
		SetEmployee(emp).
		SetRENTALAMOUNT(obj.RENTALAMOUNT).
		SetRENTALDATE(time1).
		SetRETURNDATE(timere).
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
		"error":  er,
	})
}

// GetEquipmentrental handles GET requests to retrieve a equipmentrental entity
// @Summary Get a equipmentrental entity by ID
// @Description get equipmentrental by ID
// @ID get-equipmentrental
// @Produce  json
// @Param id path int true "Equipmentrental ID"
// @Success 200 {object} Equipmentrental
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmentrentals/{id} [get]
func (ctl *EquipmentrentalController) GetEquipmentrental(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	er, err := ctl.client.Equipmentrental.
		Query().
		WithMember().
		WithEquipment().
		WithEquipmenttype().
		WithEmployee().
		Where(equipmentrental.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, er)
}

// ListEquipmentrental handles request to get a list of equipmentrental entities
// @Summary List equipmentrental entities
// @Description list equipmentrental entities
// @ID list-equipmentrental
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Equipmentrental
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipmentrentals [get]
func (ctl *EquipmentrentalController) ListEquipmentrental(c *gin.Context) {
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

	Eqrentals, err := ctl.client.Equipmentrental.
		Query().
		WithMember().
		WithEquipment().
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

	c.JSON(200, Eqrentals)
}

// NewEquipmentrentalController creates and registers handles for the equipmentrental controller
func NewEquipmentrentalController(router gin.IRouter, client *ent.Client) *EquipmentrentalController {
	erc := &EquipmentrentalController{
		client: client,
		router: router,
	}
	erc.register()
	return erc

}

// InitEquipmentrentalController registers routes to the main engine
func (ctl *EquipmentrentalController) register() {
	Eqrentals := ctl.router.Group("/equipmentrentals")

	Eqrentals.GET("", ctl.ListEquipmentrental)

	// CRUD
	Eqrentals.POST("", ctl.CreateEquipmentrental)
	Eqrentals.GET(":id", ctl.GetEquipmentrental)
}
