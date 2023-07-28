package repository

import "github.com/opensourceways/app-cla-stat/signing/domain"

type FindLinksOpt struct {
	Platform string
	Orgs     []string
}

type LinkSummary struct {
	Id  string
	Org domain.OrgInfo
}

type Link interface {
	FindAll(*FindLinksOpt) ([]LinkSummary, error)
	Find(linkId string) (domain.OrgInfo, error)
}
