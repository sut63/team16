package controllers

import (
	"context"
	"strconv"
	

	"github.com/G16/app/ent/promotionamount"

	"github.com/G16/app/ent"
	"github.com/gin-gonic/gin"
)

// PromotionamountController defines the struct for the promotionamount controller
type PromotionamountController struct {
	client *ent.Client
	router gin.IRouter
}

// Promotionamount defines the struct for the promotionamount controller
type Promotionamount struct {
	AMOUNT int
}

// CreatePromotionamount handles POST requests for adding promotionamount entities
// @Summary Create promotionamount
// @Description Create promotionamount
// @ID create-promotionamount
// @Accept   json
// @Produce  json
// @Param promotionamount body ent.Promotionamount true "Promotionamount entity"
// @Success 200 {object} ent.Promotionamount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotionamounts [post]
func (ctl *PromotionamountController) CreatePromotionamount(c *gin.Context) {
	obj := Promotionamount{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotionamount binding failed",
		})
		return
	}

	prm, err := ctl.client.Promotionamount.
		Create().
		SetAMOUNT(obj.AMOUNT).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, prm)
}

// GetPromotionamount handles GET requests to retrieve a promotionamount entity
// @Summary Get a promotionamount entity by ID
// @Description get promotionamount by ID
// @ID get-promotionamount
// @Produce  json
// @Param id path int true "Promotionamount ID"
// @Success 200 {object} Promotionamount
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotionamounts/{id} [get]
func (ctl *PromotionamountController) GetPromotionamount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	prm, err := ctl.client.Promotionamount.
		Query().
		Where(promotionamount.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prm)
}

// ListPromotionamount handles request to get a list of promotionamount entities
// @Summary List promotionamount entities
// @Description list promotionamount entities
// @ID list-promotionamount
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotionamount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotionamounts [get]
func (ctl *PromotionamountController) ListPromotionamount(c *gin.Context) {
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

	promotionamounts, err := ctl.client.Promotionamount.
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

	c.JSON(200, promotionamounts)
}

// NewPromotionamountController creates and registers handles for the promotionamount controller
func NewPromotionamountController(router gin.IRouter, client *ent.Client) *PromotionamountController {
	prmc := &PromotionamountController{
		client: client,
		router: router,
	}
	prmc.register()
	return prmc

}

// InitPromotionamountController registers routes to the main engine
func (ctl *PromotionamountController) register() {
	promotionamounts := ctl.router.Group("/promotionamounts")

	promotionamounts.GET("", ctl.ListPromotionamount)

	// CRUD
	promotionamounts.POST("", ctl.CreatePromotionamount)
	promotionamounts.GET(":id", ctl.GetPromotionamount)
}
