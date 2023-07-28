package repository

import "github.com/opensourceways/app-cla-stat/signing/domain"

type CorpSigningSummary struct {
	Id     string
	Date   string
	HasPDF bool
	Link   domain.LinkInfo
	Rep    domain.Representative
	Corp   domain.Corporation
	Admin  domain.Manager
}

type CorpSigning interface {
	FindAll(linkId string) ([]CorpSigningSummary, error)
	FindEmployees(string) ([]domain.EmployeeSigning, error)
}
