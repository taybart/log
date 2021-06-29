package log

import "fmt"

var TmplFuncs = map[string]interface{}{
	"red":    red,
	"gray":   gray,
	"purple": purple,
	"blue":   blue,
	"yellow": yellow,
	"green":  green,
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", Gray, s, Reset)
}

func purple(s string) string {
	return fmt.Sprintf("%s%s%s", Purple, s, Reset)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}

func yellow(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Reset)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Reset)
}
