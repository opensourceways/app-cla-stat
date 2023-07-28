package repositoryimpl

import "github.com/opensourceways/app-cla-stat/signing/domain"

const fieldId = "id"

// managerDO
type managerDO struct {
	Id string `bson:"id" json:"id"`

	RepDO `bson:",inline"`
}

func (do *managerDO) isEmpty() bool {
	return do.Id == ""
}

func (do *managerDO) toManager() (m domain.Manager, err error) {
	if do.isEmpty() {
		return
	}

	if m.Representative, err = do.RepDO.toRep(); err != nil {
		return
	}

	m.Id = do.Id

	return
}
