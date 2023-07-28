package app

import "github.com/opensourceways/app-cla-stat/signing/domain/repository"

func NewEmployeeSigningService(
	repo repository.CorpSigning,
) *employeeSigningService {
	return &employeeSigningService{
		repo: repo,
	}
}

type EmployeeSigningService interface {
	List(csId string) ([]IndividualSigningDTO, error)
}

type employeeSigningService struct {
	repo repository.CorpSigning
}

// List
func (s *employeeSigningService) List(csId string) ([]IndividualSigningDTO, error) {
	v, err := s.repo.FindEmployees(csId)
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
