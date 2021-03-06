package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/course"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/promotionamount"
	"github.com/G16/app/ent/promotiontype"
	"github.com/gin-gonic/gin"
)

// PromotionController defines the struct for the promotion controller
type PromotionController struct {
	client *ent.Client
	router gin.IRouter
}

// Promotion defines the struct for the promotion controller
type Promotion struct {
	PROMOTIONTYPE   int
	PROMOTIONAMOUNT int
	EMPLOYEE        int
	COURSE			int
	NAME            string
	DESC            string
	CODE            string
	DATE            string
}

// CreatePromotion handles POST requests for adding promotion entities
// @Summary Create promotion
// @Description Create promotion
// @ID create-promotion
// @Accept   json
// @Produce  json
// @Param promotion body ent.Promotion true "Promotion entity"
// @Success 200 {object} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [post]
func (ctl *PromotionController) CreatePromotion(c *gin.Context) {
	obj := Promotion{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotion binding failed",
		})
		return
	}

	prt, err := ctl.client.Promotiontype.
		Query().
		Where(promotiontype.IDEQ(int(obj.PROMOTIONTYPE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "promotiontype not found",
		})
		return
	}

	pa, err := ctl.client.Promotionamount.
		Query().
		Where(promotionamount.IDEQ(int(obj.PROMOTIONAMOUNT))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "promotionamount not found",
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
			"error": "course not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.DATE)
	pr, err := ctl.client.Promotion.
		Create().
		SetPromotiontype(prt).
		SetPromotionamount(pa).
		SetEmployee(em).
		SetCourse(co).
		SetNAME(obj.NAME).
		SetDESC(obj.DESC).
		SetCODE(obj.CODE).
		SetDATE(time).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"error": pr,
	})
}

// GetPromotion handles GET requests to retrieve a promotion entity
// @Summary Get a promotion entity by ID
// @Description get promotion by ID
// @ID get-promotion
// @Produce  json
// @Param id path int true "Promotion ID"
// @Success 200 {object} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions/{id} [get]
func (ctl *PromotionController) GetPromotion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pr, err := ctl.client.Promotion.
		Query().
		WithPromotiontype().
		WithPromotionamount().
		WithEmployee().
		WithCourse().
		Where(promotion.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListPromotion handles request to get a list of promotion entities
// @Summary List promotion entities
// @Description list promotion entities
// @ID list-promotion
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [get]
func (ctl *PromotionController) ListPromotion(c *gin.Context) {
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

	promotions, err := ctl.client.Promotion.
		Query().
		WithPromotiontype().
		WithPromotionamount().
		WithEmployee().
		WithCourse().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, promotions)
}

// NewPromotionController creates and registers handles for the promotion controller
func NewPromotionController(router gin.IRouter, client *ent.Client) *PromotionController {
	prc := &PromotionController{
		client: client,
		router: router,
	}
	prc.register()
	return prc

}

// InitPromotionController registers routes to the main engine
func (ctl *PromotionController) register() {
	promotions := ctl.router.Group("/promotions")

	promotions.GET("", ctl.ListPromotion)

	// CRUD
	promotions.POST("", ctl.CreatePromotion)
	promotions.GET(":id", ctl.GetPromotion)
}
