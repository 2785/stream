package minmax

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testutil "github.com/alexander-yu/stream/util/test"
)

func TestNewMin(t *testing.T) {
	t.Run("pass: valid Min is valid", func(t *testing.T) {
		min, err := NewMin(3)
		require.NoError(t, err)
		assert.Equal(t, 3, min.window)
		assert.Equal(t, uint64(0), min.queue.Len())
		assert.Equal(t, 0, min.deque.Len())
		assert.Equal(t, math.Inf(1), min.min)
		assert.Equal(t, 0, min.count)
	})

	t.Run("fail: negative window returns error", func(t *testing.T) {
		_, err := NewMin(-1)
		testutil.ContainsError(t, err, "-1 is a negative window")
	})
}

func TestMinString(t *testing.T) {
	expectedString := "minmax.Min_{window:3}"
	min, err := NewMin(3)
	require.NoError(t, err)
	assert.Equal(t, expectedString, min.String())
}

func TestMinValue(t *testing.T) {
	t.Run("pass: returns global minimum for a window of 0", func(t *testing.T) {
		min, err := NewMin(0)
		require.NoError(t, err)

		vals := []float64{9, 4, 6, 1, 8, 2, 2, 5, 5, 3}
		for _, val := range vals {
			err := min.Push(val)
			require.NoError(t, err)
		}

		val, err := min.Value()
		require.NoError(t, err)
		testutil.Approx(t, 1., val)
	})

	t.Run("pass: returns minimum for a provided window", func(t *testing.T) {
		min, err := NewMin(5)
		require.NoError(t, err)

		vals := []float64{9, 4, 6, 1, 8, 2, 2, 5, 5, 3}
		for _, val := range vals {
			err = min.Push(val)
			require.NoError(t, err)
		}

		val, err := min.Value()
		require.NoError(t, err)
		testutil.Approx(t, 2., val)
	})

	t.Run("fail: if no values seen, return error", func(t *testing.T) {
		min, err := NewMin(3)
		require.NoError(t, err)

		_, err = min.Value()
		assert.EqualError(t, err, "no values seen yet")
	})

	t.Run("fail: if queue retrieval fails, return error", func(t *testing.T) {
		min, err := NewMin(3)
		require.NoError(t, err)

		for i := 0.; i < 3; i++ {
			err = min.Push(i)
			require.NoError(t, err)
		}

		// dispose the queue to simulate an error when we try to retrieve from the queue
		min.queue.Dispose()
		err = min.Push(3.)
		testutil.ContainsError(t, err, "error popping item from queue")
	})

	t.Run("fail: if queue insertion fails, return error", func(t *testing.T) {
		min, err := NewMin(3)
		require.NoError(t, err)

		// dispose the queue to simulate an error when we try to insert into the queue
		min.queue.Dispose()
		val := 3.
		err = min.Push(val)
		testutil.ContainsError(t, err, fmt.Sprintf("error pushing %f to queue", val))
	})
}

func TestMinClear(t *testing.T) {
	min, err := NewMin(3)
	require.NoError(t, err)

	for i := 0.; i < 3; i++ {
		err = min.Push(i)
		require.NoError(t, err)
	}

	min.Clear()
	assert.Equal(t, 0, min.count)
	assert.Equal(t, math.Inf(1), min.min)
	assert.Equal(t, uint64(0), min.queue.Len())
	assert.Equal(t, 0, min.deque.Len())
}
