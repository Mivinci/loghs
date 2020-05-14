package loghs

// FieldType fieldtype
type FieldType int8

// Field types
const (
	UnknownType FieldType = iota
	StringType
	IntTpye
	Int64Type
	UintType
	Uint64Type
	Float32Type
	Float64Type
	DurationType
)

// Field field
type Field struct {
	Type   FieldType
	Key    string
	String string
	Int64  int64
}

// String string field
func String(key, value string) Field {
	return Field{Type: StringType, Key: key, String: value}
}

// Int64 int64 field
func Int64(key string, value int64) Field {
	return Field{Type: Int64Type, Key: key, Int64: value}
}
