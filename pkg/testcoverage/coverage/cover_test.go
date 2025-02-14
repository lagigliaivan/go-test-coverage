package coverage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lagigliaivan/go-test-coverage/pkg/testcoverage/coverage"
	"github.com/lagigliaivan/go-test-coverage/pkg/testcoverage/testdata"
)

const (
	profileOK  = "../testdata/" + testdata.ProfileOK
	profileNOK = "../testdata/" + testdata.ProfileNOK
	prefix     = "github.com/vladopajic/go-test-coverage/v2"
)

func Test_GenerateCoverageStats(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		return
	}

	// should not be able to read directory
	stats, err := GenerateCoverageStats(Config{Profile: t.TempDir()})
	assert.Error(t, err)
	assert.Empty(t, stats)

	// should get error parsing invalid profile file
	stats, err = GenerateCoverageStats(Config{Profile: profileNOK})
	assert.Error(t, err)
	assert.Empty(t, stats)

	// should be okay to read valid profile
	stats1, err := GenerateCoverageStats(Config{Profile: profileOK})
	assert.NoError(t, err)
	assert.NotEmpty(t, stats1)

	// should be okay to read valid profile
	stats2, err := GenerateCoverageStats(Config{
		Profile:      profileOK,
		ExcludePaths: []string{`cover\.go$`},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, stats2)
	// stats2 should have less total statements because cover.go should have be excluded
	assert.Greater(t, CalcTotalStats(stats1).Total, CalcTotalStats(stats2).Total)

	// should remove prefix from stats
	stats3, err := GenerateCoverageStats(Config{
		Profile:     profileOK,
		LocalPrefix: prefix,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, stats3)
	assert.Equal(t, CalcTotalStats(stats1), CalcTotalStats(stats3))
	assert.NotContains(t, stats3[0].Name, prefix)
}
