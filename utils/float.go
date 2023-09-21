package utils

import (
	"fmt"
	"math"
	"strconv"
)

// ReservedOneDecimal 保留一位小数
func ReservedOneDecimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return value
}

// ReservedTwoDecimal 保留两位小数
func ReservedTwoDecimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}
