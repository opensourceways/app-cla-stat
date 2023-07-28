package repositoryimpl

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	mongodbCmdOr        = "$or"
	mongodbCmdIn        = "$in"
	mongodbCmdLt        = "$lt"
	mongodbCmdElemMatch = "$elemMatch"
)

type dao interface {
	IsDocNotExists(error) bool

	DocIdFilter(s string) (bson.M, error)

	GetDoc(filter, project bson.M, result interface{}) error
	GetDocs(filter, project bson.M, result interface{}) error
}

func linkIdFilter(v string) bson.M {
	return bson.M{
		fieldLinkId: v,
	}
}

func childField(fields ...string) string {
	return strings.Join(fields, ".")
}
