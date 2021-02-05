package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent/zone"

	"github.com/G16/app/ent"
	"github.com/gin-gonic/gin"
)

// ZoneController defines the struct for the zone controller
type ZoneController struct {
	client *ent.Client
	router gin.IRouter
}

// Zone defines the struct for the zone controller
type Zone struct {
	EQUIPMENTZONE string
}

// CreateZone handles POST requests for adding zone entities
// @Summary Create zone
// @Description Create zone
// @ID create-zone
// @Accept   json
// @Produce  json
// @Param zone body ent.Zone true "Zone entity"
// @Success 200 {object} ent.Zone
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zones [post]
func (ctl *ZoneController) CreateZone(c *gin.Context) {
	obj := Zone{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "zone binding failed",
		})
		return
	}

	zn, err := ctl.client.Zone.
		Create().
		SetEQUIPMENTZONE(obj.EQUIPMENTZONE).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, zn)
}

// GetZone handles GET requests to retrieve a zone entity
// @Summary Get a zone entity by ID
// @Description get zone by ID
// @ID get-zone
// @Produce  json
// @Param id path int true "Zone ID"
// @Success 200 {object} ent.Zone
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zones/{id} [get]
func (ctl *ZoneController) GetZone(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	zn, err := ctl.client.Zone.
		Query().
		Where(zone.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, zn)
}

// ListZone handles request to get a list of zone entities
// @Summary List zone entities
// @Description list zone entities
// @ID list-zone
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Zone
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zones [get]
func (ctl *ZoneController) ListZone(c *gin.Context) {
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

	zones, err := ctl.client.Zone.
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

	c.JSON(200, zones)
}

// NewZoneController creates and registers handles for the zone controller
func NewZoneController(router gin.IRouter, client *ent.Client) *ZoneController {
	znc := &ZoneController{
		client: client,
		router: router,
	}
	znc.register()
	return znc

}

// InitZoneController registers routes to the main engine
func (ctl *ZoneController) register() {
	zones := ctl.router.Group("/zones")

	zones.GET("", ctl.ListZone)

	// CRUD
	zones.POST("", ctl.CreateZone)
	zones.GET(":id", ctl.GetZone)
}
