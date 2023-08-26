package utils

// Converitr un slice de interface{} en []string
func ConvertSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if str, ok := v.(string); ok {
			result[i] = str
		} else {
			// Gérer le cas où l'élément n'est pas une chaîne de caractères
			// Par exemple, vous pouvez attribuer une valeur par défaut ou ignorer l'élément.
			result[i] = "" // Valeur par défaut
		}
	}
	return result
}
