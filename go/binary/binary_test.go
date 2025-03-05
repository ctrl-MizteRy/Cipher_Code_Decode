package binary

import "testing"

func TestEncodecipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"hello world",
			"01101000 01100101 01101100 01101100 01101111 00100000 01110111 01101111 01110010 01101100 01100100",
		},
		{
			"This is a BINARY msg!!!!",
			"01010100 01101000 01101001 01110011 00100000 01101001 01110011 00100000 01100001 00100000 01000010 01001001 01001110 01000001 01010010 01011001 00100000 01101101 01110011 01100111 00100001 00100001 00100001 00100001",
		},
		{
			"Beep Boop Bop",
			"01000010 01100101 01100101 01110000 00100000 01000010 01101111 01101111 01110000 00100000 01000010 01101111 01110000",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "code")

		if test.expect != expect {
			t.Fatalf("Mismatch enconding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}

func TestDecodingCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"01000010 01100101 01100101 01110000 00100000 01000010 01101111 01101111 01110000 00100000 01000010 01101111 01110000",
			"Beep Boop Bop",
		},
		{
			"01010100 01101000 01101001 01110011 00100000 01101001 01110011 00100000 01100001 00100000 01000010 01001001 01001110 01000001 01010010 01011001 00100000 01101101 01110011 01100111 00100001 00100001 00100001 00100001",
			"This is a BINARY msg!!!!",
		},
		{
			"01101000 01100101 01101100 01101100 01101111 00100000 01110111 01101111 01110010 01101100 01100100",
			"hello world",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
