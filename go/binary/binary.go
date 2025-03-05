package binary

import (
	"fmt"
	"strconv"
	"strings"
)

func GetMessage(text string, cod string) string {
	cod = strings.ToLower(cod)

	if cod == "code" {
		return encodingCipher(text)
	} else if cod == "decode" {
		return decodingCipher(text)
	} else {
		return "We could not understand the instruction of" + cod + " only code or decode"
	}
}

func encodingCipher(input string) string {
	var msg = []string{}
	text := []byte(input)

	for _, letter := range text {
		binary := fmt.Sprintf("%08b", letter)
		msg = append(msg, binary)
	}

	return strings.Join(msg, " ")
}

func decodingCipher(input string) string {
	text := strings.Split(input, " ")
	var msg = []string{}
	if len(text) == 1 {
		text_len := len(text[0])
		message := text[0]
		for i := 0; i < text_len; i += 7 {
			if i+7 >= text_len {
				remain_len := len(message[i:])
				zeros := strings.Repeat("0", 8-remain_len)
				bit := zeros + message[i:]
				val, err := strconv.ParseInt(bit, 2, 8)
				if err != nil {
					panic(fmt.Sprintf("Cannot convert %s into a binary string", bit))
				}
				msg = append(msg, string(byte(val)))
			} else {
				bit := message[i : i+8]
				val, err := strconv.ParseInt(bit, 2, 8)
				if err != nil {
					panic(fmt.Sprintf("Cannot convert %s into a binary string", bit))
				}
				msg = append(msg, string(byte(val)))
			}
		}
	} else {
		for _, bit := range text {
			val, err := strconv.ParseInt(bit, 2, 8)
			if err != nil {
				panic(fmt.Sprintf("Cannot convert %s into a binary string", bit))
			}
			msg = append(msg, string(byte(val)))
		}
	}

	return strings.Join(msg, "")
}
