package repositoryimpl

import "github.com/opensourceways/app-cla-stat/signing/domain"

const (
	fieldOrg      = "org"
	fieldCLAs     = "clas"
	fieldRemoved  = "removed"
	fieldPlatform = "platform"
)

type linkDO struct {
	Id  string    `bson:"id"         json:"id"          required:"true"`
	Org orgInfoDO `bson:"org"        json:"org"         required:"true"`
}

// orgInfoDO
type orgInfoDO struct {
	Org      string `bson:"org"        json:"org"         required:"true"`
	Platform string `bson:"platform"   json:"platform"    required:"true"`
}

func (do *orgInfoDO) toOrgInfo() domain.OrgInfo {
	return domain.OrgInfo{
		Org:      do.Org,
		Platform: do.Platform,
	}
}
