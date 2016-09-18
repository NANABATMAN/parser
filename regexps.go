package main

import (
	"regexp"
)

var (
	wordCharacterRegExp    = regexp.MustCompile(`^\w$`)
	littleLetterRegExp     = regexp.MustCompile(`^[a-z]$`)
	urlRegExp              = regexp.MustCompile(`^(https?:\/\/)?([\w\.]+)\.([a-z]{2,6}\.?)(\/[\w\.]*)*\/?$`)
	allowdedCharacterInUrl = regexp.MustCompile(`^[A-Za-z0-9-._~:\/?#[\]@!$&'()*+,;=%]$`)
)

func isWordCharacter(char string) bool {
	return wordCharacterRegExp.MatchString(char)
}

func isLittleLetter(char string) bool {
	return littleLetterRegExp.MatchString(char)
}

func isUrl(s string) bool {
	return urlRegExp.MatchString(s)
}

func isAllowdedCharacterInUrl(char string) bool {
	return allowdedCharacterInUrl.MatchString(char)
}
