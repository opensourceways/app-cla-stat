package models

// link

func ListLink(platform string, orgs []string) ([]LinkInfo, error) {
	return linkAdapterInstance.List(platform, orgs)
}

func FindLink(linkId string) (LinkInfo, error) {
	return linkAdapterInstance.Find(linkId)
}

// corp signing

func ListCorpSigning(linkID string) ([]CorporationSigningSummary, error) {
	return corpSigningAdapterInstance.List(linkID)
}

// employee signing

func ListEmployeeSignings(csId string) ([]IndividualSigningBasicInfo, error) {
	return employeeSigningAdapterInstance.List(csId)
}

// individual signing

func ListIndividualSignings(linkId string) ([]IndividualSigningBasicInfo, error) {
	return individualSigningAdapterInstance.List(linkId)
}

// access token
func NewAccessToken(payload []byte) (string, error) {
	return accessTokenAdapterInstance.Add(payload)
}

func Validate(token string) ([]byte, error) {
	return accessTokenAdapterInstance.Validate(token)
}
