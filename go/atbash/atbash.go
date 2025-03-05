package atbash

import "strings"

var keyMapCap = map[byte]byte{
	'A': 'Z', 'B': 'Y', 'C': 'X', 'D': 'W', 'E': 'V', 'F': 'U', 'G': 'T',
	'H': 'S', 'I': 'R', 'J': 'Q', 'K': 'P', 'L': 'O', 'M': 'N', 'N': 'M',
	'O': 'L', 'P': 'K', 'Q': 'J', 'R': 'I', 'S': 'H', 'T': 'G', 'U': 'F',
	'V': 'E', 'W': 'D', 'X': 'C', 'Y': 'B', 'Z': 'A',
}

var keyMapReg = map[byte]byte{
	'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 'g': 't',
	'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm',
	'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g', 'u': 'f',
	'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b', 'z': 'a',
}

func GetMessage(text string, cod string) string {
	codeType := strings.ToLower(cod)

	if codeType == "code" || codeType == "decode" {
		return encodeCipher(text)
	}
	return "We could not understand the instruction of " + cod
}

func encodeCipher(text string) string {
	letters := []byte(text)
	message := []string{}

	for _, letter := range letters {
		if letter >= 65 && letter <= 90 {
			message = append(message, string(keyMapCap[letter]))
		} else if letter >= 97 && letter <= 122 {
			message = append(message, string(keyMapReg[letter]))
		} else {
			message = append(message, string(letter))
		}
	}

	return strings.Join(message, "")
}
