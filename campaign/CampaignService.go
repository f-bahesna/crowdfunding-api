package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(campaign GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	if campaign, err := s.repository.FindByID(input.ID); err != nil {
		return campaign, err
	} else {
		return campaign, nil
	}
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	//create slug
	slugName := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugName)

	if newCampaign, err := s.repository.Save(campaign); err != nil {
		return newCampaign, err
	} else {
		return newCampaign, nil
	}
}

func (s *service) UpdateCampaign(campaign GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	if campaign, err := s.repository.FindByID(campaign.ID); err != nil {
		return campaign, err
	} else {

		if campaign.UserID != input.User.ID {
			return campaign, errors.New("your not authorized to edit this campaign")
		}

		campaign.Name = input.Name
		campaign.ShortDescription = input.ShortDescription
		campaign.Description = input.Description
		campaign.Perks = input.Perks
		campaign.GoalAmount = input.GoalAmount

		if update, err := s.repository.UpdateCampaign(campaign); err != nil {
			return update, err
		} else {
			return update, nil
		}
	}
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("your not authorized to edit this campaign")
	}

	IsPrimary := false

	if input.IsPrimary {
		IsPrimary = true

		if _, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID); err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = IsPrimary
	campaignImage.FileName = fileLocation

	if campaignImageCreated, err := s.repository.CreateImage(campaignImage); err != nil {
		return campaignImageCreated, err
	} else {
		return campaignImageCreated, nil
	}
}
