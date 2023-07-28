package controllers

import (
	"errors"
	"fmt"

	platformAuth "github.com/opensourceways/app-cla-stat/code-platform-auth"
	"github.com/opensourceways/app-cla-stat/code-platform-auth/platforms"
	"github.com/opensourceways/app-cla-stat/models"
)

type AuthController struct {
	baseController
}

func (ctl *AuthController) Prepare() {
	ctl.apiPrepare("")
}

type userAccount struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type accessToken struct {
	Token string `json:"access_token"`
}

// @Title Auth
// @Description authentication by user's password of code platform
// @Param	platform	path 	string				true	"gitee/github"
// @Param	body		body 	controllers.userAccount		true	"body for auth on code platform"
// @Success 201 {object} controllers.accessToken
// @Failure 400 missing_url_path_parameter: missing url path parameter
// @Failure 401 error_parsing_api_body:     parse payload of request failed
// @Failure 402 unsupported_code_platform: unsupported code platform
// @Failure 500 system_error:              system error
// @router /:platform [post]
func (ctl *AuthController) Auth() {
	action := "auth by pw"
	platform := ctl.GetString(":platform")

	var body userAccount

	if fr := ctl.fetchInputPayload(&body); fr != nil {
		ctl.sendFailedResultAsResp(fr, action)
		return
	}

	cp, err := platformAuth.Auth[platformAuth.AuthApplyToLogin].GetAuthInstance(platform)
	if err != nil {
		ctl.sendFailedResponse(400, errUnsupportedCodePlatform, err, action)
		return
	}

	token, err := cp.PasswordCredentialsToken(body.UserName, body.Password)
	if err != nil {
		ctl.sendFailedResponse(401, errWrongIDOrPassword, err, action)
		return
	}

	permission := PermissionOwnerOfOrg
	pl, ec, err := ctl.genACPayload(platform, permission, token)
	if err != nil {
		ctl.sendFailedResponse(500, ec, err, action)
		return
	}

	at, err := ctl.newApiToken(permission, pl)
	if err != nil {
		ctl.sendFailedResponse(500, errSystemError, err, action)
		return
	}

	ctl.sendSuccessResp(accessToken{at})
}

func (ctl *AuthController) genACPayload(platform, permission, platformToken string) (*acForCodePlatformPayload, string, error) {
	pt, err := platforms.NewPlatform(platform)
	if err != nil {
		return nil, errSystemError, err
	}

	orgs, err := pt.ListOrg(platformToken)
	if err != nil {
		return nil, errSystemError, err
	}
	if len(orgs) == 0 {
		return nil, errNoOrg, errors.New("no org")
	}

	user, err := pt.GetUser(platformToken)
	if err != nil {
		return nil, errSystemError, err
	}

	return &acForCodePlatformPayload{
		User:     user,
		Platform: platform,
		Orgs:     orgs,
	}, "", nil
}

type authCodeURL struct {
	URL string `json:"url"`
}

type acForCodePlatformPayload struct {
	User     string   `json:"user"`
	Orgs     []string `json:"orgs"`
	Platform string   `json:"platform"`
}

func (pl *acForCodePlatformPayload) isOwnerOfLink(link string) *failedApiResult {
	v, err := models.FindLink(link)
	if err != nil {
		return parseModelError(err)
	}

	return pl.isOwnerOfOrg(v.Platform, v.OrgID)
}

func (pl *acForCodePlatformPayload) isOwnerOfOrg(platform, org string) *failedApiResult {
	if pl.Platform == platform {
		for _, v := range pl.Orgs {
			if v == org {
				return nil
			}
		}
	}

	return newFailedApiResult(400, errNotYoursOrg, fmt.Errorf("not the org of owner"))
}
