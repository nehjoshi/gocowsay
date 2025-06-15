package cowsay

import (
	"strings"
	"testing"
)

func TestTabsToSpaces(t *testing.T) {
	in := []string{"hello\tworld"}
	expected := []string{"hello    world"}
	out := TabsToSpaces(in)
	if strings.Join(out, "") != strings.Join(expected, "") {
		t.Errorf("tabsToSpaces failed: got %v, want %v", out, expected)
	}
}

func TestCalculateMaxWidth(t *testing.T) {
	in := []string{"short", "longer line", "mid"}
	expected := len("longer line")
	out := CalculateMaxWidth(in)
	if out != expected {
		t.Errorf("calculateMaxWidth failed: got %d, want %d", out, expected)
	}
}

func TestNormalizeStringsLength(t *testing.T) {
	in := []string{"a", "bb", "ccc"}
	max := 3
	expected := []string{"a  ", "bb ", "ccc"}
	out := NormalizeStringsLength(in, max)
	for i := range expected {
		if out[i] != expected[i] {
			t.Errorf("normalizeStringsLength failed at line %d: got %q, want %q", i, out[i], expected[i])
		}
	}
}

func TestBuildBalloonSingleLine(t *testing.T) {
	lines := []string{"Hello, world!"}
	width := CalculateMaxWidth(lines)
	lines = NormalizeStringsLength(lines, width)
	out := BuildBalloon(lines, width)

	if !strings.Contains(out, "< Hello, world! >") {
		t.Errorf("Expected balloon to contain '< Hello, world! >', got:\n%s", out)
	}
}
