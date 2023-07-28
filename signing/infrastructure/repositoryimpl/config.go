package repositoryimpl

type Config struct {
	Collections Collections `json:"collections" required:"true"`
}

type Collections struct {
	Link              string `json:"link"               required:"true"`
	CorpSigning       string `json:"corp_signing"       required:"true"`
	IndividualSigning string `json:"individual_signing" required:"true"`
}
