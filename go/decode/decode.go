package decode

import (
	"cipher/a1z26"
	"cipher/atbash"
	"cipher/binary"
	"cipher/caesar"
	"cipher/fence"
	"cipher/hex"
	"cipher/monosubstitution"
	"cipher/words"
)

func Decode(input string, key string, cipherType string, cod string) string {
	dictionary := words.ReadWords()
	switch cipherType {
	case "caesar":
		return caesar.GetMessage(dictionary, input, key, cod)
	case "substitution":
		return monosubstitution.GetMessage(input, key, cod)
	case "atbash":
		return atbash.GetMessage(input, cod)
	case "binary":
		return binary.GetMessage(input, cod)
	case "hex":
		return hex.GetMessage(input, cod)
	case "fence":
		return fence.GetMessage(dictionary, input, key, cod)
	default:
		return a1z26.GetMessage(input, cod)
	}
}
