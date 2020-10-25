package jpg2png

import "testing"

func TestIsExtensionOk(t *testing.T) {
	patterns := []struct {
		extension string
		expected  bool
	}{
		{".jpg", true},
		{".gif", true},
		{".png", false},
		{"jpg", false},
		{"gif", false},
		{"png", false},
	}

	for index, pattern := range patterns {
		actual := isExtensionOk(pattern.extension)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %t, actual %t", index, pattern.expected, actual)
		}
	}
}
