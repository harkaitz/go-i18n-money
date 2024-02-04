package i18nm

import (
	"fmt"
)

// Cents holds a monetary value in cents.
type Cents uint64

// Value returns a Cents value from a pair of integers. For example
// Money(1, 23) returns 123 cents.
func Value(a, b int) Cents {
	return Cents(a*100 + b)
}

// ParseMoney parses a monetary value string. It support english
// notation "1.23" "$1,000,000.00" and spanish notation "1,23" and
// "1.000.000,00". It uses the position of the comma or dot to
// distinguish.
//
// The string may also contain a leading minus. All other characters
// are ignored.
func ParseMoney(s string) (value Cents, negative bool, err error) {
	var s2 string
	var a, m, l int64 = 0, 0, -1
	var c rune
	for _, c = range s {
		switch {
		case c == '-' && len(s2) == 0:
			negative = true
		case c == '.' || c == ',' || c == '\'':
			l = 0
		case c >= '0' && c <= '9':
			if l != -1 {
				l++
			}
			s2 += string(c)
		}
	}
	switch l {
	case -1: m = 100
	case  1: m = 10
	case  2: m = 1
	case  3: m = 100
	default: return Cents(0), false, fmt.Errorf("Invalid monetary format")
	}
	_, err = fmt.Sscanf(s2, "%d", &a)
	if err != nil { return Cents(0), false, err }
	return Cents(a * int64(m)), negative, nil
}

func (c Cents) String() string {
	return fmt.Sprintf("%d.%02d", c / 100, c % 100)
}
