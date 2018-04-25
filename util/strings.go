package cloudimageeditor

import (
	"fmt"
)

/*
 * cloudimageeditor package
 *
 * Authors:
 *     zhenwei pi <cloudimageeditor@126.com>
 *
 * This project is under Apache v2 License.
 *
 */

func isAlpha(c byte) bool {
	if ((c >= 'a') && (c <= 'z')) || ((c >= 'A') && (c <= 'Z')) {
		return true
	}

	return false
}

func isNum(c byte) bool {
	if (c >= '0') && (c <= '9') {
		return true
	}

	return false
}

// StringKeepAlpha : only keep alphabetic character
func StringKeepAlpha(src string) (dest string) {
	for index := 0; index < len(src); index++ {
		b := src[index]
		if isAlpha(b) {
			dest = fmt.Sprintf("%s%c", dest, b)
		}
	}

	return dest
}

// StringKeepNumber : only keep alphanumeric character
func StringKeepNumber(src string) (dest string) {
	for index := 0; index < len(src); index++ {
		b := src[index]
		if isNum(b) {
			dest = fmt.Sprintf("%s%c", dest, b)
		}
	}

	return dest
}

// StringKeepFloatNumber : only keep alphanumeric character and '.'
func StringKeepFloatNumber(src string) (dest string) {
	for index := 0; index < len(src); index++ {
		b := src[index]
		if isNum(b) || (b == '.') {
			dest = fmt.Sprintf("%s%c", dest, b)
		}
	}

	return dest
}

// StringKeepAlphaAndNum : only keep alphabetic character and alphanumeric character
func StringKeepAlphaAndNum(src string) (dest string) {
	for index := 0; index < len(src); index++ {
		b := src[index]
		if isAlpha(b) || isNum(b) {
			dest = fmt.Sprintf("%s%c", dest, b)
		}
	}

	return dest
}

// ParseSizeWithUnit : parse size with unit k/m/g/t
func ParseSizeWithUnit(size string, unit string) (value uint64, err error) {
	var fValue float64
	fmt.Sscanf(size, "%f", &fValue)
	if len(unit) == 0 {
		return uint64(fValue), nil
	}

	kb := float64(1024)
	mb := 1024 * kb
	gb := 1024 * mb
	tb := 1024 * gb
	switch unit {
	case "K":
		fallthrough
	case "k":
		fValue *= kb

	case "M":
		fallthrough
	case "m":
		fValue *= mb

	case "G":
		fallthrough
	case "g":
		fValue *= gb

	case "T":
		fallthrough
	case "t":
		fValue *= tb
	default:
		return 0, fmt.Errorf("Unrecongized unit %v", unit)
	}

	return uint64(fValue), nil
}
