package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// FormatIDR returns formatted Rupiah string like "Rp 1.234.567".
func FormatIDR(amount int64) string {
	return FormatCurrency(float64(amount), "Rp", 0)
}

// FormatCurrency formats value using thousands separator and decimals.
func FormatCurrency(amount float64, symbol string, decimals int) string {
	rounded := math.Round(amount*math.Pow10(decimals)) / math.Pow10(decimals)
	parts := strings.Split(fmt.Sprintf("%.*f", decimals, math.Abs(rounded)), ".")
	intPart := parts[0]
	fracPart := ""
	if decimals > 0 && len(parts) > 1 {
		fracPart = parts[1]
	}
	if len(intPart) > 3 {
		n := len(intPart)
		var groups []string
		for n > 3 {
			groups = append([]string{intPart[n-3:]}, groups...)
			intPart = intPart[:n-3]
			n = len(intPart)
		}
		if intPart != "" {
			groups = append([]string{intPart}, groups...)
		}
		intPart = strings.Join(groups, ".")
	}
	if decimals > 0 {
		return fmt.Sprintf("%s %s,%s", symbol, intPart, fracPart)
	}
	return fmt.Sprintf("%s %s", symbol, intPart)
}

// ParseCurrency parses amount string to float64. Accepts ".", "," and no separator.
func ParseCurrency(value string) (float64, error) {
	if value == "" {
		return 0, fmt.Errorf("empty value")
	}
	clean := strings.TrimSpace(value)
	clean = strings.ReplaceAll(clean, "Rp", "")
	clean = strings.ReplaceAll(clean, " ", "")
	clean = strings.ReplaceAll(clean, ".", "")
	clean = strings.ReplaceAll(clean, ",", ".")
	return strconv.ParseFloat(clean, 64)
}

// ConvertCurrency converts amount via rate.
func ConvertCurrency(amount float64, rate float64) float64 {
	return amount * rate
}

// RoundCurrency rounds to n decimals.
func RoundCurrency(amount float64, decimals int) float64 {
	factor := math.Pow10(decimals)
	return math.Round(amount*factor) / factor
}
