package models

var (
	linkAdapterInstance              linkAdapter
	corpSigningAdapterInstance       corpSigningAdapter
	accessTokenAdapterInstance       accessTokenAdapter
	employeeSigningAdapterInstance   employeeSigningAdapter
	individualSigningAdapterInstance individualSigningAdapter
)

type corpSigningAdapter interface {
	List(linkId string) ([]CorporationSigningSummary, error)
}

func RegisterCorpSigningAdapter(a corpSigningAdapter) {
	corpSigningAdapterInstance = a
}

// employeeSigningAdapter
type employeeSigningAdapter interface {
	List(csId string) ([]IndividualSigningBasicInfo, error)
}

func RegisterEmployeeSigningAdapter(a employeeSigningAdapter) {
	employeeSigningAdapterInstance = a
}

// individualSigningAdapter
type individualSigningAdapter interface {
	List(linkId string) ([]IndividualSigningBasicInfo, error)
}

func RegisterIndividualSigningAdapter(a individualSigningAdapter) {
	individualSigningAdapterInstance = a
}

// linkAdapter
type linkAdapter interface {
	List(platform string, orgs []string) ([]LinkInfo, error)
	Find(linkId string) (r LinkInfo, err error)
}

func RegisterLinkAdapter(a linkAdapter) {
	linkAdapterInstance = a
}

// accessTokenAdapter
type accessTokenAdapter interface {
	Add(payload []byte) (string, error)
	Validate(string) ([]byte, error)
}

func RegisterAccessTokenAdapter(at accessTokenAdapter) {
	accessTokenAdapterInstance = at
}
