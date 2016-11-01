package strutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	// Strings accepted as boolean.
	boolString = map[string]bool{
		"t":     true,
		"true":  true,
		"y":     true,
		"yes":   true,
		"on":    true,
		"1":     true,
		"f":     false,
		"false": false,
		"n":     false,
		"no":    false,
		"off":   false,
		"0":     false,
	}
)

//没找到并且没有默认值则 panic抛错
func Stou64(v string, def ...uint64) uint64 {
	if n, err := strconv.ParseUint(v, 10, 64); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint64 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou32(v string, def ...uint32) uint32 {
	if n, err := strconv.ParseUint(v, 10, 64); err == nil {
		return uint32(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint32 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou16(v string, def ...uint16) uint16 {
	if n, err := strconv.ParseUint(v, 10, 64); err == nil {
		return uint16(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint16 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou8(v string, def ...uint8) uint8 {
	if n, err := strconv.ParseUint(v, 10, 64); err == nil {
		return uint8(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint8 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou(v string, def ...uint) uint {
	if n, err := strconv.ParseUint(v, 10, 64); err == nil {
		return uint(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi64(v string, def ...int64) int64 {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int64 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi32(v string, def ...int32) int32 {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return int32(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int32 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi16(v string, def ...int16) int16 {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return int16(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int16 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi8(v string, def ...int8) int8 {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return int8(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int8 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi(v string, def ...int) int {
	if n, err := strconv.ParseInt(v, 10, 64); err == nil {
		return int(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func StoBol(v string, def ...bool) bool {
	if value, ok := boolString[strings.ToLower(v)]; ok {
		return value
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid bool param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stof32(v string, def ...float32) float32 {
	if val, err := strconv.ParseFloat(v, 32); err == nil {
		return float32(val)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid float32 param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stof64(v string, def ...float64) float64 {
	if val, err := strconv.ParseFloat(v, 64); err == nil {
		return val
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid float64 param:%v", v))
	}
}

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func Median(items []float64) float64 {
	n := len(items)
	switch {
	case n == 0:
		return 0
	case n%2 == 1:
		return items[n/2]
	default:
		a := items[n/2-1]
		b := items[n/2]
		return (a + b) / 2
	}
}

func DurationString(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func ParseFloats(items []string) []float64 {
	result := make([]float64, len(items))
	for i, item := range items {
		f, _ := strconv.ParseFloat(item, 64)
		result[i] = f
	}
	return result
}

func ParseInts(items []string) []int {
	result := make([]int, len(items))
	for i, item := range items {
		f, _ := strconv.ParseInt(item, 0, 0)
		result[i] = int(f)
	}
	return result
}

func Fract(x float64) float64 {
	_, x = math.Modf(x)
	return x
}

func Clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}

func ClampInt(x, lo, hi int) int {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}

func ShowProgress(start time.Time, rays uint64, crtValue, totalValue int) {
	pct := int(100 * float64(crtValue) / float64(totalValue))
	elapsed := time.Since(start)
	fmt.Printf("\r%4d / %d (%3d%%) [", crtValue, totalValue, pct)
	for p := 0; p < 100; p += 3 {
		if pct > p {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %s ", DurationString(elapsed))
}
