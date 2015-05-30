package timeutil

import (
	"fmt"
	"strconv"
	"time"
)

func IsGreaterThan(timeLeft time.Time, timeRight time.Time) bool {
	durDelta := timeLeft.Sub(timeRight)
	if durZero, _ := time.ParseDuration("0ns"); durDelta > durZero {
		return true
	}
	return false
}

func IsLessThan(timeLeft time.Time, timeRight time.Time) bool {
	durDelta := timeLeft.Sub(timeRight)
	if durZero, _ := time.ParseDuration("0ns"); durDelta < durZero {
		return true
	}
	return false
}

func Dt8NowUtc() int32 {
	tm := time.Now()
	tm = tm.UTC()
	return Dt8ForInts(tm.Year(), int(tm.Month()), tm.Day())
}

func Dt8ForInts(yyyy int, mm int, dd int) int32 {
	sDt8 := fmt.Sprintf("%04d%02d%02d", yyyy, mm, dd)
	iDt8, _ := strconv.ParseInt(sDt8, 10, 32)
	return int32(iDt8)
}

func DurationForNowSubDt8(dt8 int32) (time.Duration, error) {
	t, err := TimeForDt8(dt8)
	if err != nil {
		var d time.Duration
		return d, err
	}
	now := time.Now()
	return now.Sub(t), nil
}

func TimeForDt8(dt8 int32) (time.Time, error) {
	return time.Parse("20060102", strconv.FormatInt(int64(dt8), 10))
}

func Dt14NowUtc() int64 {
	tm := time.Now()
	tm = tm.UTC()
	return Dt14ForInts(tm.Year(), int(tm.Month()), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

func Dt14ForInts(yyyy int, mm int, dd int, hr int, mn int, dy int) int64 {
	sDt14 := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", yyyy, mm, dd, hr, mn, dy)
	iDt14, _ := strconv.ParseInt(sDt14, 10, 64)
	return int64(iDt14)
}
