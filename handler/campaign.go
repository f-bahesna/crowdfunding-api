package handler

import (
	"golang-practice/campaign"
	"golang-practice/helper"
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
	// handler : mapping id yang di input ke input => di service call formatter
	// service : manggil repo dari yang terima id dari input.
	// repository: get campaign by id
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
