package badges

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerate_MatchesExpected(t *testing.T) {
	// args designed to produce the expected.svg in testdata
	args := []string{
		"License:MIT:#007acc",
		"Go:1.24.4:#00ADD8",
		"Tests:Passing:#4c1",
		"Coverage:73%:#dfb317",
		"Race:Clean:#4c1",
		"Vet:OK:#4c1",
	}

	h := NewBadgeHandler(args...)
	if h == nil {
		t.Fatalf("NewBadgeHandler returned nil")
	}
	if err := h.Err(); err != nil {
		t.Fatalf("NewBadgeHandler error: %v", err)
	}

	svgBytes, _, err := h.GenerateSVG()
	if err != nil {
		t.Fatalf("GenerateSVG error: %v", err)
	}
	svg := string(svgBytes)

	wantPath := filepath.Join("testdata", "expected.svg")
	wantBytes, rerr := os.ReadFile(wantPath)
	if rerr != nil {
		t.Fatalf("read expected file: %v", rerr)
	}

	got := strings.TrimSpace(svg)
	want := strings.TrimSpace(string(wantBytes))

	if got != want {
		t.Fatalf("generated SVG differs from expected\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}
