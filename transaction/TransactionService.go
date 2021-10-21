package transaction

import (
	"errors"
	"golang-practice/campaign"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	if campaign, err := s.campaignRepository.FindByID(input.ID); err != nil {
		return []Transaction{}, err
	} else {
		if campaign.UserID == input.User.ID {
			if transactions, err := s.repository.GetByCampaignID(input.ID); err != nil {
				return transactions, err
			} else {
				return transactions, nil
			}
		}

		return []Transaction{}, errors.New("your not authorized to edit this campaign")
	}
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	if transactions, err := s.repository.GetByUserID(userID); err != nil {
		return transactions, err
	} else {
		return transactions, nil
	}
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	if TransactionRepository, err := s.repository.Save(transaction); err != nil {
		return TransactionRepository, err
	} else {
		return TransactionRepository, nil
	}
}
