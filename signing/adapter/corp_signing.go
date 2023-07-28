package adapter

import (
	"github.com/opensourceways/app-cla-stat/models"
	"github.com/opensourceways/app-cla-stat/signing/app"
)

func NewCorpSigningAdapter(
	s app.CorpSigningService,
) *corpSigningAdatper {
	return &corpSigningAdatper{
		s: s,
	}
}

type corpSigningAdatper struct {
	s app.CorpSigningService
}

// List
func (adapter *corpSigningAdatper) List(linkId string) (
	[]models.CorporationSigningSummary, error,
) {
	v, err := adapter.s.List(linkId)
	if err != nil {
		return nil, err
	}

	r := make([]models.CorporationSigningSummary, len(v))
	for i := range v {
		item := &v[i]

		r[i] = models.CorporationSigningSummary{
			CorporationSigningBasicInfo: models.CorporationSigningBasicInfo{
				Date:            item.Date,
				AdminName:       item.RepName,
				AdminEmail:      item.RepEmail,
				CLALanguage:     item.Language,
				CorporationName: item.CorpName,
			},
			Id:          item.Id,
			AdminAdded:  item.HasAdminAdded,
			PDFUploaded: item.HasPDFUploaded,
		}
	}

	return r, nil
}
