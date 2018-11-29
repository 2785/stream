package median

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexander-yu/stream/testutil"
)

func TestNewAVLMedian(t *testing.T) {
	t.Run("pass: nonnegative window is valid", func(t *testing.T) {
		median, err := NewAVLMedian(0)
		require.NoError(t, err)
		assert.Equal(t, 0, median.window)

		median, err = NewAVLMedian(5)
		require.NoError(t, err)
		assert.Equal(t, 5, median.window)
	})

	t.Run("fail: negative window is invalid", func(t *testing.T) {
		_, err := NewAVLMedian(-1)
		assert.EqualError(t, err, "-1 is a negative window")
	})
}

func TestAVLMedianPush(t *testing.T) {
	median, err := NewAVLMedian(3)
	require.NoError(t, err)
	for i := 0.; i < 5; i++ {
		err := median.Push(i)
		require.NoError(t, err)
	}

	assert.Equal(t, uint64(3), median.queue.Len())
	for i := 2.; i < 5; i++ {
		val, err := median.queue.Get()
		y := val.(float64)
		require.NoError(t, err)
		testutil.Approx(t, i, y)
	}

	assert.Equal(
		t,
		strings.Join([]string{
			"│   ┌── 4.000000",
			"└── 3.000000",
			"    └── 2.000000",
			"",
		}, "\n"),
		median.tree.String(),
	)
}

func TestAVLMedianValue(t *testing.T) {
	t.Run("pass: if number of values is even, return average of middle two", func(t *testing.T) {
		median, err := NewAVLMedian(4)
		require.NoError(t, err)
		for i := 0.; i < 6; i++ {
			err := median.Push(i)
			require.NoError(t, err)
		}

		value, err := median.Value()
		require.NoError(t, err)

		assert.Equal(t, 3.5, value)
	})

	t.Run("pass: if number of values is odd, return middle value", func(t *testing.T) {
		median, err := NewAVLMedian(3)
		require.NoError(t, err)
		for i := 0.; i < 5; i++ {
			err := median.Push(i)
			require.NoError(t, err)
		}

		value, err := median.Value()
		require.NoError(t, err)

		assert.Equal(t, float64(3), value)
	})
}
