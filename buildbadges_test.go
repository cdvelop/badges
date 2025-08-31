package badges

import (
	"os"
	"path/filepath"
	"testing"
)

// TestBuildBadges_SectionArgsOrder ensures BuildBadges returns the four
// arguments required by SectionUpdate in the exact order:
// [sectionID, afterLine, content, readmeFile]
func TestBuildBadges_SectionArgsOrder(t *testing.T) {
	tmp := t.TempDir()
	out := filepath.Join(tmp, "badges.svg")

	// create handler with explicit output and readme file settings and one badge
	h := NewBadgeHandler("output_svgfile:"+out, "readmefile:MYREADME.md", "License:MIT:#007acc")
	if h == nil {
		t.Fatal("NewBadgeHandler returned nil")
	}
	if err := h.Err(); err != nil {
		t.Fatalf("handler init error: %v", err)
	}

	args, err := h.BuildBadges()
	if err != nil {
		t.Fatalf("BuildBadges error: %v", err)
	}

	if len(args) != 4 {
		t.Fatalf("expected 4 args, got %d: %#v", len(args), args)
	}

	// Check order
	if args[0] != "BADGES_SECTION" {
		t.Errorf("sectionID order mismatch: want BADGES_SECTION, got %q", args[0])
	}
	if args[1] != "1" {
		t.Errorf("afterLine order mismatch: want \"1\", got %q", args[1])
	}
	if args[2] != h.BadgeMarkdown() {
		t.Errorf("content mismatch: want BadgeMarkdown(), got %q", args[2])
	}
	if args[3] != "MYREADME.md" {
		t.Errorf("readmeFile mismatch: want MYREADME.md, got %q", args[3])
	}

	// Ensure output file was written and looks like an SVG
	info, err := os.Stat(out)
	if err != nil {
		t.Fatalf("expected output file %s to exist: %v", out, err)
	}
	if info.Size() == 0 {
		t.Fatalf("output file %s is empty", out)
	}
}
