package app

import "github.com/opensourceways/app-cla-stat/signing/domain/repository"

func NewIndividualSigningService(
	repo repository.IndividualSigning,
) *individualSigningService {
	return &individualSigningService{
		repo: repo,
	}
}

type IndividualSigningService interface {
	List(linkId string) ([]IndividualSigningDTO, error)
}

type individualSigningService struct {
	repo repository.IndividualSigning
}

func (s *individualSigningService) List(linkId string) ([]IndividualSigningDTO, error) {
	v, err := s.repo.FindAll(linkId)
	if err != nil || len(v) == 0 {
		return nil, err
	}

	r := make([]IndividualSigningDTO, len(v))

	for i := range v {
		item := &v[i]

		r[i] = IndividualSigningDTO{
			ID:    item.Id,
			Name:  item.Rep.Name.Name(),
			Date:  item.Date,
			Email: item.Rep.EmailAddr.EmailAddr(),
		}
	}

	return r, nil
}

type IndividualSigningDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Email string `json:"email"`
}
