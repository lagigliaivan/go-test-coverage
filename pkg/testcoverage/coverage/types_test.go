package coverage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lagigliaivan/go-test-coverage/pkg/testcoverage/coverage"
)

func TestCoveredPercentage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		percentage float64
		total      int64
		covered    int64
	}{
		{percentage: 0, total: 0, covered: 0},
		{percentage: 0, total: 0, covered: 1},
		{percentage: 100, total: 1, covered: 1},
		{percentage: 10, total: 10, covered: 1},
		{percentage: 22.22, total: 9, covered: 2}, // 22.222.. should round down to 22
		{percentage: 66.67, total: 9, covered: 6}, // 66.666.. should round down to 66
	}

	for _, tc := range tests {
		assert.Equal(t, fmt.Sprintf("%.2f", tc.percentage), fmt.Sprintf("%.2f", CoveredPercentage(tc.total, tc.covered)))
	}
}
