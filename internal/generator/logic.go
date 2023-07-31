/* this file contains the logic for the password generator */

package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func GeneratePassword(length int, useUpper, useLower, useNumbers, useSpecial bool, filter string) string {
	chars := ""
	if useUpper {
		chars += upperChars
	}
	if useLower {
		chars += lowerChars
	}
	if useNumbers {
		chars += numberChars
	}
	if useSpecial {
		chars += specialChars
	}
	if chars == "" {
		return ""
	}

	/* funtion to filter the characters according to the filter */
	filterFunc := func(r rune) rune {
		if !strings.ContainsRune(filter, r) {
			return r
		}
		return rune(0)
	}

	/* if filter is not empty, apply the filter to the characters */
	if filter != "" {
		chars = MapFilter(filterFunc, chars)
	}

	password := make([]byte, length)
	maxIdx := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, maxIdx)
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}
		password[i] = chars[idx.Int64()]
	}

	return string(password)
}

/* MapFilter function to apply the mapping to a string */
func MapFilter(f func(rune) rune, s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		result[i] = f(r)
	}
	return string(result)
}
