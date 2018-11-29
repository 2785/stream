package moment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexander-yu/stream"
	"github.com/alexander-yu/stream/testutil"
)

func TestNewMoment(t *testing.T) {
	t.Run("pass: nonnegative moment is valid", func(t *testing.T) {
		moment, err := NewMoment(0)
		require.NoError(t, err)
		assert.Equal(t, 0, moment.k)

		moment, err = NewMoment(5)
		require.NoError(t, err)
		assert.Equal(t, 5, moment.k)
	})

	t.Run("fail: negative moment is invalid", func(t *testing.T) {
		_, err := NewMoment(-1)
		assert.EqualError(t, err, "-1 is a negative moment")
	})
}

func TestMoment(t *testing.T) {
	t.Run("pass: returns the kth moment", func(t *testing.T) {
		moment, err := NewMoment(2)
		require.NoError(t, err)

		stream.TestData(moment)

		value, err := moment.Value()
		require.NoError(t, err)

		testutil.Approx(t, 7, value)
	})

	t.Run("pass: 0th moment always returns 1", func(t *testing.T) {
		moment, err := NewMoment(0)
		require.NoError(t, err)

		stream.TestData(moment)

		value, err := moment.Value()
		require.NoError(t, err)

		testutil.Approx(t, 1, value)

		err = moment.core.Push(10)
		require.NoError(t, err)

		value, err = moment.Value()
		require.NoError(t, err)

		testutil.Approx(t, 1, value)
	})

	t.Run("pass: 1st moment always returns 0", func(t *testing.T) {
		moment, err := NewMoment(1)
		require.NoError(t, err)

		stream.TestData(moment)

		value, err := moment.Value()
		require.NoError(t, err)

		testutil.Approx(t, 0, value)

		err = moment.core.Push(10)
		require.NoError(t, err)

		value, err = moment.Value()
		require.NoError(t, err)

		testutil.Approx(t, 0, value)
	})

	t.Run("fail: error if no values are seen", func(t *testing.T) {
		moment, err := NewMoment(2)
		require.NoError(t, err)

		stream.NewCore(&stream.CoreConfig{}, moment)

		_, err = moment.Value()
		assert.EqualError(t, err, "no values seen yet")
	})
}
