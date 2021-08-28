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

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
