package main

import (
	accesstokenservice "github.com/opensourceways/app-cla-stat/accesstoken/domain/service"
	"github.com/opensourceways/app-cla-stat/accesstoken/infrastructure/symmetricencryptionimpl"
	"github.com/opensourceways/app-cla-stat/common/infrastructure/mongodb"
	"github.com/opensourceways/app-cla-stat/config"
	"github.com/opensourceways/app-cla-stat/models"
	"github.com/opensourceways/app-cla-stat/signing/adapter"
	"github.com/opensourceways/app-cla-stat/signing/app"
	"github.com/opensourceways/app-cla-stat/signing/infrastructure/repositoryimpl"
)

func initSigning(cfg *config.Config) error {
	symmetric, err := symmetricencryptionimpl.NewSymmetricEncryptionImpl(&cfg.Symmetric)
	if err != nil {
		return err
	}

	repo := repositoryimpl.NewCorpSigning(
		mongodb.DAO(cfg.Mongodb.Collections.CorpSigning),
	)

	// corp
	models.RegisterCorpSigningAdapter(
		adapter.NewCorpSigningAdapter(
			app.NewCorpSigningService(repo),
		),
	)

	// employee
	models.RegisterEmployeeSigningAdapter(
		adapter.NewEmployeeSigningAdapter(
			app.NewEmployeeSigningService(repo),
		),
	)

	// individual
	individual := repositoryimpl.NewIndividualSigning(
		mongodb.DAO(cfg.Mongodb.Collections.IndividualSigning),
	)

	models.RegisterIndividualSigningAdapter(
		adapter.NewIndividualSigningAdapter(app.NewIndividualSigningService(
			individual,
		)),
	)

	// access token
	models.RegisterAccessTokenAdapter(
		accesstokenservice.NewAccessTokenService(symmetric),
	)

	// link
	linkRepo := repositoryimpl.NewLink(
		mongodb.DAO(cfg.Mongodb.Collections.Link),
	)

	models.RegisterLinkAdapter(
		adapter.NewLinkAdapter(
			app.NewLinkService(linkRepo),
		),
	)

	return nil
}
