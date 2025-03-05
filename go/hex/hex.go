package hex

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func GetMessage(input string, cod string) string {
	cod = strings.ToLower(cod)

	if cod == "code" {
		return encodeCipher(input)
	} else if cod == "decode" {
		return decodeCipher(input)
	} else {
		return "We could not understand the command of " + cod + ", only code and decode"
	}
}

func encodeCipher(input string) string {
	text := []byte(input)
	var msg = []string{}
	val := hex.EncodeToString(text)

	for i := 0; i < len(val); i += 2 {
		msg = append(msg, strings.ToUpper(val[i:i+2]))
	}
	return strings.Join(msg, " ")
}

func decodeCipher(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	text, err := hex.DecodeString(input)
	if err != nil {
		panic(fmt.Sprintf("%s does not seem to be a hex message", input))
	}

	var msg = []string{}
	for _, letter := range text {
		msg = append(msg, string(letter))
	}

	return strings.Join(msg, "")
}
