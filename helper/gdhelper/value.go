package gdhelper

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func IIF[T any](v bool, trueVal T, falseVal T) T {
	if v {
		return trueVal
	}
	return falseVal
}

func SetIfHasValue[T comparable](dst *T, val T, zero T) {
	if val != zero {
		*dst = val
	}
}

func StringToUnit(input string) uint {
	num, err := strconv.ParseUint(input, 10, 64)
	if err == nil {
		return uint(num)
	}
	return uint(0)
}

func UnitToString(input uint) string {
	return strconv.FormatUint(uint64(input), 10)
}

func EnableToggle(input string) bool {
	return input == "1"
}

func RemoveAccents(s string) string {
	replacements := map[*regexp.Regexp]string{
		regexp.MustCompile(`[áàâãªä]`): "a",
		regexp.MustCompile(`[ÁÀÂÃÄ]`):  "A",
		regexp.MustCompile(`[ÍÌÎÏ]`):   "I",
		regexp.MustCompile(`[íìîï]`):   "i",
		regexp.MustCompile(`[éèêë]`):   "e",
		regexp.MustCompile(`[ÉÈÊË]`):   "E",
		regexp.MustCompile(`[óòôõºö]`): "o",
		regexp.MustCompile(`[ÓÒÔÕÖ]`):  "O",
		regexp.MustCompile(`[úùûü]`):   "u",
		regexp.MustCompile(`[ÚÙÛÜ]`):   "U",
		regexp.MustCompile(`ç`):        "c",
		regexp.MustCompile(`Ç`):        "C",
		regexp.MustCompile(`ñ`):        "n",
		regexp.MustCompile(`Ñ`):        "N",
		regexp.MustCompile(`–`):        "-",
		regexp.MustCompile(`[’‘‹›‚]`):  "",
		regexp.MustCompile(`[“”«»„]`):  "",
	}

	for re, repl := range replacements {
		s = re.ReplaceAllString(s, repl)
	}

	reSpecial := regexp.MustCompile(`[!@#$%^*():|<>~]`)
	s = reSpecial.ReplaceAllString(s, "")

	return s
}

func AddSlashes(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	s = strings.ReplaceAll(s, "\x00", `\0`)
	return s
}

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func SubString(s string, start, length int) string {
	runes := []rune(s)
	if start > len(runes) {
		return ""
	}
	end := start + length
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

func AppendIfNotExistPtr[T comparable](slice *[]T, value T) {
	if !slices.Contains(*slice, value) {
		*slice = append(*slice, value)
	}
}

func Ptr[T any](v T) *T { return &v }
