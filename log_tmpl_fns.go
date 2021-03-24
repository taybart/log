package log

import "fmt"

var Funcs = map[string]interface{}{
	"red":    red,
	"gray":   gray,
	"purple": purple,
	"blue":   blue,
	"yellow": yellow,
	"green":  green,
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Rtd)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", Gray, s, Rtd)
}

func purple(s string) string {
	return fmt.Sprintf("%s%s%s", Purple, s, Rtd)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Rtd)
}

func yellow(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Rtd)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Rtd)
}
