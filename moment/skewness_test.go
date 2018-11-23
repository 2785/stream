package moment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alexander-yu/stream"
	"github.com/alexander-yu/stream/testutil"
)

func TestSkewness(t *testing.T) {
	skewness, err := NewSkewness()
	require.NoError(t, err)

	stream.TestData(skewness)

	value, err := skewness.Value()
	require.NoError(t, err)

	adjust := 3.
	moment := 9.
	variance := 7.

	testutil.Approx(t, adjust*moment/math.Pow(variance, 1.5), value)
}