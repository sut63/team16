package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/G16/app/ent"
	"github.com/G16/app/ent/paymenttype"
)

// PaymenttypeController defines the struct for the paymenttype controller
type PaymenttypeController struct {
	client *ent.Client
	router gin.IRouter
}

// Paymenttype defines the struct for the paymenttype controller
type Paymenttype struct {
	TYPE string
}

// CreatePaymenttype handles POST requests for adding paymenttype entities
// @Summary Create paymenttype
// @Description Create paymenttype
// @ID create-paymenttype
// @Accept   json
// @Produce  json
// @Param paymenttype body ent.Paymenttype true "Paymenttype entity"
// @Success 200 {object} ent.Paymenttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes [post]
func (ctl *PaymenttypeController) CreatePaymenttype(c *gin.Context) {
	obj := ent.Paymenttype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype binding failed",
		})
		return
	}

	pt, err := ctl.client.Paymenttype.
		Create().
		SetTYPE(obj.TYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pt)
}

// GetPaymenttype handles GET requests to retrieve a paymenttype entity
// @Summary Get a paymenttype entity by ID
// @Description get paymenttype by ID
// @ID get-paymenttype
// @Produce  json
// @Param id path int true "Paymenttype ID"
// @Success 200 {object} ent.Paymenttype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes/{id} [get]
func (ctl *PaymenttypeController) GetPaymenttype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pt, err := ctl.client.Paymenttype.
		Query().
		Where(paymenttype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pt)
}

// ListPaymenttype handles request to get a list of paymenttype entities
// @Summary List paymenttype entities
// @Description list paymenttype entities
// @ID list-paymenttype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Paymenttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes [get]
func (ctl *PaymenttypeController) ListPaymenttype(c *gin.Context) {
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

	paymenttypes, err := ctl.client.Paymenttype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paymenttypes)
}

// NewPaymenttypeController creates and registers handles for the paymenttype controller
func NewPaymenttypeController(router gin.IRouter, client *ent.Client) *PaymenttypeController {
	ptc := &PaymenttypeController{
		client: client,
		router: router,
	}
	ptc.register()
	return ptc
}

// InitPaymenttypeController registers routes to the main engine
func (ctl *PaymenttypeController) register() {
	paymenttypes := ctl.router.Group("/paymenttypes")

	paymenttypes.GET("", ctl.ListPaymenttype)

	// CRUD
	paymenttypes.POST("", ctl.CreatePaymenttype)
	paymenttypes.GET(":id", ctl.GetPaymenttype)
}
