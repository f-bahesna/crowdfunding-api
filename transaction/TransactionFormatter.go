package transaction

import "time"

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
