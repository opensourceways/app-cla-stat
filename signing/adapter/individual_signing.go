package adapter

import (
	"github.com/opensourceways/app-cla-stat/models"
	"github.com/opensourceways/app-cla-stat/signing/app"
)

func NewIndividualSigningAdapter(s app.IndividualSigningService) *individualSigningAdatper {
	return &individualSigningAdatper{s}
}

type individualSigningAdatper struct {
	s app.IndividualSigningService
}

func (adapter *individualSigningAdatper) List(linkId string) (
	[]models.IndividualSigningBasicInfo, error,
) {
	v, err := adapter.s.List(linkId)
	if err != nil {
		return nil, err
	}

	r := make([]models.IndividualSigningBasicInfo, len(v))

	for i := range v {
		item := &v[i]

		r[i] = models.IndividualSigningBasicInfo{
			ID:    item.ID,
			Name:  item.Name,
			Email: item.Email,
			Date:  item.Date,
		}
	}

	return r, nil
}
