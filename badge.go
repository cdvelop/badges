package badges

import (
	"fmt"
	"strings"
)

// Badge represents a single badge definition (exported)
type Badge struct {
	Label string
	Value string
	Color string
}

// parseBadge parses a string in the format label:value:color (unexported)
// It treats special commands (output_svgfile/readmefile) as a sentinel error.
func parseBadge(s string) (Badge, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return Badge{}, fmt.Errorf("invalid badge format: %s", s)
	}
	if parts[0] == "output_svgfile" || parts[0] == "readmefile" {
		return Badge{}, fmt.Errorf("special command")
	}
	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return Badge{}, fmt.Errorf("empty fields in badge: %s", s)
	}
	return Badge{Label: parts[0], Value: parts[1], Color: parts[2]}, nil
}
