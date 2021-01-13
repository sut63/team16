package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent/member"

	"github.com/G16/app/ent"
	"github.com/gin-gonic/gin"
)

// MemberController defines the struct for the member controller
type MemberController struct {
	client *ent.Client
	router gin.IRouter
}

// Member defines the struct for the member controller
type Member struct {
	MEMBERID 	string
	NAME		string
}

// CreateMember handles POST requests for adding member entities
// @Summary Create member
// @Description Create member
// @ID create-member
// @Accept   json
// @Produce  json
// @Param member body ent.Member true "Member entity"
// @Success 200 {object} ent.Member
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members [post]
func (ctl *MemberController) CreateMember(c *gin.Context) {
	obj := Member{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "member binding failed",
		})
		return
	}

	mb, err := ctl.client.Member.
		Create().
		SetMEMBERID(obj.MEMBERID).
		SetNAME(obj.NAME).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, mb)
}

// GetMember handles GET requests to retrieve a member entity
// @Summary Get a member entity by ID
// @Description get member by ID
// @ID get-member
// @Produce  json
// @Param id path int true "member ID"
// @Success 200 {object} member
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members/{id} [get]
func (ctl *MemberController) GetMember(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	mb, err := ctl.client.Member.
		Query().
		Where(member.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, mb)
}

// ListMember handles request to get a list of member entities
// @Summary List member entities
// @Description list member entities
// @ID list-member
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Member
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /members [get]
func (ctl *MemberController) ListMember(c *gin.Context) {
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

	members, err := ctl.client.Member.
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

	c.JSON(200, members)
}

// NewMemberController creates and registers handles for the member controller
func NewMemberController(router gin.IRouter, client *ent.Client) *MemberController {
	mbc := &MemberController{
		client: client,
		router: router,
	}
	mbc.register()
	return mbc

}

// InitMemberController registers routes to the main engine
func (ctl *MemberController) register() {
	members := ctl.router.Group("/members")

	members.GET("", ctl.ListMember)

	// CRUD
	members.POST("", ctl.CreateMember)
	members.GET(":id", ctl.GetMember)
}
