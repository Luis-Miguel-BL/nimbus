package vo

type Schema struct {
	schema map[SchemaKey]SchemaKind
}

type SchemaKey string
type SchemaKind string

const (
	SchemaKindString SchemaKind = "string"
	SchemaKindNumber SchemaKind = "number"
	SchemaKindBool   SchemaKind = "boolean"
)
