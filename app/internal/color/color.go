package color

import "fmt"

var (
	Red    = Color("\033[1;31m%s\033[0m")
	Green  = Color("\033[1;32m%s\033[0m")
	Yellow = Color("\033[1;33m%s\033[0m")
)

func Color(colorStr string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(colorStr, fmt.Sprint(args...))
	}
}
