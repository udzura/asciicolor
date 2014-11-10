package asciicolor

import (
	"testing"
)

// see http://en.wikichip.org/wiki/irc/colors
//     http://www.visualirc.net/tech-attrs.php
//     https://golang.org/ref/spec#String_literals

func assertEqual(t *testing.T, src string, expected string) {
	if src != expected {
		t.Logf("\"%+v\" expected, got \"%+v\"", expected, src)
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	src := "<red>Hello, colored<end>"
	expected := "\x034Hello, colored\x0f"
	assertEqual(t, src, expected)
}
