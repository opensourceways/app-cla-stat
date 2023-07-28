package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/opensourceways/app-cla-stat/signing/domain"
)

func (impl *corpSigning) FindEmployees(csId string) ([]domain.EmployeeSigning, error) {
	filter, err := impl.toCorpSigningIndex(csId)
	if err != nil {
		return nil, err
	}

	var do corpSigningDO

	if err = impl.dao.GetDoc(filter, bson.M{fieldEmployees: 1}, &do); err != nil {
		if impl.dao.IsDocNotExists(err) {
			err = nil
		}

		return nil, err
	}

	return do.toEmployeeSignings()
}
