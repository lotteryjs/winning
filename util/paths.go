package util

import (
	"strings"
)

// Path prefixes.
const (
	PathRoot     = "/"
	PathPlatInfo = "/plat/info"
)

var reservedPaths = []string{PathPlatInfo}

// IsReservedPath checks the specified path is a reserved path or not.
func IsReservedPath(path string) bool {
	path = strings.TrimSpace(path)
	if PathRoot == path {
		return true
	}

	for _, reservedPath := range reservedPaths {
		if strings.HasPrefix(path, reservedPath) {
			return true
		}
	}

	return false
}
