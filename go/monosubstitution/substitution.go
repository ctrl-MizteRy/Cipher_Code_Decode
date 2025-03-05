package monosubstitution

import (
	"fmt"
	"strings"
)

func GetMessage(input string, key string, cod string) string {
	cipherType := strings.ToLower(cod)
	if len(key) != 26 {
		panic(fmt.Sprintf("Incorrect length: expect=26 got=%d", len(key)))
	}
	key = strings.ToUpper(key)

	if cipherType == "code" {
		return encodeCipher(input, key)
	} else if cipherType == "decode" {
		return decodeCipher(input, key)
	} else {
		return cod + " is not part of our function, only code or decode"
	}
}

func encodeCipher(input string, key string) string {
	substitute := make(map[byte]byte)
	keyByte := []byte(key)

	for i := 0; i < len(keyByte); i++ {
		substitute[byte(65+i)] = keyByte[i]
	}

	text := []byte(input)
	return substituting(text, substitute)
}

func decodeCipher(input string, key string) string {
	substitute := make(map[byte]byte)
	keyByte := []byte(key)

	for i := 0; i < len(keyByte); i++ {
		substitute[keyByte[i]] = byte(65 + i)
	}
	text := []byte(input)
	return substituting(text, substitute)
}

func substituting(text []byte, substitute map[byte]byte) string {
	msg := []string{}
	for _, letter := range text {
		if letter >= 65 && letter <= 90 {
			val, _ := substitute[letter]
			msg = append(msg, string(val))
		} else if letter >= 97 && letter <= 122 {
			upper := strings.ToUpper(string(letter))
			val, _ := substitute[[]byte(upper)[0]]
			msg = append(msg, strings.ToLower(string(val)))
		} else {
			msg = append(msg, string(letter))
		}
	}

	return strings.Join(msg, "")
}
