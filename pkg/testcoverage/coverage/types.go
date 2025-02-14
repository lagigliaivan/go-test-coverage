package coverage

import (
	"regexp"
	"strings"

	"github.com/lagigliaivan/go-test-coverage/pkg/testcoverage/path"
)

type Stats struct {
	Name      string
	Total     int64
	Covered   int64
	Threshold int
}

func (s Stats) CoveredPercentage() float64 {
	return CoveredPercentage(s.Total, s.Covered)
}

//nolint:gomnd // relax
func CoveredPercentage(total, covered int64) float64 {
	if total == 0 {
		return 0
	}

	if covered == total {
		return 100
	}

	return float64(covered*100) / float64(total)
}

func stripPrefix(name, prefix string) string {
	if prefix == "" {
		return name
	}

	if string(prefix[len(prefix)-1]) != "/" {
		prefix += "/"
	}

	return strings.Replace(name, prefix, "", 1)
}

func matches(regexps []*regexp.Regexp, str string) bool {
	for _, r := range regexps {
		if r.MatchString(str) {
			return true
		}
	}

	return false
}

func compileExcludePathRules(excludePaths []string) []*regexp.Regexp {
	if len(excludePaths) == 0 {
		return nil
	}

	compiled := make([]*regexp.Regexp, len(excludePaths))

	for i, pattern := range excludePaths {
		pattern = path.NormalizePathInRegex(pattern)
		compiled[i] = regexp.MustCompile(pattern)
	}

	return compiled
}

func CalcTotalStats(coverageStats []Stats) Stats {
	totalStats := Stats{}

	for _, stats := range coverageStats {
		totalStats.Total += stats.Total
		totalStats.Covered += stats.Covered
	}

	return totalStats
}
