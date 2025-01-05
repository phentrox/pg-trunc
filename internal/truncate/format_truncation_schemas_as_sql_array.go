package truncate

import "strings"

func FormatTruncationSchemasAsSqlArray(schemas []string) string {
	return "{" + strings.Join(schemas, ",") + "}"
	//if len(schemas) == 0 {
	//	return "ARRAY[]"
	//}
	//
	//stringBuilder := strings.Builder{}
	//
	//stringBuilder.WriteString("ARRAY[")
	//
	//for _, schema := range schemas {
	//	stringBuilder.WriteString("'")
	//	stringBuilder.WriteString(schema)
	//	stringBuilder.WriteString("'")
	//	stringBuilder.WriteString(", ")
	//}
	//stringBuilderString := stringBuilder.String()
	//
	//// remove comma and whitespace at the end of the string
	//schemaList := stringBuilderString[:len(stringBuilderString)-2]
	//
	//// add closing bracket
	//return schemaList + "]"
}
