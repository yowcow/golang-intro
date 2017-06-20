package hello

import (
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	re := regexp.MustCompile("\\A\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}\\z")
	now := time.Now().Format("2006-01-02 15:04:05")
	assert.True(t, re.MatchString(now))
}

func TestTimezone(t *testing.T) {
	assert := assert.New(t)

	jploc, e := time.LoadLocation("Asia/Tokyo")

	assert.Nil(e)

	tm := time.Date(2000, time.January, 1, 0, 0, 0, 0, jploc)

	assert.Equal("2000-01-01T00:00:00+09:00", tm.Format(time.RFC3339))
	assert.Equal("Asia/Tokyo", tm.Location().String())

	utcloc, _ := time.LoadLocation("UTC")
	cnloc, _ := time.LoadLocation("Asia/Shanghai")

	assert.Equal("1999-12-31T15:00:00Z", tm.In(utcloc).Format(time.RFC3339))
	assert.Equal("1999-12-31T23:00:00+08:00", tm.In(cnloc).Format(time.RFC3339))
}

func TestAddDuration(t *testing.T) {
	assert := assert.New(t)

	tm1, _ := time.Parse(time.RFC3339, "2000-01-01T00:00:00+09:00")
	tm2 := tm1.Add(10 * time.Minute)

	assert.Equal("2000-01-01T00:00:00+09:00", tm1.Format(time.RFC3339))
	assert.Equal("2000-01-01T00:10:00+09:00", tm2.Format(time.RFC3339))
}

func TestAfter(t *testing.T) {
	assert := assert.New(t)

	wg := &sync.WaitGroup{}
	ok := false

	wg.Add(1)
	go func(o *bool) {
		defer wg.Done()
		select {
		case <-time.After(100 * time.Millisecond):
			*o = true
			return
		}
	}(&ok)

	wg.Wait()

	assert.True(ok)
}
