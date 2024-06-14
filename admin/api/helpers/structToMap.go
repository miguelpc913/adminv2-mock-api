//We convert structs to maps in order to update properly the booleans in gorm.

package helpers

import (
	"encoding/json"
	"unicode"
)

// Function to capitalize the first letter of a string
func capitalize(str string) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Function to capitalize the keys of a map
func CapitalizeMapKeys(m map[string]interface{}) map[string]interface{} {
	capitalizedMap := make(map[string]interface{})
	for k, v := range m {
		capitalizedKey := capitalize(k)
		capitalizedMap[capitalizedKey] = v
	}
	return capitalizedMap
}

func StructToMap[T any](str T) map[string]interface{} {
	result := make(map[string]interface{})
	inrec, _ := json.Marshal(str)
	json.Unmarshal(inrec, &result)
	return CapitalizeMapKeys(result)
}
