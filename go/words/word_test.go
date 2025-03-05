package words

import "testing"

func TestGetDict(t *testing.T) {
	dict := ReadWords()

	if len(dict) != 10000 {
		t.Fatalf("Unexpect length, expect= 10000, got=%d", len(dict))
	}
}
