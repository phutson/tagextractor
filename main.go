package tagextractor

import (
	"reflect"
	"regexp"
	"strings"
)

//ExtractTag extracts the tag associiated with the key and any special parameters seperated by a comma
func ExtractTag(key string, field reflect.StructField) (string, []string) {
	fieldVal := ""
	specials := make([]string, 0)
	rg := regexp.MustCompile(",")
	value, ok := field.Tag.Lookup(key)
	if !ok {
		return "", nil
	}
	if strings.Contains(value, ",") {
		splits := rg.Split(value, -1)
		if len(splits) > 0 {
			fieldVal = splits[0]
		}
	}
	return fieldVal, specials
}
