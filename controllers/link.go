package controllers

import "github.com/opensourceways/app-cla-stat/models"

type LinkController struct {
	baseController
}

func (ctl *LinkController) Prepare() {
	ctl.apiPrepare(PermissionOwnerOfOrg)
}

// @Title ListLinks
// @Description list all links
// @Success 200 {object} models.LinkInfo
// @Failure 401 missing_token:              token is missing
// @Failure 402 unknown_token:              token is unknown
// @Failure 403 expired_token:              token is expired
// @Failure 404 unauthorized_token:         the permission of token is unmatched
// @Failure 500 system_error:               system error
// @router / [get]
func (ctl *LinkController) ListLinks() {
	action := "list links"

	pl, fr := ctl.tokenPayloadBasedOnCodePlatform()
	if fr != nil {
		ctl.sendFailedResultAsResp(fr, action)
		return
	}

	r, merr := models.ListLink(pl.Platform, pl.Orgs)
	if merr != nil {
		ctl.sendModelErrorAsResp(merr, action)
	} else {
		ctl.sendSuccessResp(r)
	}
}
