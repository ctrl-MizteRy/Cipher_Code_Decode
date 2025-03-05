package caesar

import (
	"cipher/words"
	"testing"
)

var dict = words.ReadWords()

func TestCaeserEncode(t *testing.T) {
	tests := []struct {
		text     string
		key      string
		expected string
	}{
		{
			"Hello World",
			"7",
			"Olssv Dvysk",
		},
		{
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"15",
			"PQRSTUVWXYZABCDEFGHIJKLMNO",
		},
		{
			"abcdefghijklmnopqrstuvwxyz",
			"10",
			"klmnopqrstuvwxyzabcdefghij",
		},
		{
			"this is a test :)",
			"26",
			"this is a test :)",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dict, test.text, test.key, "code")

		if test.expected != expect {
			t.Fatalf("Mismatch encoding error: expect=%s, got=%s", test.expected, expect)
		}
	}
}

func TestCaesarDecodeWithKey(t *testing.T) {
	tests := []struct {
		input    string
		key      string
		expected string
	}{
		{
			"klmnopqrstuvwxyzabcdefghij",
			"10",
			"abcdefghijklmnopqrstuvwxyz",
		},
		{
			"PQRSTUVWXYZABCDEFGHIJKLMNO",
			"15",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			"Olssv Dvysk",
			"7",
			"Hello World",
		},
		{
			"Dro kwlob nbyzvod rexq pbyw dro lbkxmr, bokmrsxq pevvxocc kxn bokni dy nbyz. Sd gksdon. Grsvo wkxi yp dro ydrob nbyzvodc gobo ckdscpson dy pybw kc lsq kc droi myevn kxn bovokco, drsc nbyzvod rkn ydrob zvkxc. Sd gkxdon dy lo zkbd yp rscdybi. Sd gkxdon dy lo bowowlobon vyxq kpdob kvv dro ydrob nbyzvodc rkn nsccyvfon sxdy rscdybi",
			"10",
			"The amber droplet hung from the branch, reaching fullness and ready to drop. It waited. While many of the other droplets were satisfied to form as big as they could and release, this droplet had other plans. It wanted to be part of history. It wanted to be remembered long after all the other droplets had dissolved into history",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dict, test.input, test.key, "decode")

		if test.expected != expect {
			t.Fatalf("Mismatch decoding, expect=%s, got=%s", test.expected, expect)
		}
	}
}

func TestCaesarDecodeWithoutKey(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			"Dro kwlob nbyzvod rexq pbyw dro lbkxmr, bokmrsxq pevvxocc kxn bokni dy nbyz. Sd gksdon. Grsvo wkxi yp dro ydrob nbyzvodc gobo ckdscpson dy pybw kc lsq kc droi myevn kxn bovokco, drsc nbyzvod rkn ydrob zvkxc. Sd gkxdon dy lo zkbd yp rscdybi. Sd gkxdon dy lo bowowlobon vyxq kpdob kvv dro ydrob nbyzvodc rkn nsccyvfon sxdy rscdybi",
			"The amber droplet hung from the branch, reaching fullness and ready to drop. It waited. While many of the other droplets were satisfied to form as big as they could and release, this droplet had other plans. It wanted to be part of history. It wanted to be remembered long after all the other droplets had dissolved into history",
		},
		{
			"Olssv Dvysk",
			"Hello World",
		},
		{
			"drsc sc K DOCD!!!! Iokr",
			"this is A TEST!!!! Yeah",
		},
		{
			"hello there :)",
			"hello there :)",
		},
		{
			"Thfil? Aopz! Dpss.. Nv dyvun",
			"Maybe? This! Will.. Go wrong",
		},
	}

	for _, test := range tests {
		expect := GetMessage(dict, test.input, "", "decode")

		if test.expect != expect {
			t.Fatalf("Mismatch decode error: expect=%s, got=%s", test.expect, expect)
		}
	}
}
