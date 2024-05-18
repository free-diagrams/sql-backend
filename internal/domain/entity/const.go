package entity

import (
	"errors"
	internalerr "github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/free-diagrams/sql-backend/pkg/errs"
)

type ContextKey string

const (
	ContextLang        ContextKey = "ctx_lang"
	ContextUser        ContextKey = "ctx_user"
	ContextTransaction ContextKey = "ctx_transaction"
)

type AccessType string

const (
	AccessOwner   AccessType = "owner"
	AccessEditor  AccessType = "editor"
	AccessViewer  AccessType = "viewer"
	AccessUnknown AccessType = "unknown"
)

func (c AccessType) Value() string {
	return string(c)
}

func ParseAccessType(value string) (AccessType, error) {
	switch value {
	case AccessOwner.Value():
		return AccessOwner, nil
	case AccessEditor.Value():
		return AccessEditor, nil
	case AccessViewer.Value():
		return AccessViewer, nil
	default:
		return AccessUnknown, errs.WrapErrorError(internalerr.Parsing, errors.New("unknown access type"))
	}
}

type RelationType string

const (
	RelationMany    RelationType = "many"
	RelationOne     RelationType = "one"
	RelationUnknown RelationType = "unknown"
)

func (c RelationType) Value() string {
	return string(c)
}

func ParseRelationType(value string) (RelationType, error) {
	switch value {
	case RelationMany.Value():
		return RelationMany, nil
	case RelationOne.Value():
		return RelationOne, nil
	default:
		return RelationUnknown, errs.WrapErrorError(internalerr.Parsing, errors.New("unknown relation type"))
	}
}

type NotImplemented struct {
}
