package truncate

func Truncate(schemas []string) {
	schemasAsSqlArray := FormatTruncationSchemasAsSqlArray(schemas)
	err := TruncateAllSchemasInList(schemasAsSqlArray)
	if err != nil {
		panic(err)
	}
}
