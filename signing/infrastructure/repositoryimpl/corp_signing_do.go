package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/opensourceways/app-cla-stat/signing/domain"
	"github.com/opensourceways/app-cla-stat/signing/domain/dp"
	"github.com/opensourceways/app-cla-stat/signing/domain/repository"
)

const (
	fieldPDF       = "pdf"
	fieldRep       = "rep"
	fieldDate      = "date"
	fieldCorp      = "corp"
	fieldName      = "name"
	fieldLang      = "lang"
	fieldAdmin     = "admin"
	fieldEmail     = "email"
	fieldHasPDF    = "has_pdf"
	fieldLinkId    = "link_id"
	fieldDomain    = "domain"
	fieldDomains   = "domains"
	fieldDeleted   = "deleted"
	fieldVersion   = "version"
	fieldManagers  = "managers"
	fieldEmployees = "employees"
)

// corpSigningDO
type corpSigningDO struct {
	Id       primitive.ObjectID `bson:"_id"      json:"-"`
	Date     string             `bson:"date"     json:"date"     required:"true"`
	LinkId   string             `bson:"link_id"  json:"link_id"  required:"true"`
	Language string             `bson:"lang"     json:"lang"     required:"true"`
	Rep      RepDO              `bson:"rep"      json:"rep"      required:"true"`
	Corp     corpDO             `bson:"corp"     json:"corp"     required:"true"`

	HasPDF    bool                `bson:"has_pdf"       json:"has_pdf"`
	Admin     managerDO           `bson:"admin"         json:"admin"`
	Employees []employeeSigningDO `bson:"employees"     json:"employees"`
}

func (do *corpSigningDO) index() string {
	return do.Id.Hex()
}

func (do *corpSigningDO) toCorpSigningSummary(cs *repository.CorpSigningSummary) (err error) {
	rep, err := do.Rep.toRep()
	if err != nil {
		return
	}

	corp, err := do.Corp.toCorp()
	if err != nil {
		return
	}

	admin, err := do.Admin.toManager()
	if err != nil {
		return
	}

	*cs = repository.CorpSigningSummary{
		Id:     do.index(),
		Date:   do.Date,
		HasPDF: do.HasPDF,
		Rep:    rep,
		Corp:   corp,
		Admin:  admin,
	}

	cs.Link.Id = do.LinkId
	cs.Link.Language, err = dp.NewLanguage(do.Language)

	return
}

func (do *corpSigningDO) toEmployeeSignings() (es []domain.EmployeeSigning, err error) {
	es = make([]domain.EmployeeSigning, len(do.Employees))

	for i := range do.Employees {
		if err = do.Employees[i].toEmployeeSigning(&es[i]); err != nil {
			return
		}
	}

	return
}

// representative DO
type RepDO struct {
	Name  string `bson:"name"  json:"name"  required:"true"`
	Email string `bson:"email" json:"email" required:"true"`
}

func (do *RepDO) toRep() (rep domain.Representative, err error) {
	if rep.Name, err = dp.NewName(do.Name); err != nil {
		return
	}

	rep.EmailAddr, err = dp.NewEmailAddr(do.Email)

	return
}

// corporation DO
type corpDO struct {
	Name    string   `bson:"name"     json:"name"      required:"true"`
	Domain  string   `bson:"domain"   json:"domain"    required:"true"`
	Domains []string `bson:"domains"  json:"domains"   required:"true"`
}

func (do *corpDO) toCorp() (c domain.Corporation, err error) {
	if c.Name, err = dp.NewCorpName(do.Name); err != nil {
		return
	}

	c.PrimaryEmailDomain = do.Domain
	c.AllEmailDomains = do.Domains

	return
}
