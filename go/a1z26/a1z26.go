package a1z26

import (
	"log"
	"strconv"
	"strings"
)

func GetMessage(text string, cod string) string {
	direction := strings.ToLower(cod)

	if direction == "code" {
		return encodeCipher(text)
	} else if direction == "decode" {
		return decodeCipher(text)
	} else {
		return ("unknown type " + direction)
	}
}

func encodeCipher(input string) string {
	msg := []byte(input)
	message := []string{}

	for _, letter := range msg {
		if letter >= 65 && letter <= 90 {
			l := (letter % 65) + 1
			message = append(message, strconv.Itoa(int(l)))
		} else if letter >= 97 && letter <= 122 {
			l := (letter % 97) + 1
			message = append(message, strconv.Itoa(int(l)))
		}
	}

	return strings.Join(message, " ")
}

func decodeCipher(input string) string {
	letters := strings.Split(input, " ")
	message := []string{}

	for _, letter := range letters {
		l, err := strconv.Atoi(letter)
		if err != nil {
			log.Fatalf("Unexpected value '%s' in the A1Z26 decoder", letter)
		}

		val := byte(l) - 1
		message = append(message, string(val+65))
	}

	return strings.Join(message, " ")
}
