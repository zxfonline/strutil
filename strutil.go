package strutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
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
	if n, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint64 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou32(v string, def ...uint32) uint32 {
	if n, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err == nil {
		return uint32(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint32 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou16(v string, def ...uint16) uint16 {
	if n, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err == nil {
		return uint16(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint16 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou8(v string, def ...uint8) uint8 {
	if n, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err == nil {
		return uint8(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint8 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stou(v string, def ...uint) uint {
	if n, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64); err == nil {
		return uint(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid uint param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi64(v string, def ...int64) int64 {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return n
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int64 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi32(v string, def ...int32) int32 {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return int32(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int32 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi16(v string, def ...int16) int16 {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return int16(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int16 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi8(v string, def ...int8) int8 {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return int8(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int8 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stoi(v string, def ...int) int {
	if n, err := strconv.ParseInt(strings.TrimSpace(v), 0, 0); err == nil {
		return int(n)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid int param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func StoBol(v string, def ...bool) bool {
	if value, ok := boolString[strings.ToLower(strings.TrimSpace(v))]; ok {
		return value
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid bool param:%v", v))
	}
}

//没找到并且没有默认值则 panic抛错
func Stof32(v string, def ...float32) float32 {
	if val, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
		return float32(val)
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid float32 param:%v ,err:%v", v, err))
	}
}

//没找到并且没有默认值则 panic抛错
func Stof64(v string, def ...float64) float64 {
	if val, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
		return val
	} else if len(def) > 0 {
		return def[0]
	} else {
		panic(fmt.Errorf("invalid float64 param:%v ,err:%v", v, err))
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

//忽略空字符串
func ParseFloat64s(items []string) ([]float64, error) {
	result := make([]float64, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				result = append(result, f)
			} else {
				return nil, fmt.Errorf("invalid []float64 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseFloat32s(items []string) ([]float32, error) {
	result := make([]float32, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				result = append(result, float32(f))
			} else {
				return nil, fmt.Errorf("invalid []float32 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseInts(items []string) ([]int, error) {
	result := make([]int, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseInt(t, 0, 0); err == nil {
				result = append(result, int(f))
			} else {
				return nil, fmt.Errorf("invalid []int param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseInt64s(items []string) ([]int64, error) {
	result := make([]int64, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseInt(t, 0, 0); err == nil {
				result = append(result, f)
			} else {
				return nil, fmt.Errorf("invalid []int64 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseInt32s(items []string) ([]int32, error) {
	result := make([]int32, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseInt(t, 0, 0); err == nil {
				result = append(result, int32(f))
			} else {
				return nil, fmt.Errorf("invalid []int32 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseInt16s(items []string) ([]int16, error) {
	result := make([]int16, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseInt(t, 0, 0); err == nil {
				result = append(result, int16(f))
			} else {
				return nil, fmt.Errorf("invalid []int16 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseInt8s(items []string) ([]int8, error) {
	result := make([]int8, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, err := strconv.ParseInt(t, 0, 0); err == nil {
				result = append(result, int8(f))
			} else {
				return nil, fmt.Errorf("invalid []int8 param:%v ,err:%v", items, err)
			}
		}
	}
	return result, nil
}

//忽略空字符串
func ParseBools(items []string) ([]bool, error) {
	result := make([]bool, 0, len(items))
	for i := 0; i < len(items); i++ {
		t := strings.TrimSpace(items[i])
		if len(t) != 0 {
			if f, ok := boolString[strings.ToLower(t)]; ok {
				result = append(result, f)
			} else {
				return nil, fmt.Errorf("invalid []bool param:%v", items)
			}
		}
	}
	return result, nil
}

func Float64sToStrs(items []float64) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func Float32sToStrs(items []float32) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func IntsToStrs(items []int) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func Int64sToStrs(items []int64) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func Int32sToStrs(items []int32) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func Int16sToStrs(items []int16) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func Int8sToStrs(items []int8) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
	}
	return result
}

func BoolsToStrs(items []bool) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%v", item)
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

func FirstIsLower(s string) bool {
	if r := s[0]; r < utf8.RuneSelf {
		return 'a' <= r && r <= 'z'
	}
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(r)
}

func FirstIsUpper(s string) bool {
	if r := s[0]; r < utf8.RuneSelf {
		return 'A' <= r && r <= 'Z'
	}
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func FirstToUpper(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	if r == '_' {
		return s
	}
	return string(unicode.ToUpper(r)) + s[n:]
}

func FirstToLower(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	if r == '_' {
		return s
	}
	return string(unicode.ToLower(r)) + s[n:]
}
