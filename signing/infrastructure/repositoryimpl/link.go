package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"

	commonRepo "github.com/opensourceways/app-cla-stat/common/domain/repository"
	"github.com/opensourceways/app-cla-stat/signing/domain"
	"github.com/opensourceways/app-cla-stat/signing/domain/repository"
)

func NewLink(dao dao) *link {
	return &link{
		dao: dao,
	}
}

type link struct {
	dao dao
}

func (impl *link) docFilter(linkId string) bson.M {
	return bson.M{
		fieldId:      linkId,
		fieldDeleted: false,
	}
}

func (impl *link) Find(linkId string) (r domain.OrgInfo, err error) {
	var do linkDO

	err = impl.dao.GetDoc(impl.docFilter(linkId), bson.M{fieldOrg: 1}, &do)
	if err != nil {
		if impl.dao.IsDocNotExists(err) {
			err = commonRepo.NewErrorResourceNotFound(err)
		}
	} else {
		r = do.Org.toOrgInfo()
	}

	return
}

func (impl *link) FindAll(opt *repository.FindLinksOpt) ([]repository.LinkSummary, error) {
	filter := bson.M{
		fieldDeleted:                        false,
		childField(fieldOrg, fieldOrg):      bson.M{mongodbCmdIn: opt.Orgs},
		childField(fieldOrg, fieldPlatform): opt.Platform,
	}

	var dos []linkDO

	project := bson.M{
		fieldCLAs:    0,
		fieldRemoved: 0,
	}

	err := impl.dao.GetDocs(filter, project, &dos)
	if err != nil || len(dos) == 0 {
		return nil, err
	}

	r := make([]repository.LinkSummary, len(dos))
	for i := range dos {
		item := &dos[i]

		r[i] = repository.LinkSummary{
			Id:  item.Id,
			Org: item.Org.toOrgInfo(),
		}
	}

	return r, nil
}
