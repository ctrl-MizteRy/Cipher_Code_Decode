package fence

import (
	"cipher/words"
	"testing"
)

var dictionary = words.ReadWords()

func TestEncodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		key    string
		expect string
	}{
		{
			"Hello world",
			"4",
			"Hwe olordll",
		},
		{
			"This IS A test?? Maybe",
			"9",
			"T h?Mi?asty sbIeeSt  A",
		},
		{
			"Maybe? This! Will.. Go wrong",
			"6",
			"MsGai! oyh . bTW.wge ilrn?lo",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dict, test.input, test.key, "code")

		if test.expect != expect {
			t.Fatalf("Mismatch encoding, expect=%s, got=%s", test.expect, expect)
		}
	}
}

func TestDecodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"Hwe olordll",
			"Hello world",
		},
		{
			"MsGai! oyh . bTW.wge ilrn?lo",
			"Maybe? This! Will.. Go wrong",
		},
		{
			"T h?Mi?asty sbIeeSt  A",
			"This IS A test?? Maybe",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dictionary, test.input, "", "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}

func TestDecodeWithKey(t *testing.T) {
	tests := []struct {
		input  string
		key    string
		expect string
	}{
		{
			"Hwe olordll",
			"4",
			"Hello world",
		},
		{
			"T h?Mi?asty sbIeeSt  A",
			"9",
			"This IS A test?? Maybe",
		},
		{
			"MsGai! oyh . bTW.wge ilrn?lo",
			"6",
			"Maybe? This! Will.. Go wrong",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dictionary, test.input, test.key, "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
