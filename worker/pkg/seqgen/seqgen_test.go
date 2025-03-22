package seqgen

import "testing"

func TestNext(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
	maxWordLen := 3
	offset, nWords := int64(1260), int64(109)
	sg := New(alphabet, maxWordLen, offset, nWords)

	expected := []string{
		"8a", "8b", "8c", "8d", "8e", "8f", "8g", "8h", "8i", "8j", "8k", "8l", "8m",
		"8n", "8o", "8p", "8q", "8r", "8s", "8t", "8u", "8v", "8w", "8x", "8y", "8z",
		"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
		"9a", "9b", "9c", "9d", "9e", "9f", "9g", "9h", "9i", "9j", "9k", "9l", "9m",
		"9n", "9o", "9p", "9q", "9r", "9s", "9t", "9u", "9v", "9w", "9x", "9y", "9z",
		"90", "91", "92", "93", "94", "95", "96", "97", "98", "99",
		"aaa", "aab", "aac", "aad", "aae", "aaf", "aag", "aah", "aai", "aaj", "aak", "aal", "aam",
		"aan", "aao", "aap", "aaq", "aar", "aas", "aat", "aau", "aav", "aaw", "aax", "aay", "aaz",
		"aa0", "aa1", "aa2", "aa3", "aa4", "aa5", "aa6", "aa7", "aa8", "aa9",
		"aba",
	}

	for currentWordIndex := offset; currentWordIndex < offset+nWords; currentWordIndex++ {
		nextWord, err := sg.Next()

		if err != nil {
			t.Errorf("Unexpected error while generating next word: %v", err)
		}

		if nextWord != expected[currentWordIndex-offset] {
			t.Errorf("Expected %s, got %s", expected[currentWordIndex-offset], nextWord)
		}
	}

	nextWord, err := sg.Next()
	if err == nil || err.Error() != "word out of range" || nextWord != "" {
		t.Errorf("Expected error while generating next word")
	}
}
