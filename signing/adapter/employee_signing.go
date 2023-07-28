package adapter

import (
	"github.com/opensourceways/app-cla-stat/models"
	"github.com/opensourceways/app-cla-stat/signing/app"
)

func NewEmployeeSigningAdapter(s app.EmployeeSigningService) *employeeSigningAdatper {
	return &employeeSigningAdatper{s}
}

type employeeSigningAdatper struct {
	s app.EmployeeSigningService
}

// List
func (adapter *employeeSigningAdatper) List(csId string) (
	[]models.IndividualSigningBasicInfo, error,
) {
	v, err := adapter.s.List(csId)
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
