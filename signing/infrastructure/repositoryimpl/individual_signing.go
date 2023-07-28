package repositoryimpl

import "github.com/opensourceways/app-cla-stat/signing/domain"

func NewIndividualSigning(dao dao) *individualSigning {
	return &individualSigning{
		dao: dao,
	}
}

type individualSigning struct {
	dao dao
}

func (impl *individualSigning) FindAll(linkId string) ([]domain.IndividualSigning, error) {
	filter := linkIdFilter(linkId)
	filter[fieldDeleted] = false

	var dos []individualSigningDO

	if err := impl.dao.GetDocs(filter, nil, &dos); err != nil || len(dos) == 0 {
		return nil, err
	}

	r := make([]domain.IndividualSigning, len(dos))

	for i := range dos {
		dos[i].toIndividualSigning(&r[i])
	}

	return r, nil
}
