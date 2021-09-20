package handler

import (
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
