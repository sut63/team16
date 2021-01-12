package controllers

import (
	"context"
	"strconv"

	"github.com/G16/app/ent"
	"github.com/G16/app/ent/classifier"
	"github.com/gin-gonic/gin"
)

// ClassifierController defines the struct for the classifier controller
type ClassifierController struct {
	client *ent.Client
	router gin.IRouter
}

// Classifier defines the struct for the classifier controller
type Classifier struct {
	EQUIPMENTCLASSIFIER string
}

// CreateClassifier handles POST requests for adding classifier entities
// @Summary Create classifier
// @Description Create classifier
// @ID create-classifier
// @Accept   json
// @Produce  json
// @Param classifier body ent.Classifier true "Classifier entity"
// @Success 200 {object} ent.Classifier
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classifiers [post]
func (ctl *ClassifierController) CreateClassifier(c *gin.Context) {
	obj := ent.Classifier{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "classifier binding failed",
		})
		return
	}

	cf, err := ctl.client.Classifier.
		Create().
		SetEQUIPMENTCLASSIFIER(obj.EQUIPMENTCLASSIFIER).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cf)
}

// GetClassifier handles GET requests to retrieve a classifier entity
// @Summary Get a classifier entity by ID
// @Description get classifier by ID
// @ID get-classifier
// @Produce  json
// @Param id path int true "Classifier ID"
// @Success 200 {object} ent.Classifier
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classifiers/{id} [get]
func (ctl *ClassifierController) GetClassifier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cf, err := ctl.client.Classifier.
		Query().
		Where(classifier.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cf)
}

// ListClassifier handles request to get a list of classifier entities
// @Summary List classifier entities
// @Description list classifier entities
// @ID list-classifier
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Classifier
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /classifiers [get]
func (ctl *ClassifierController) ListClassifier(c *gin.Context) {
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

	classifiers, err := ctl.client.Classifier.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, classifiers)
}

// NewClassifierController creates and registers handles for the classifier controller
func NewClassifierController(router gin.IRouter, client *ent.Client) *ClassifierController {
	cfc := &ClassifierController{
		client: client,
		router: router,
	}
	cfc.register()
	return cfc
}

// InitClassifierController registers routes to the main engine
func (ctl *ClassifierController) register() {
	classifiers := ctl.router.Group("/classifiers")

	classifiers.GET("", ctl.ListClassifier)

	// CRUD
	classifiers.POST("", ctl.CreateClassifier)
	classifiers.GET(":id", ctl.GetClassifier)
}
