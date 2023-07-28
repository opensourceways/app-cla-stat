package adapter

import (
	"github.com/opensourceways/app-cla-stat/models"
	"github.com/opensourceways/app-cla-stat/signing/app"
)

func NewLinkAdapter(
	s app.LinkService,
) *linkAdatper {
	return &linkAdatper{
		s: s,
	}
}

type linkAdatper struct {
	s app.LinkService
}

// List
func (adapter *linkAdatper) List(platform string, orgs []string) ([]models.LinkInfo, error) {
	v, err := adapter.s.List(&app.CmdToListLink{
		Platform: platform,
		Orgs:     orgs,
	})
	if err != nil {
		return nil, err
	}

	r := make([]models.LinkInfo, len(v))
	for i := range v {
		item := &v[i]

		r[i] = models.LinkInfo{
			LinkID:   item.Id,
			OrgID:    item.Org.Org,
			Platform: item.Org.Platform,
		}
	}

	return r, nil
}

// Find
func (adapter *linkAdatper) Find(linkId string) (r models.LinkInfo, err error) {
	v, err := adapter.s.Find(linkId)
	if err != nil {
		return
	}

	r.LinkID = linkId
	r.OrgID = v.Org
	r.Platform = v.Platform

	return
}
