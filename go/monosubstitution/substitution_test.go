package monosubstitution

import "testing"

func TestEncoding(t *testing.T) {
	tests := []struct {
		input  string
		key    string
		expect string
	}{
		{
			"Maybe? This! Will.. Go wrong",
			"zyxwvutsrqponmlkjihgfedcba",
			"Nzbyv? Gsrh! Droo.. Tl dilmt",
		},
		{
			"Hello World",
			"iajktqpsdowmchxlzgfenvruby",
			"Stmmx Rxgmk",
		},
		{
			"According to all known laws of aviation, there is no way a bee should be able to fly. Its wings are too small to get its fat little body off the ground. The bee, of course, flies anyway because bees don't care what humans think is impossible.",
			"svrfnubqomzlckegyhiwpjtxad",
			"Srrehfokb we sll zketk lsti eu sjoswoek, wqnhn oi ke tsa s vnn iqeplf vn svln we ula. Owi tokbi shn wee icsll we bnw owi usw lowwln vefa euu wqn bhepkf. Wqn vnn, eu rephin, uloni skatsa vnrspin vnni fek'w rshn tqsw qpcski wqokz oi ocgeiiovln.",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, test.key, "code")

		if expect != test.expect {
			t.Fatalf("Mismatch encoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}

func TestDecoding(t *testing.T) {
	tests := []struct {
		input  string
		key    string
		expect string
	}{
		{
			"Nzbyv? Gsrh! Droo.. Tl dilmt",
			"zyxwvutsrqponmlkjihgfedcba",
			"Maybe? This! Will.. Go wrong",
		},
		{
			"Stmmx Rxgmk",
			"iajktqpsdowmchxlzgfenvruby",
			"Hello World",
		},
		{
			"Srrehfokb we sll zketk lsti eu sjoswoek, wqnhn oi ke tsa s vnn iqeplf vn svln we ula. Owi tokbi shn wee icsll we bnw owi usw lowwln vefa euu wqn bhepkf. Wqn vnn, eu rephin, uloni skatsa vnrspin vnni fek'w rshn tqsw qpcski wqokz oi ocgeiiovln.",
			"svrfnubqomzlckegyhiwpjtxad",
			"According to all known laws of aviation, there is no way a bee should be able to fly. Its wings are too small to get its fat little body off the ground. The bee, of course, flies anyway because bees don't care what humans think is impossible.",
		},
	}

	for _, test := range tests {
		expect := GetMessage(test.input, test.key, "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decoding error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
