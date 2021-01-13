package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/G16/app/ent/promotiontime"

	"github.com/G16/app/ent"
	"github.com/gin-gonic/gin"
)

// PromotiontimeController defines the struct for the promotiontime controller
type PromotiontimeController struct {
	client *ent.Client
	router gin.IRouter
}

// Promotiontime defines the struct for the promotiontime controller
type Promotiontime struct {
	DATE 	string
	HOUR 	int
	MINUTE	int
}

// CreatePromotiontime handles POST requests for adding promotiontime entities
// @Summary Create promotiontime
// @Description Create promotiontime
// @ID create-promotiontime
// @Accept   json
// @Produce  json
// @Param promotiontime body ent.Promotiontime true "Promotiontime entity"
// @Success 200 {object} ent.Promotiontime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontimes [post]
func (ctl *PromotiontimeController) CreatePromotiontime(c *gin.Context) {
	obj := Promotiontime{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotiontime binding failed",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.DATE)
	ptm, err := ctl.client.Promotiontime.
		Create().
		SetDATE(time).
		SetHOUR(obj.HOUR).
		SetMINUTE(obj.MINUTE).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ptm)
}

// GetPromotiontime handles GET requests to retrieve a promotiontime entity
// @Summary Get a promotiontime entity by ID
// @Description get promotiontime by ID
// @ID get-promotiontime
// @Produce  json
// @Param id path int true "Promotiontime ID"
// @Success 200 {object} Promotiontime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontimes/{id} [get]
func (ctl *PromotiontimeController) GetPromotiontime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	prm, err := ctl.client.Promotiontime.
		Query().
		Where(promotiontime.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prm)
}

// ListPromotiontime handles request to get a list of promotiontime entities
// @Summary List promotiontime entities
// @Description list promotiontime entities
// @ID list-promotiontime
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotiontime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotiontimes [get]
func (ctl *PromotiontimeController) ListPromotiontime(c *gin.Context) {
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

	promotiontimes, err := ctl.client.Promotiontime.
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

	c.JSON(200, promotiontimes)
}

// NewPromotiontimeController creates and registers handles for the promotiontime controller
func NewPromotiontimeController(router gin.IRouter, client *ent.Client) *PromotiontimeController {
	ptmc := &PromotiontimeController{
		client: client,
		router: router,
	}
	ptmc.register()
	return ptmc

}

// InitPromotiontimeController registers routes to the main engine
func (ctl *PromotiontimeController) register() {
	promotiontimes := ctl.router.Group("/promotiontimes")

	promotiontimes.GET("", ctl.ListPromotiontime)

	// CRUD
	promotiontimes.POST("", ctl.CreatePromotiontime)
	promotiontimes.GET(":id", ctl.GetPromotiontime)
}
