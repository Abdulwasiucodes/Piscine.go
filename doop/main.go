package main

import (
	"os"
)

const (
	MaxInt64 = 1<<63 - 1
	MinInt64 = -1 << 63
)

// parseInt64 parses base-10 int64 from s.
// returns (value, ok). ok==false on any invalid input or overflow.
func parseInt64(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	i := 0
	neg := false
	switch s[0] {
	case '+':
		i = 1
	case '-':
		neg = true
		i = 1
	}
	if i >= len(s) {
		return 0, false
	}

	if !neg {
		var acc int64 = 0
		for ; i < len(s); i++ {
			ch := s[i]
			if ch < '0' || ch > '9' {
				return 0, false
			}
			d := int64(ch - '0')
			// check acc > (MaxInt64 - d)/10
			if acc > (MaxInt64-d)/10 {
				return 0, false
			}
			acc = acc*10 + d
		}
		return acc, true
	} else {
		var acc int64 = 0 // accumulate negative: acc = acc*10 - d
		for ; i < len(s); i++ {
			ch := s[i]
			if ch < '0' || ch > '9' {
				return 0, false
			}
			d := int64(ch - '0')
			// check acc < (MinInt64 + d)/10  (would underflow)
			if acc < (MinInt64+d)/10 {
				return 0, false
			}
			acc = acc*10 - d
		}
		return acc, true
	}
}

func itoaInt64(n int64) string {
	if n == 0 {
		return "0"
	}
	// handle MinInt64 explicitly
	if n == MinInt64 {
		return "-9223372036854775808"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	// produce digits reversed
	buf := make([]byte, 0, 20)
	for n > 0 {
		d := byte(n%10) + '0'
		buf = append(buf, d)
		n /= 10
	}
	if neg {
		buf = append(buf, '-')
	}
	// reverse
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func write(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, okA := parseInt64(aStr)
	b, okB := parseInt64(bStr)
	if !okA || !okB {
		return
	}

	switch op {
	case "+":
		// overflow if (b>0 && a > Max-b) or (b<0 && a < Min-b)
		if b > 0 && a > MaxInt64-b {
			return
		}
		if b < 0 && a < MinInt64-b {
			return
		}
		write(itoaInt64(a+b) + "\n")
	case "-":
		// a - b overflow checks:
		// if b > 0 and a < Min + b -> underflow
		// if b < 0 and a > Max + b -> overflow
		if b > 0 {
			if a < MinInt64+b {
				return
			}
		} else {
			if a > MaxInt64+b {
				return
			}
		}
		write(itoaInt64(a-b) + "\n")
	case "*":
		// handle simple cases
		if a == 0 || b == 0 {
			write("0\n")
			return
		}
		// special overflow case: MinInt64 * -1 (but we handle via unsigned check)
		// use absolute values in unsigned to check: ua > Max/ub -> overflow
		var ua, ub uint64
		if a < 0 {
			if a == MinInt64 {
				ua = 1 << 63
			} else {
				ua = uint64(-a)
			}
		} else {
			ua = uint64(a)
		}
		if b < 0 {
			if b == MinInt64 {
				ub = 1 << 63
			} else {
				ub = uint64(-b)
			}
		} else {
			ub = uint64(b)
		}
		if ub != 0 && ua > uint64(MaxInt64)/ub {
			return
		}
		write(itoaInt64(a*b) + "\n")
	case "/":
		if b == 0 {
			write("No division by 0\n")
			return
		}
		// overflow: MinInt64 / -1
		if a == MinInt64 && b == -1 {
			return
		}
		write(itoaInt64(a/b) + "\n")
	case "%":
		if b == 0 {
			write("No modulo by 0\n")
			return
		}
		write(itoaInt64(a%b) + "\n")
	default:
		return
	}
}
