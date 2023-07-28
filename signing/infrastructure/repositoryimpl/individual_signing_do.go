package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/opensourceways/app-cla-stat/signing/domain"
)

// individualSigningDO
type individualSigningDO struct {
	Id   primitive.ObjectID `bson:"_id"  json:"-"`
	Date string             `bson:"date" json:"date"     required:"true"`

	RepDO `bson:",inline"`
}

func (do *individualSigningDO) toIndividualSigning(v *domain.IndividualSigning) (err error) {
	rep, err := do.RepDO.toRep()
	if err != nil {
		return
	}

	*v = domain.IndividualSigning{
		Id:   do.Id.Hex(),
		Rep:  rep,
		Date: do.Date,
	}

	return
}
