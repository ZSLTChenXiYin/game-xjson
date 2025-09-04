package converter

type ColumnType string

const (
	COLUMN_TYPE_STRING ColumnType = "string"
	COLUMN_TYPE_INT    ColumnType = "int"
	COLUMN_TYPE_FLOAT  ColumnType = "float"
	COLUMN_TYPE_BOOL   ColumnType = "bool"
	COLUMN_TYPE_JSON   ColumnType = "json"
)

var (
	type_map = map[string]ColumnType{
		"string":  COLUMN_TYPE_STRING,
		"text":    COLUMN_TYPE_STRING,
		"int":     COLUMN_TYPE_INT,
		"integer": COLUMN_TYPE_INT,
		"float":   COLUMN_TYPE_FLOAT,
		"double":  COLUMN_TYPE_FLOAT,
		"number":  COLUMN_TYPE_FLOAT,
		"bool":    COLUMN_TYPE_BOOL,
		"boolean": COLUMN_TYPE_BOOL,
		"json":    COLUMN_TYPE_JSON,
		"object":  COLUMN_TYPE_JSON,
		"array":   COLUMN_TYPE_JSON,
		"list":    COLUMN_TYPE_JSON,
	}
)
