package gdhelper

import "strings"

func ParseTranslationKey(key, categoryDefault, keyDefault string) (string, string) {
	parts := strings.SplitN(key, ".", 2)
	if len(parts) != 2 {
		return categoryDefault, keyDefault
	}
	return parts[0], parts[1]
}

func FirstArgOrNil(data []map[string]interface{}) map[string]interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return nil
}
