//
// This file is forked from https://github.com/pinzolo/casee
//
// The MIT License (MIT)
//
// Copyright (c) 2016 pinzolo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify, merge,
// publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons
// to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or
// substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
// FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package strutils

import (
	"strings"
	"unicode"
)

// ToSnakeCase convert argument to snake_case style string.
// If argument is empty, return itself.
func ToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "_")
}

// IsSnakeCase check whether argument is snake_case style string, return true.
func IsSnakeCase(s string) bool {
	if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByLowerAndDigit(s)

}

// ToChainCase convert argument to chain-case style string.
// If argument is empty, return itself.
func ToChainCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "-")
}

// IsChainCase whether argument is chain-case style string, return true.
func IsChainCase(s string) bool {
	if strings.Contains(s, "-") {
		fields := strings.Split(s, "-")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByLowerAndDigit(s)
}

// ToCamelCase convert argument to camelCase style string
// If argument is empty, return itself
func ToCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		if i != 0 {
			fields[i] = toUpperFirstRune(f)
		}
	}
	return strings.Join(fields, "")
}

// IsCamelCase whether argument is camelCase style string, return true.
// If first character is digit, always returns false
func IsCamelCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByAlphanumeric(s) && isFirstRuneLower(s)
}

// ToPascalCase convert argument to PascalCase style string
// If argument is empty, return itself
func ToPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		fields[i] = toUpperFirstRune(f)
	}
	return strings.Join(fields, "")
}

// IsPascalCase whether argument is PascalCase style string, return true.
// If first character is digit, always returns false
func IsPascalCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByAlphanumeric(s) && isFirstRuneUpper(s)
}

// ToFlatCase convert argument to flatcase style string
// If argument is empty, return itself
func ToFlatCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "")
}

// IsFlatCase whether argument is flatcase style string, return true.
// If first character is digit, always returns false
func IsFlatCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByLowerAndDigit(s)
}

// ToUpperCase convert argument to UPPER_CASE style string.
// If argument is empty, return itself.
func ToUpperCase(s string) string {
	return strings.ToUpper(ToSnakeCase(s))
}

// IsUpperCase whether argument is UPPER_CASE style string, return true.
func IsUpperCase(s string) bool {
	if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByUpperAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByUpperAndDigit(s)
}

func isMadeByLowerAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByUpperAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByAlphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isFirstRuneUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsUpper(getRuneAt(s, 0))
}

func isFirstRuneLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsLower(getRuneAt(s, 0))
}

func isFirstRuneDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsDigit(getRuneAt(s, 0))
}

func getRuneAt(s string, i int) rune {
	if len(s) == 0 {
		return 0
	}

	rs := []rune(s)
	return rs[0]
}

func splitToLowerFields(s string) []string {
	defaultCap := len([]rune(s)) / 3
	fields := make([]string, 0, defaultCap)

	for _, sf := range strings.Fields(s) {
		for _, su := range strings.Split(sf, "_") {
			for _, sh := range strings.Split(su, "-") {
				for _, sc := range Split(sh) {
					fields = append(fields, strings.ToLower(sc))
				}
			}
		}
	}
	return fields
}

func toUpperFirstRune(s string) string {
	rs := []rune(s)
	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}

//---------The following code does not exist in the original version-------------

// IsChainCaseForRouting whether argument is chain-case style string, return true.
func IsChainCaseForRouting(s string) bool {
	if strings.Contains(s, "-") {
		fields := strings.Split(s, "-")
		for _, field := range fields {
			if !isMadeByLowerAndDigitForRouting(field) {
				return false
			}
		}
		return true
	}
	return isMadeByLowerAndDigitForRouting(s)
}

func isMadeByLowerAndDigitForRouting(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) && !(r == rune(':')) && !(r == rune('/')) {
			return false
		}
	}

	return true
}

// ToChainCaseForRouting convert argument to chain case style string.
// If argument is empty, return itself.
func ToChainCaseForRouting(s string) string {
	if len(s) == 0 {
		return s
	}

	str := ""
	for i, f := range splitToLowerFields(s) {
		if i == 0 {
			str = f
			continue
		} else if f == "/:" || strings.HasSuffix(str, "/:") {
			str = str + f
			continue
		}
		str = str + "-" + f
	}
	return str
}
