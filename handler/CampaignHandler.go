package handler

import (
	"fmt"
	"golang-practice/campaign"
	"golang-practice/helper"
	"golang-practice/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//  tangkap parameter di handler
// 	handler => service
// 	service yang menentukan repository yang di call (method mana yang di call)
// 	repo : getAll, getByUserId
// 	db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.FindCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error fetching campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignDetail, err := h.service.FindCampaignByID(input); err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.APIResponse("Campaign Detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
		c.JSON(http.StatusOK, response)
	}
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)

	input.User = currentUser

	if newCampaign, err := h.service.CreateCampaign(input); err != nil {
		response := helper.APIResponse("create campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.APIResponse("create campaign success", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
		c.JSON(http.StatusOK, response)
	}
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var campaignID campaign.GetCampaignDetailInput

	if err := c.ShouldBindUri(&campaignID); err != nil {
		response := helper.APIResponse("update campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.CreateCampaignInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("update campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = c.MustGet("current_user").(user.User)

	if updated, err := h.service.UpdateCampaign(campaignID, input); err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		response := helper.APIResponse("update campaign success", http.StatusOK, "success", campaign.FormatCampaign(updated))
		c.JSON(http.StatusOK, response)
	}
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	//catch form data request
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("upload campaign failed, some error in input", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	//get the file
	if file, err := c.FormFile("file"); err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("upload campaign file failed", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	} else {

		path := fmt.Sprintf("images/campaign/%d-%s", userID, file.Filename)

		//save file
		if err = c.SaveUploadedFile(file, path); err != nil {
			// c.Error(err)
			// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("failed save uploaded file", http.StatusBadRequest, "error", data)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		// TODO: failed to save filepath in db (13/10/2021)
		// Clear: save image to db when is_primary false/true (14/10/2021)
		//save to db
		if _, err = h.service.SaveCampaignImage(input, path); err != nil {
			// c.Error(err)
			// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("failed save campaign image", http.StatusBadRequest, "error", data)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		data := gin.H{"is_uploaded": true}
		response := helper.APIResponse("upload campaign success", http.StatusOK, "success", data)

		c.JSON(http.StatusOK, response)
	}
}
