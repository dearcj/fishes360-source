package main

import (
	"math"
	"strconv"
	"strings"
)

var finfp = math.Inf(1)
var finfn = math.Inf(-1)

func InterfaceMapToMapID(m map[string]interface{}) map[MapID]string {
	var nw map[MapID]string
	nw = make(map[MapID]string)
	for a, b := range m {
		i := b.(MapID)
		nw[i] = a

	}

	return nw
}

func AddLowerCaseMap(m map[string]interface{}) map[string]interface{} {
	for a, b := range m {
		m[strings.ToLower(a)] = b
	}

	return m
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func boolToStr(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func strFloat32(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', 6, 32)
}

func iMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func iMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
