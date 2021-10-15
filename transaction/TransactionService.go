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
