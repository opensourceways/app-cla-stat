package controllers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/opensourceways/app-cla-stat/models"
)

type failedApiResult struct {
	reason     error
	errCode    string
	statusCode int
}

func newFailedApiResult(statusCode int, errCode string, err error) *failedApiResult {
	return &failedApiResult{
		statusCode: statusCode,
		errCode:    errCode,
		reason:     err,
	}
}

type baseController struct {
	beego.Controller

	ac *accessController
}

func (ctl *baseController) sendResponse(body interface{}, statusCode int) {
	if statusCode != 0 {
		// if success, don't set status code, otherwise the header set in ctl.ServeJSON
		// will not work. The reason maybe the same as above.
		ctl.Ctx.ResponseWriter.WriteHeader(statusCode)
	}

	ctl.Data["json"] = struct {
		Data interface{} `json:"data"`
	}{
		Data: body,
	}

	ctl.ServeJSON()
}

func (ctl *baseController) sendSuccessResp(body interface{}) {
	ctl.sendResponse(body, 0)
}

func (ctl *baseController) newFuncForSendingFailedResp(action string) func(fr *failedApiResult) {
	return func(fr *failedApiResult) {
		ctl.sendFailedResponse(fr.statusCode, fr.errCode, fr.reason, action)
	}
}

func (ctl *baseController) sendModelErrorAsResp(err error, action string) {
	ctl.sendFailedResultAsResp(parseModelError(err), action)
}

func (ctl *baseController) sendFailedResultAsResp(fr *failedApiResult, action string) {
	ctl.sendFailedResponse(fr.statusCode, fr.errCode, fr.reason, action)
}

func (ctl *baseController) sendFailedResponse(statusCode int, errCode string, reason error, action string) {
	if statusCode >= 500 {
		beego.Error(fmt.Sprintf("Failed to %s, errCode: %s, err: %s", action, errCode, reason.Error()))

		errCode = errSystemError
		reason = fmt.Errorf("system error")
	}

	d := struct {
		ErrCode string `json:"error_code"`
		ErrMsg  string `json:"error_message"`
	}{
		ErrCode: fmt.Sprintf("cla.%s", errCode),
		ErrMsg:  reason.Error(),
	}

	ctl.sendResponse(d, statusCode)
}

func (ctl *baseController) newApiToken(permission string, pl interface{}) (string, error) {
	addr, fr := ctl.getRemoteAddr()
	if fr != nil {
		return "", fr.reason
	}

	ac := &accessController{
		Permission: permission,
		Payload:    pl,
		RemoteAddr: addr,
	}

	v, err := json.Marshal(ac)
	if err != nil {
		return "", err
	}

	return models.NewAccessToken(v)
}

func (ctl *baseController) tokenPayloadBasedOnCodePlatform() (*acForCodePlatformPayload, *failedApiResult) {
	ac, fr := ctl.getAccessController()
	if fr != nil {
		return nil, fr
	}

	if pl, ok := ac.Payload.(*acForCodePlatformPayload); ok {
		return pl, nil
	}
	return nil, newFailedApiResult(500, errSystemError, fmt.Errorf("invalid token payload"))
}

func (ctl *baseController) fetchInputPayload(info interface{}) *failedApiResult {
	return fetchInputPayloadData(ctl.Ctx.Input.RequestBody, info)
}

func (ctl *baseController) checkPathParameter() *failedApiResult {
	rp := ctl.routerPattern()
	if rp == "" {
		return nil
	}

	items := strings.Split(rp, "/")
	for _, item := range items {
		if strings.HasPrefix(item, ":") && ctl.GetString(item) == "" {
			return newFailedApiResult(
				400, errMissingURLPathParameter,
				fmt.Errorf("missing path parameter:%s", item))
		}
	}

	return nil
}

func (ctl *baseController) routerPattern() string {
	if v, ok := ctl.Data["RouterPattern"]; ok {
		return v.(string)
	}
	return ""
}

func (ctl *baseController) apiPrepare(permission string) {
	if permission != "" {
		ctl.apiPrepareWithAC(
			ctl.newAccessController(permission),
			[]string{permission},
		)
	} else {
		ctl.apiPrepareWithAC(nil, nil)
	}
}

func (ctl *baseController) apiPrepareWithAC(ac *accessController, permission []string) {
	if fr := ctl.checkPathParameter(); fr != nil {
		ctl.sendFailedResultAsResp(fr, "")
		ctl.StopRun()
	}

	if ac != nil && len(permission) != 0 {
		if fr := ctl.checkApiReqToken(ac, permission); fr != nil {
			ctl.sendFailedResultAsResp(fr, "")
			ctl.StopRun()
		}

		ctl.Data[apiAccessController] = *ac
	}
}

func (ctl *baseController) newAccessController(permission string) *accessController {
	return &accessController{Payload: &acForCodePlatformPayload{}}
}

func (ctl *baseController) checkApiReqToken(ac *accessController, permission []string) *failedApiResult {
	token := ctl.apiReqHeader(headerToken)
	if token == "" {
		return newFailedApiResult(401, errMissingToken, fmt.Errorf("no token passed"))
	}

	v, err := models.Validate(token)
	if err != nil {
		return parseModelError(err)
	}

	if err := json.Unmarshal(v, ac); err != nil {
		return newFailedApiResult(500, errSystemError, err)
	}

	addr, fr := ctl.getRemoteAddr()
	if fr != nil {
		return fr
	}

	if err := ac.verify(permission, addr); err != nil {
		return newFailedApiResult(403, errUnauthorizedToken, err)
	}

	return nil
}

func (ctl *baseController) getAccessController() (*accessController, *failedApiResult) {
	ac, ok := ctl.Data[apiAccessController]
	if !ok {
		return nil, newFailedApiResult(500, errSystemError, fmt.Errorf("no access controller"))
	}

	if v, ok := ac.(accessController); ok {
		return &v, nil
	}

	return nil, newFailedApiResult(500, errSystemError, fmt.Errorf("can't convert to access controller instance"))
}

func (ctl *baseController) apiReqHeader(h string) string {
	return ctl.Ctx.Input.Header(h)
}

func (ctl *baseController) apiRequestMethod() string {
	return ctl.Ctx.Request.Method
}

func (ctl *baseController) isPostRequest() bool {
	return ctl.apiRequestMethod() == http.MethodPost
}

func (ctl *baseController) getRemoteAddr() (string, *failedApiResult) {
	ips := ctl.Ctx.Request.Header.Get("x-forwarded-for")
	for _, item := range strings.Split(ips, ", ") {
		if net.ParseIP(item) != nil {
			return item, nil
		}
	}

	return "", newFailedApiResult(400, errCanNotFetchClientIP, fmt.Errorf("can not fetch client ip"))
}
