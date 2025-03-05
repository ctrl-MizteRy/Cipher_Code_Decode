package caesar

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetMessage(words map[string]byte, msg string, key string, cod string) string {
	var message string
	if strings.ToLower(cod) == "code" {
		message = encodeCaesar(msg, key)
	} else if strings.ToLower(cod) == "decode" {
		message = decodeCaesar(words, msg, key)
	}

	return message
}

func decodeCaesar(words map[string]byte, text string, key string) string {
	if len(key) == 0 {
		return decodeCaesarNoKey(words, text)
	} else {
		n, err := strconv.Atoi(key)
		if err != nil {
			panic(fmt.Sprintf("%s cannot be convert to an integer for the caesar key", key))
		}

		shift := byte(n)
		msg := []byte(text)

		for i := range len(msg) {
			if (msg[i] >= 65 && msg[i] <= 90) || (msg[i] >= 97 && msg[i] <= 122) {
				result := msg[i] - shift
				if result < 65 && msg[i] >= 65 && msg[i] <= 90 {
					newResult := 91 - (65 - result)
					msg[i] = newResult
				} else if result < 97 && msg[i] >= 97 && msg[i] <= 122 {
					newResult := 123 - (97 - result)
					msg[i] = newResult
				} else {
					msg[i] = result
				}
			}
		}
		return string(msg)
	}
}

func decodeCaesarNoKey(words map[string]byte, text string) string {
	for i := byte(0); i < 26; i++ {
		msg := []byte(text)
		for j := range len(msg) {
			if (msg[j] >= 65 && msg[j] <= 90) || (msg[j] >= 97 && msg[j] <= 122) {
				result := msg[j] - i
				if result < 65 && msg[j] >= 65 && msg[j] <= 90 {
					newResult := 91 - (65 - result)
					msg[j] = newResult
				} else if result < 97 && msg[j] >= 97 && msg[j] <= 122 {
					newResult := 123 - (97 - result)
					msg[j] = newResult
				} else {
					msg[j] = result
				}
			}
		}
		if isEnglishWords(words, string(msg)) {
			return string(msg)
		}
	}
	return "We could not find decode this message with Caesar"
}

func encodeCaesar(text string, key string) string {
	n, err := strconv.Atoi(key)
	if err != nil {
		panic(fmt.Sprintf("%s cannot be convert into an integer for the Caesar key", key))
	}

	shift := byte(n)
	msg := []byte(text)
	for i := range len(msg) {
		if (msg[i] >= 65 && msg[i] <= 90) || (msg[i] >= 97 && msg[i] <= 122) {
			result := msg[i] + shift
			if result > 122 && msg[i] >= 97 {
				newResult := 97 - (123 - result)
				msg[i] = newResult
			} else if result > 90 && msg[i] >= 65 && msg[i] <= 90 {
				newResult := 65 - (91 - result)
				msg[i] = newResult
			} else {
				msg[i] = result
			}
		}

	}

	return string(msg)
}

func isEnglishWords(dict map[string]byte, text string) bool {
	words := strings.Split(text, " ")
	words_len := len(words)
	word_count := 0

	for _, word := range words {
		regex := regexp.MustCompile("[!.?\":'()\\[\\]{}]+")
		newWord := regex.ReplaceAllString(word, "")
		if _, ok := dict[strings.ToLower(newWord)]; ok {
			word_count++
		}
	}

	if words_len < 10 && (float64(word_count)/float64(words_len)) > 0.5 {
		return true
	} else if float64(word_count)/float64(words_len) >= 0.7 {
		return true
	} else {
		return false
	}
}
