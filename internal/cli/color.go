package cli

import (
	"os"
	"strings"
)

type Color struct {
	enabled bool
}

func NewColor(enabled bool) Color {
	return Color{enabled: enabled}
}

func (c Color) Bold(value string) string {
	return c.wrap("1", value)
}

func (c Color) Cyan(value string) string {
	return c.wrap("36", value)
}

func (c Color) Green(value string) string {
	return c.wrap("32", value)
}

func (c Color) Yellow(value string) string {
	return c.wrap("33", value)
}

func (c Color) Dim(value string) string {
	return c.wrap("2", value)
}

func (c Color) wrap(code string, value string) string {
	if !c.enabled {
		return value
	}
	return "\x1b[" + code + "m" + value + "\x1b[0m"
}

func colorEnabled(noColor bool) bool {
	if noColor {
		return false
	}
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}
	term := strings.TrimSpace(os.Getenv("TERM"))
	if term == "" || term == "dumb" {
		return false
	}
	return true
}
