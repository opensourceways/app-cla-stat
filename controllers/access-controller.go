package controllers

import "errors"

const PermissionOwnerOfOrg = "owner of org"

type accessController struct {
	RemoteAddr string      `json:"remote_addr"`
	Permission string      `json:"permission"`
	Payload    interface{} `json:"payload"`
}

func (ctl *accessController) verify(permission []string, addr string) error {
	if ctl.RemoteAddr != addr {
		return errors.New("unmatched remote address")
	}

	for _, p := range permission {
		if p == ctl.Permission {
			return nil
		}
	}

	return errors.New("not allowed permission")
}
