package path

import (
	"path/filepath"
	"runtime"
	"strings"
)

const separatorToReplace = string(filepath.Separator)

func NormalizePathInRegex(path string) string {
	if runtime.GOOS != "windows" {
		return path
	}

	return strings.ReplaceAll(path, "/", separatorToReplace)
}
