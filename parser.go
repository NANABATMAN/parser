package main

import (
	"bytes"
	"encoding/json"
)

func Parse(message string) *bytes.Buffer {
	messageContent := MessageContent{}

	messageLenght := len(message)

	for i := 0; i < messageLenght; i++ {
		currentChar := string(message[i])
		switch currentChar {
		case AT_LEXEMA:
			mention, continueFrom := findMention(message, i)
			i = continueFrom
			if mention != EMPTY_STRING {
				messageContent.AddMention(mention)
			}
		case OPEN_BRACKET_LEXEMA:
			emoticon, continueFrom := findEmoticon(message, i)
			i = continueFrom
			if emoticon != EMPTY_STRING {
				messageContent.AddEmoticon(emoticon)
			}
		case LITTLE_H_LEXEMA:
			if i+4 < messageLenght && message[i:i+4] == HTTP_LEXEMA {
				url, continueFrom := findUrl(message, i+4)
				i = continueFrom
				if url != EMPTY_STRING {
					messageContent.AddLink(url)
				}
			}
		}
	}

	jsonBuffer := new(bytes.Buffer)
	json.NewEncoder(jsonBuffer).Encode(messageContent)
	return jsonBuffer
}

func findMention(message string, currentIndex int) (string, int) {
	var mention bytes.Buffer
	mention.WriteString(AT_LEXEMA)

	messageLenght := len(message)
	startIdnex := currentIndex + 1

	if startIdnex < messageLenght {
		for i := startIdnex; i < messageLenght; i++ {
			currentChar := string(message[i])
			if isWordCharacter(currentChar) {
				mention.WriteString(currentChar)
				if i == messageLenght-1 {
					return mention.String(), i
				}
			} else {
				if len(mention.String()) == 1 {
					return EMPTY_STRING, i - 1
				} else {
					return mention.String(), i - 1
				}
			}
		}
	}

	return EMPTY_STRING, currentIndex
}

func findEmoticon(message string, currentIndex int) (string, int) {
	var emoticon bytes.Buffer
	emoticon.WriteString(OPEN_BRACKET_LEXEMA)

	messageLenght := len(message)
	startIdnex := currentIndex + 1

	if startIdnex < messageLenght {
		for i := startIdnex; i < messageLenght; i++ {
			currentChar := string(message[i])
			if isLittleLetter(currentChar) {
				emoticon.WriteString(currentChar)
			} else {
				emoticonLength := len(emoticon.String())
				if emoticonLength > 1 && emoticonLength <= EMOTICON_MAX_LENGTH+1 &&
					currentChar == CLOSE_BRACKET_LEXEMA {
					emoticon.WriteString(CLOSE_BRACKET_LEXEMA)
					return emoticon.String(), i
				} else {
					return EMPTY_STRING, i - 1
				}
			}
		}
	}

	return EMPTY_STRING, currentIndex
}

func findUrl(message string, currentIndex int) (string, int) {
	var url bytes.Buffer
	url.WriteString(HTTP_LEXEMA)

	messageLenght := len(message)

	for i := currentIndex; i < messageLenght; i++ {
		currentChar := string(message[i])
		if isAllowdedCharacterInUrl(currentChar) {
			url.WriteString(currentChar)
			if i == messageLenght-1 {
				if !isUrl(url.String()) {
					return EMPTY_STRING, i
				} else {
					return url.String(), i
				}
			}
		} else {
			if !isUrl(url.String()) {
				return EMPTY_STRING, i - 1
			} else {
				return url.String(), i - 1
			}
		}
	}

	return EMPTY_STRING, currentIndex
}
