package controllers

import (
	"context"
	"strconv"
	

	"github.com/G16/app/ent/promotiontype"

	"github.com/G16/app/ent"
	"github.com/gin-gonic/gin"
)

// PromotiontypeController defines the struct for the promotiontype controller
type PromotiontypeController struct {
	client *ent.Client
	router gin.IRouter
}

// Promotiontype defines the struct for the promotiontype controller
type Promotiontype struct {
	TYPE string
}

// CreatePromotiontype handles POST requests for adding promotiontype entities
// @Summary Create promotiontype
// @Description Create promotiontype
// @ID create-promotiontype
// @Accept   json
// @Produce  json
// @Param promotiontype body ent.Promotiontype true "Promotiontype entity"
// @Success 200 {object} ent.Promotiontype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontypes [post]
func (ctl *PromotiontypeController) CreatePromotiontype(c *gin.Context) {
	obj := Promotiontype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotiontype binding failed",
		})
		return
	}

	prt, err := ctl.client.Promotiontype.
		Create().
		SetTYPE(obj.TYPE).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, prt)
}

// GetPromotiontype handles GET requests to retrieve a promotiontype entity
// @Summary Get a promotiontype entity by ID
// @Description get promotiontype by ID
// @ID get-promotiontype
// @Produce  json
// @Param id path int true "Promotiontype ID"
// @Success 200 {object} Promotiontype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontypes/{id} [get]
func (ctl *PromotiontypeController) GetPromotiontype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	prt, err := ctl.client.Promotiontype.
		Query().
		Where(promotiontype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prt)
}

// ListPromotiontype handles request to get a list of promotiontype entities
// @Summary List promotiontype entities
// @Description list promotiontype entities
// @ID list-promotiontype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotiontype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontypes [get]
func (ctl *PromotiontypeController) ListPromotiontype(c *gin.Context) {
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

	promotiontypes, err := ctl.client.Promotiontype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, promotiontypes)
}

// NewPromotiontypeController creates and registers handles for the promotiontype controller
func NewPromotiontypeController(router gin.IRouter, client *ent.Client) *PromotiontypeController {
	prtc := &PromotiontypeController{
		client: client,
		router: router,
	}
	prtc.register()
	return prtc

}

// InitPromotiontypeController registers routes to the main engine
func (ctl *PromotiontypeController) register() {
	promotiontypes := ctl.router.Group("/promotiontypes")

	promotiontypes.GET("", ctl.ListPromotiontype)

	// CRUD
	promotiontypes.POST("", ctl.CreatePromotiontype)
	promotiontypes.GET(":id", ctl.GetPromotiontype)
}
