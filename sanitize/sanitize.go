package sanitize

import "strings"

func Sanitize(string string) string {
	string = strings.ReplaceAll(string, "/", "-")
	return strings.ReplaceAll(string, ":", " -")
}
