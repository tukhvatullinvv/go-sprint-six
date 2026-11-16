package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Transform(s string) string {
	if s == "" {
		return "initial string is empty"
	}
	if strings.Contains(s, "-") {
		return morse.ToText(s)
	} else {
		return morse.ToMorse(s)
	}
}
