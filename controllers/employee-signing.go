package controllers

import "github.com/opensourceways/app-cla-stat/models"

type EmployeeSigningController struct {
	baseController
}

func (ctl *EmployeeSigningController) Prepare() {
	ctl.apiPrepare(PermissionOwnerOfOrg)
}

// @Title List
// @Description get all the employees by community manager
// @Param	:link_id	path 	string		true		"link id"
// @Param	:signing_id	path 	string		true		"signing id"
// @Success 200 {object} models.IndividualSigningBasicInfo
// @Failure 400 missing_url_path_parameter: missing url path parameter
// @Failure 401 missing_token:              token is missing
// @Failure 402 unknown_token:              token is unknown
// @Failure 403 expired_token:              token is expired
// @Failure 404 unauthorized_token:         the permission of token is unmatched
// @Failure 405 unknown_link:               unkown link id
// @Failure 406 not_yours_org:              the link doesn't belong to your community
// @Failure 500 system_error:               system error
// @router /:link_id/:signing_id [get]
func (ctl *EmployeeSigningController) List() {
	action := "list employees"
	linkId := ctl.GetString(":link_id")
	csId := ctl.GetString(":signing_id")

	pl, fr := ctl.tokenPayloadBasedOnCodePlatform()
	if fr != nil {
		ctl.sendFailedResultAsResp(fr, action)
		return
	}

	if fr := pl.isOwnerOfLink(linkId); fr != nil {
		ctl.sendFailedResultAsResp(fr, action)
		return
	}

	r, merr := models.ListEmployeeSignings(csId)
	if merr != nil {
		ctl.sendModelErrorAsResp(merr, action)
	} else {
		ctl.sendSuccessResp(r)
	}
}
