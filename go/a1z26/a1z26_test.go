package a1z26

import "testing"

func TestEncodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"hello there",
			"8 5 12 12 15 20 8 5 18 5",
		},
		{
			"meeting will be at three",
			"13 5 5 20 9 14 7 23 9 12 12 2 5 1 20 20 8 18 5 5",
		},
		{
			"I'm writing this in golang",
			"9 13 23 18 9 20 9 14 7 20 8 9 19 9 14 7 15 12 1 14 7",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "code")

		if test.expect != expect {
			t.Fatalf("Mismatch encoding error: expec=%s, got=%s", test.expect, expect)
		}
	}
}

func TestDecodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"9 13 23 18 9 20 9 14 7 20 8 9 19 9 14 7 15 12 1 14 7",
			"I M W R I T I N G T H I S I N G O L A N G",
		},
		{
			"13 5 5 20 9 14 7 23 9 12 12 2 5 1 20 20 8 18 5 5",
			"M E E T I N G W I L L B E A T T H R E E",
		},
		{
			"8 5 12 12 15 23 15 18 12 4",
			"H E L L O W O R L D",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoing output, expect=%s, got=%s", test.expect, expect)
		}
	}
}
