package repositoryimpl

import "github.com/opensourceways/app-cla-stat/signing/domain"

const fieldEnabled = "enabled"

// employeeSigningDO
type employeeSigningDO struct {
	Id      string `bson:"id"       json:"id"       required:"true"`
	Date    string `bson:"date"     json:"date"     required:"true"`
	Enabled bool   `bson:"enabled"  json:"enabled"`

	RepDO `bson:",inline"`
}

func (do *employeeSigningDO) toEmployeeSigning(es *domain.EmployeeSigning) (err error) {
	rep, err := do.RepDO.toRep()
	if err != nil {
		return
	}

	*es = domain.EmployeeSigning{
		Id:   do.Id,
		Rep:  rep,
		Date: do.Date,
	}

	return
}
