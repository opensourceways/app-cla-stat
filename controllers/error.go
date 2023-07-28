package controllers

import "github.com/opensourceways/app-cla-stat/staterror"

const (
	errNoOrg                   = "no_org"
	errNotYoursOrg             = "not_yours_org"
	errSystemError             = "system_error"
	errMissingToken            = "missing_token"
	errUnknownToken            = "unknown_token"
	errParsingApiBody          = "error_parsing_api_body"
	errUnauthorizedToken       = "unauthorized_token"
	errWrongIDOrPassword       = "wrong_id_or_pw"
	errCanNotFetchClientIP     = "can_not_fetch_client_ip"
	errMissingURLPathParameter = "missing_url_path_parameter"
	errUnsupportedCodePlatform = "unsupported_code_platform"
)

type errorCode interface {
	ErrorCode() string
}

type errorNotFound interface {
	NotFound()
}

func parseModelError(err error) *failedApiResult {
	if err == nil {
		return nil
	}

	sc := 500
	code := errSystemError

	if v, ok := err.(errorCode); ok {
		code = v.ErrorCode()
		sc = 400

		if code == staterror.ErrorCodeAccessTokenInvalid {
			sc = 401
		}

		if code == staterror.ErrorCodeAccessTokenExpiry {
			sc = 401
		}
	}

	if _, ok := err.(errorNotFound); ok {
		sc = 404
	}

	return newFailedApiResult(sc, code, err)
}
