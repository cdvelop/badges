package gitutils

import (
	"os/exec"
	"regexp"
	"strings"
)

// GetGitHostUserPath returns the git host and user/organization path from the remote origin URL.
// For example, "github.com/cdvelop"
func GetGitHostUserPath() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	url := strings.TrimSpace(string(output))
	re := regexp.MustCompile(`(?:git@|https://)([^:/]+)[/:]([^/]+)/.*`)
	matches := re.FindStringSubmatch(url)
	if len(matches) > 2 {
		return matches[1] + "/" + matches[2], nil
	}

	return "", nil
}
