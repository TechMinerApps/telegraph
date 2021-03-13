package telegraph_test

import (
	"testing"
	"time"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetViews(t *testing.T) {
	c := telegraph.NewClient()
	t.Run("Invalid", func(t *testing.T) {
		t.Run("Path", func(t *testing.T) {
			_, err := c.GetViews("wtf", time.Time{})
			assert.Error(t, err)
		})
		t.Run("Year", func(t *testing.T) {
			dt := time.Date(1980, 0, 0, 0, 0, 0, 0, time.UTC)
			_, err := c.GetViews("Sample-Page-12-15", dt)
			assert.Error(t, err)
		})
		t.Run("Month", func(t *testing.T) {
			dt := time.Date(2000, 22, 0, 0, 0, 0, 0, time.UTC)
			result, err := c.GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
		t.Run("Day", func(t *testing.T) {
			dt := time.Date(2000, time.February, 42, 0, 0, 0, 0, time.UTC)
			result, err := c.GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
		t.Run("Hour", func(t *testing.T) {
			dt := time.Date(2000, time.February, 12, 65, 0, 0, 0, time.UTC)
			result, err := c.GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	})
	t.Run("Valid", func(t *testing.T) {
		dt := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)
		stats, err := c.GetViews("Sample-Page-12-15", dt)
		assert.NoError(t, err)
		if !assert.NotNil(t, stats) {
			t.FailNow()
		}

		assert.NotZero(t, stats.Views)
	})
}
