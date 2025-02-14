package testcoverage

import (
	"strings"

	"golang.org/x/exp/maps"

	"github.com/lagigliaivan/go-test-coverage/pkg/testcoverage/coverage"
)

type AnalyzeResult struct {
	Threshold              Threshold
	FilesBelowThreshold    []coverage.Stats
	PackagesBelowThreshold []coverage.Stats
	MeetsTotalCoverage     bool
	TotalCoverage          float64
}

func (r *AnalyzeResult) Pass() bool {
	return r.MeetsTotalCoverage &&
		len(r.FilesBelowThreshold) == 0 &&
		len(r.PackagesBelowThreshold) == 0
}

func packageForFile(filename string) string {
	i := strings.LastIndex(filename, "/")
	if i == -1 {
		return filename
	}

	return filename[:i]
}

func checkCoverageStatsBelowThreshold(
	coverageStats []coverage.Stats,
	threshold int,
	overrideRules []regRule,
) []coverage.Stats {
	var belowThreshold []coverage.Stats

	for _, s := range coverageStats {
		thr := threshold
		if override, ok := matches(overrideRules, s.Name); ok {
			thr = override
		}

		if s.CoveredPercentage() < float64(thr) {
			s.Threshold = thr
			belowThreshold = append(belowThreshold, s)
		}
	}

	return belowThreshold
}

func makePackageStats(coverageStats []coverage.Stats) []coverage.Stats {
	packageStats := make(map[string]coverage.Stats)

	for _, stats := range coverageStats {
		pkg := packageForFile(stats.Name)

		var pkgStats coverage.Stats
		if s, ok := packageStats[pkg]; ok {
			pkgStats = s
		} else {
			pkgStats = coverage.Stats{Name: pkg}
		}

		pkgStats.Total += stats.Total
		pkgStats.Covered += stats.Covered
		packageStats[pkg] = pkgStats
	}

	return maps.Values(packageStats)
}
