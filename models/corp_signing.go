package models

type CorporationSigningSummary struct {
	CorporationSigningBasicInfo

	Id          string `json:"id"`
	AdminAdded  bool   `json:"admin_added"`
	PDFUploaded bool   `json:"pdf_uploaded"`
}

type CorporationSigningBasicInfo struct {
	CLALanguage     string `json:"cla_language"`
	AdminEmail      string `json:"admin_email"`
	AdminName       string `json:"admin_name"`
	CorporationName string `json:"corporation_name"`
	Date            string `json:"date"`
}
