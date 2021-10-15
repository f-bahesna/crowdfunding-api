package handler

import (
	"golang-practice/helper"
	"golang-practice/transaction"
	"golang-practice/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	//get input from input
	if err := c.ShouldBindUri(&input); err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed get campaign transactions input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := c.MustGet("current_user").(user.User)

	input.User = user

	//panggil service
	if transactions, err := h.service.GetTransactionByCampaignID(input); err != nil {
		response := helper.APIResponse("failed get campaign transaction", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.APIResponse("transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
		c.JSON(http.StatusOK, response)
	}
}
