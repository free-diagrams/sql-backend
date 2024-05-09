package entity

type ContextKey string

const (
	ContextUser  ContextKey = "ctx_user"
	ContextLang  ContextKey = "ctx_lang"
	ContextError ContextKey = "ctx_error"
)

type AccessType string

const (
	AccessViewer AccessType = "viewer"
	AccessEditor AccessType = "editor"
	AccessOwner  AccessType = "owner"
)

func (c *AccessType) Value() string {
	return string(*c)
}

type RelationType string

const (
	RelationOne  RelationType = "one"
	RelationMany RelationType = "many"
)

func (c *RelationType) Value() string {
	return string(*c)
}
