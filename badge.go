package badges

import (
	"errors"
	"fmt"
	"strings"
)

// badge represents a single badge definition (unexported)
type badge struct {
	label string
	value string
	color string
}

// parseBadge parses a string in the format label:value:color (unexported)
func parseBadge(s string) (badge, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return badge{}, fmt.Errorf("invalid badge format: %s", s)
	}
	if parts[0] == "output_svgfile" || parts[0] == "readmefile" {
		return badge{}, errors.New("special command")
	}
	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return badge{}, fmt.Errorf("empty fields in badge: %s", s)
	}
	return badge{label: parts[0], value: parts[1], color: parts[2]}, nil
}

// generateSVG generates a simple SVG for provided badges. It uses a naive width calculation. (unexported)
func generateSVG(badges []badge) (string, error) {
	if len(badges) == 0 {
		return "", fmt.Errorf("no badges to generate")
	}
	// simple generation similar to bash script
	svgParts := []string{"<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "<svg xmlns=\"http://www.w3.org/2000/svg\">"}
	for _, b := range badges {
		// minimal element
		svgParts = append(svgParts, fmt.Sprintf("<!-- Badge: %s -->", b.label))
		svgParts = append(svgParts, fmt.Sprintf("<text>%s: %s</text>", b.label, b.value))
	}
	svgParts = append(svgParts, "</svg>")
	return strings.Join(svgParts, "\n"), nil
}
