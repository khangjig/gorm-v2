package util

import (
	"reflect"
	"regexp"
	"strings"
)

func SplitObjects(objArr []interface{}, size int) [][]interface{} {
	var (
		chunkSet [][]interface{}
		chunk    []interface{}
	)

	for len(objArr) > size {
		chunk, objArr = objArr[:size], objArr[size:]
		chunkSet = append(chunkSet, chunk)
	}

	if len(objArr) > 0 {
		chunkSet = append(chunkSet, objArr)
	}

	return chunkSet
}

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func GetDiff2Slices(slice1 []string, slice2 []string) []string {
	var diff []string
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false

			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true

					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func ConvertToNonAccentVietnamese(text string) string {
	type replace struct {
		regex string
		value string
	}

	replaces := []replace{
		{
			regex: "à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ",
			value: "a",
		},
		{
			regex: "è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ",
			value: "e",
		},
		{
			regex: "ì|í|ị|ỉ|ĩ",
			value: "i",
		},
		{
			regex: "ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ",
			value: "o",
		},
		{
			regex: "ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ",
			value: "u",
		},
		{
			regex: "ỳ|ý|ỵ|ỷ|ỹ",
			value: "y",
		},
		{
			regex: "đ",
			value: "d",
		},
	}

	for i := range replaces {
		re := regexp.MustCompile(replaces[i].regex)
		text = re.ReplaceAllString(text, replaces[i].value)

		re = regexp.MustCompile(strings.ToUpper(replaces[i].regex))
		text = re.ReplaceAllString(text, strings.ToUpper(replaces[i].value))
	}

	return text
}

func RemoveDuplicateStringSlice(values []string) []string {
	newValues := []string{}
	exist := make(map[string]bool)

	for _, value := range values {
		if exist[value] {
			continue
		}

		newValues = append(newValues, value)
		exist[value] = true
	}

	return newValues
}

// SQLEscapeString .
func SQLEscapeString(val string) string {
	replacer := strings.NewReplacer(
		"\\0", "\\\\0",
		"\n", "\\n",
		"\r", "\\r",
		"\x1a", "\\Z",
		`"`, `\"`,
		"'", `\'`,
		"\\", "\\\\",
	)

	return replacer.Replace(val)
}
