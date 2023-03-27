package controllers

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/responder"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CampaignController struct {
	CampaignService services.CampaignService
}

func NewCampaignController(cs services.CampaignService) *CampaignController {
	return &CampaignController{CampaignService: cs}
}

func (cs *CampaignController) Create(c *gin.Context) {
	//validator := NewValidator()
	//
	campaign := models.Campaign{}
	err := c.BindJSON(&campaign)

	//
	//validationErr := validator.Validate.Struct(campaign)
	//errs := validator.TranslateErrors(validationErr)

	if err != nil {
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
	}

	createdCampaign, err := cs.CampaignService.Create(campaign)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": createdCampaign})
	responder.JsonResponse(c.Writer, http.StatusOK, true, createdCampaign)
	//return createdCampaign, nil
}

func (cs *CampaignController) GetAll(c *gin.Context) {
	campaigns, err := cs.CampaignService.GetAll()
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": campaigns})
	responder.JsonResponse(c.Writer, http.StatusOK, true, campaigns)
}

func (cs *CampaignController) GetById(c *gin.Context) {
	//campaign := models.Campaign{}

	id, uuidErr := uuid.Parse(c.Param("id"))
	if uuidErr != nil {
		//c.JSON(400, gin.H{"error": uuidErr.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
	}

	campaign, err := cs.CampaignService.GetById(id)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": campaign})
	responder.JsonResponse(c.Writer, http.StatusOK, true, campaign)
}

func (cs *CampaignController) Update(c *gin.Context) {
	campaign := models.Campaign{}
	id, uuidErr := uuid.Parse(c.Param("id"))
	if uuidErr != nil {
		//c.JSON(400, gin.H{"error": uuidErr.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	err := c.BindJSON(&campaign)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	updatedCampaign, err := cs.CampaignService.Update(campaign, id)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": "success"})
	responder.JsonResponse(c.Writer, http.StatusOK, true, updatedCampaign)
}

func (cs *CampaignController) Delete(c *gin.Context) {
	id, uuidErr := uuid.Parse(c.Param("id"))
	if uuidErr != nil {
		//c.JSON(400, gin.H{"error": uuidErr.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	err := cs.CampaignService.Delete(id)
	if err != nil {
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": "deleted"})
	responder.JsonResponse(c.Writer, http.StatusOK, true, nil)
}
