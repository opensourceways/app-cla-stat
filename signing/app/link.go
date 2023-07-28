package app

import (
	commonRepo "github.com/opensourceways/app-cla-stat/common/domain/repository"
	"github.com/opensourceways/app-cla-stat/signing/domain"
	"github.com/opensourceways/app-cla-stat/signing/domain/repository"
	"github.com/opensourceways/app-cla-stat/staterror"
)

func NewLinkService(
	repo repository.Link,
) *linkService {
	return &linkService{
		repo: repo,
	}
}

type CmdToListLink = repository.FindLinksOpt

type LinkService interface {
	List(cmd *CmdToListLink) ([]repository.LinkSummary, error)
	Find(linkId string) (dto domain.OrgInfo, err error)
}

type linkService struct {
	repo repository.Link
}

func (s *linkService) List(cmd *CmdToListLink) ([]repository.LinkSummary, error) {
	return s.repo.FindAll(cmd)
}

func (s *linkService) Find(linkId string) (domain.OrgInfo, error) {
	v, err := s.repo.Find(linkId)
	if err != nil {
		if commonRepo.IsErrorResourceNotFound(err) {
			err = staterror.NewNotFound(staterror.ErrorCodeLinkNotExists)
		}
	}

	return v, err
}
