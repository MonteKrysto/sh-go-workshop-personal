package controllers

import (
	"fmt"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/responder"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ProfileController struct {
	ProfileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *ProfileController {
	return &ProfileController{
		ProfileService: profileService,
	}
}

func (pc ProfileController) Create(c *gin.Context) {
	profile := models.Profile{}

	err := c.BindJSON(&profile)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	newProfile, err := pc.ProfileService.Create(profile)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": newProfile})
	responder.JsonResponse(c.Writer, http.StatusOK, true, newProfile)
}

func (pc ProfileController) GetAll(c *gin.Context) {
	profiles, err := pc.ProfileService.GetAll()
	if err != nil {
		responder.JsonResponse(c.Writer, http.StatusBadGateway, false, nil)
		return
	}

	responder.JsonResponse(c.Writer, http.StatusOK, true, profiles)
}

func (pc ProfileController) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	profile, err := pc.ProfileService.GetById(id)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": profile})
	responder.JsonResponse(c.Writer, http.StatusOK, true, profile)
}

func (pc ProfileController) GetByUserId(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	profile, err := pc.ProfileService.GetByUserId(id)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": profile})
	responder.JsonResponse(c.Writer, http.StatusOK, true, profile)
}

func (pc ProfileController) Update(c *gin.Context) {
	profile := models.Profile{}
	id, _ := uuid.Parse(c.Param("id"))

	err := c.BindJSON(&profile)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	updatedProfile, err := pc.ProfileService.Update(profile, id)
	fmt.Println("updatedProfile", updatedProfile)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": "record updated"})
	responder.JsonResponse(c.Writer, http.StatusOK, true, updatedProfile)
}

func (pc ProfileController) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))

	err := pc.ProfileService.Delete(id)
	if err != nil {
		//c.JSON(400, gin.H{"error": err.Error()})
		responder.JsonResponse(c.Writer, http.StatusBadRequest, false, nil)
		return
	}

	//c.JSON(200, gin.H{"data": "Profile deleted"})
	responder.JsonResponse(c.Writer, http.StatusOK, true, "Profile deleted")
}
