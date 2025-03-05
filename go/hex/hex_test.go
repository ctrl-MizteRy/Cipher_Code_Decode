package hex

import "testing"

func TestEncodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"hello world",
			"68 65 6C 6C 6F 20 77 6F 72 6C 64",
		},
		{
			"This is a HEXADECIMAL msg!!!",
			"54 68 69 73 20 69 73 20 61 20 48 45 58 41 44 45 43 49 4D 41 4C 20 6D 73 67 21 21 21",
		},
		{
			"beep boop bop",
			"62 65 65 70 20 62 6F 6F 70 20 62 6F 70",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "code")

		if test.expect != expect {
			t.Fatalf("Mismatch encoding error: expect=%s got=%s", test.expect, expect)
		}
	}
}

func TestDecodingCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"62 65 65 70 20 62 6F 6F 70 20 62 6F 70",
			"beep boop bop",
		},
		{
			"54 68 69 73 20 69 73 20 61 20 48 45 58 41 44 45 43 49 4D 41 4C 20 6D 73 67 21 21 21",
			"This is a HEXADECIMAL msg!!!",
		},
		{
			"68 65 6C 6C 6F 20 77 6F 72 6C 64",
			"hello world",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "decode")

		if expect != test.expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
