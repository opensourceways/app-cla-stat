package app

import "github.com/opensourceways/app-cla-stat/signing/domain/repository"

func NewCorpSigningService(
	repo repository.CorpSigning,
) *corpSigningService {
	return &corpSigningService{
		repo: repo,
	}
}

type CorpSigningService interface {
	List(linkId string) ([]CorpSigningDTO, error)
}

type corpSigningService struct {
	repo repository.CorpSigning
}

func (s *corpSigningService) List(linkId string) ([]CorpSigningDTO, error) {
	v, err := s.repo.FindAll(linkId)
	if err != nil || len(v) == 0 {
		return nil, err
	}

	dtos := make([]CorpSigningDTO, len(v))

	for i := range v {
		item := &v[i]

		dtos[i] = CorpSigningDTO{
			Id:             item.Id,
			Date:           item.Date,
			Language:       item.Link.Language.Language(),
			CorpName:       item.Corp.Name.CorpName(),
			RepName:        item.Rep.Name.Name(),
			RepEmail:       item.Rep.EmailAddr.EmailAddr(),
			HasAdminAdded:  !item.Admin.IsEmpty(),
			HasPDFUploaded: item.HasPDF,
		}
	}

	return dtos, nil
}

type CorpSigningDTO struct {
	Id             string `json:"id"`
	Date           string `json:"date"`
	Language       string `json:"cla_language"`
	CorpName       string `json:"corporation_name"`
	RepName        string `json:"rep_name"`
	RepEmail       string `json:"rep_email"`
	HasAdminAdded  bool   `json:"has_admin_added"`
	HasPDFUploaded bool   `json:"has_pdf_uploaded"`
}
