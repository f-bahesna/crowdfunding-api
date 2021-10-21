package transaction

import "time"

//CAMPAIGN TRANSACTION FORMATTER
type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

//format one of list transaction campaign
func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

//format all of list transaction campaign
func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) != 0 {
		var transactionsFormatter []CampaignTransactionFormatter
		for _, transaction := range transactions {
			formatter := FormatCampaignTransaction(transaction)
			transactionsFormatter = append(transactionsFormatter, formatter)
		}

		return transactionsFormatter
	}

	return []CampaignTransactionFormatter{}
}

/**
USER TRANSACTION FORMATTER
**/
type UserTransactionsFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatterUserTransaction(transaction Transaction) UserTransactionsFormatter {
	userTransaction := UserTransactionsFormatter{}
	userTransaction.ID = transaction.ID
	userTransaction.Amount = transaction.Amount
	userTransaction.Status = transaction.Status
	userTransaction.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	userTransaction.Campaign = campaignFormatter

	return userTransaction
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionsFormatter {
	if len(transactions) != 0 {
		var transactionsFormatter []UserTransactionsFormatter

		for _, transaction := range transactions {
			formatter := FormatterUserTransaction(transaction)
			transactionsFormatter = append(transactionsFormatter, formatter)
		}

		return transactionsFormatter
	}

	return []UserTransactionsFormatter{}
}
