package repository

import "github.com/opensourceways/app-cla-stat/signing/domain"

type IndividualSigning interface {
	FindAll(string) ([]domain.IndividualSigning, error)
}
