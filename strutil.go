package strutil

import (
	"fmt"
	"strconv"
	"strings"
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
