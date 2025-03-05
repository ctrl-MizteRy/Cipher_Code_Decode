package atbash

import "testing"

func TestEncodeCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"PAPER JAM DIPPER SAYS: \"AUUGHWXQHGADSADUH!\"",
			"KZKVI QZN WRKKVI HZBH: \"ZFFTSDCJSTZWHZWFS!\"",
		},
		{
			"E. Pluribus trembley.",
			"V. Kofiryfh givnyovb.",
		},
		{
			"Heavy is the head that wears the fez",
			"Svzeb rh gsv svzw gszg dvzih gsv uva",
		},
		{
			"!@#$%^&*()_+",
			"!@#$%^&*()_+",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "code")

		if expect != test.expect {
			t.Fatalf("Mismatch encoding result: expect=%s, got=%s", test.expect, expect)
		}
	}

}

func TestDecodingCipher(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"Svzeb rh gsv svzw gszg dvzih gsv uva",
			"Heavy is the head that wears the fez",
		},
		{
			"V. Kofiryfh givnyovb.",
			"E. Pluribus trembley.",
		},
		{
			"KZKVI QZN WRKKVI HZBH: \"ZFFTSDCJSTZWHZWFS!\"",
			"PAPER JAM DIPPER SAYS: \"AUUGHWXQHGADSADUH!\"",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
