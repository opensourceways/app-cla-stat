package domain

import "github.com/opensourceways/app-cla-stat/signing/domain/dp"

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
)

type Representative struct {
	Name      dp.Name
	EmailAddr dp.EmailAddr
}

type CLAInfo struct {
	CLAId    string
	Language dp.Language
}

type LinkInfo struct {
	Id string

	CLAInfo
}

type Manager struct {
	Id string
	Representative
}

func (m *Manager) IsEmpty() bool {
	return m.Id == ""
}

type Corporation struct {
	Name               dp.CorpName
	AllEmailDomains    []string
	PrimaryEmailDomain string
}

type EmployeeSigning = IndividualSigning
