package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/G16/app/ent/payment"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/member"
	"github.com/G16/app/ent/paymenttype"
	"github.com/G16/app/ent/promotion"
	"github.com/gin-gonic/gin"
)

// PaymentController defines the struct for the payment controller
type PaymentController struct {
	client *ent.Client
	router gin.IRouter
}

// Payment defines the struct for the payment controller
type Payment struct {
	MEMBER        int
	PROMOTION     int
	PAYMENTTYPE   int
	EMPLOYEE      int
	PAYMENTAMOUNT string
	PHONENUMBER   string
	EMAIL         string
	PAYMENTDATE   string
}

// CreatePayment handles POST requests for adding payment entities
// @Summary Create payment
// @Description Create payment
// @ID create-payment
// @Accept   json
// @Produce  json
// @Param payment body ent.Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [post]
func (ctl *PaymentController) CreatePayment(c *gin.Context) {
	obj := Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment binding failed",
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

	pr, err := ctl.client.Promotion.
		Query().
		Where(promotion.IDEQ(int(obj.PROMOTION))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "promotion not found",
		})
		return
	}

	pt, err := ctl.client.Paymenttype.
		Query().
		Where(paymenttype.IDEQ(int(obj.PAYMENTTYPE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype not found",
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

	time, err := time.Parse(time.RFC3339, obj.PAYMENTDATE)
	pm, err := ctl.client.Payment.
		Create().
		SetMember(mem).
		SetPromotion(pr).
		SetPaymenttype(pt).
		SetEmployee(em).
		SetPAYMENTAMOUNT(obj.PAYMENTAMOUNT).
		SetPHONENUMBER(obj.PHONENUMBER).
		SetEMAIL(obj.EMAIL).
		SetPAYMENTDATE(time).
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
		"error":  pm,
	})
}

// GetPayment handles GET requests to retrieve a payment entity
// @Summary Get a payment entity by ID
// @Description get payment by ID
// @ID get-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments/{id} [get]
func (ctl *PaymentController) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pm, err := ctl.client.Payment.
		Query().
		WithMember().
		Where(payment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pm)
}

// GetPaymentbyMember handles GET requests to retrieve a GetPaymentbyMember entity
// @Summary Get a GetPaymentbyMember entity by ID
// @Description get GetPaymentbyMember by ID
// @ID get-GetPaymentbyMember
// @Produce  json
// @Param id path int true "GetPaymentbyMember ID"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentsbymembers/{id} [get]
func (ctl *PaymentController) GetPaymentbyMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Payment.
		Query().
		WithMember().
		WithPromotion().
		WithPaymenttype().
		WithEmployee().
		Where(payment.HasMemberWith(member.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListPayment handles request to get a list of payment entities
// @Summary List payment entities
// @Description list payment entities
// @ID list-payment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [get]
func (ctl *PaymentController) ListPayment(c *gin.Context) {
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

	payments, err := ctl.client.Payment.
		Query().
		WithMember().
		WithPromotion().
		WithPaymenttype().
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

	c.JSON(200, payments)
}

// NewPaymentController creates and registers handles for the payment controller
func NewPaymentController(router gin.IRouter, client *ent.Client) *PaymentController {
	pmc := &PaymentController{
		client: client,
		router: router,
	}
	pmc.register()
	return pmc

}

// InitPaymentController registers routes to the main engine
func (ctl *PaymentController) register() {
	payments := ctl.router.Group("/payments")
	paymentss := ctl.router.Group("/paymentsbymembers")

	payments.GET("", ctl.ListPayment)
	paymentss.GET("id", ctl.GetPaymentbyMember)

	// CRUD
	payments.POST("", ctl.CreatePayment)
	payments.GET(":id", ctl.GetPayment)
}
