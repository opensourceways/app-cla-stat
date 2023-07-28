package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/opensourceways/app-cla-stat/signing/domain/repository"
)

func NewCorpSigning(dao dao) *corpSigning {
	return &corpSigning{
		dao: dao,
	}
}

type corpSigning struct {
	dao dao
}

func (impl *corpSigning) toCorpSigningIndex(corpSigningId string) (bson.M, error) {
	return impl.dao.DocIdFilter(corpSigningId)
}

func (impl *corpSigning) FindAll(linkId string) ([]repository.CorpSigningSummary, error) {
	filter := linkIdFilter(linkId)

	project := bson.M{
		fieldDate:   1,
		fieldLang:   1,
		fieldRep:    1,
		fieldCorp:   1,
		fieldAdmin:  1,
		fieldLinkId: 1,
		fieldHasPDF: 1,
	}

	var dos []corpSigningDO

	if err := impl.dao.GetDocs(filter, project, &dos); err != nil {
		return nil, err
	}

	v := make([]repository.CorpSigningSummary, len(dos))
	for i := range dos {
		if err := dos[i].toCorpSigningSummary(&v[i]); err != nil {
			return nil, err
		}
	}

	return v, nil
}
