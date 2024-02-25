package youtube

import (
	"regexp"

	"github.com/rwxrob/to"
)

var LinkExp = regexp.MustCompile(`([^<]|^)(https://youtu\.be\/\S+)`)

// Linkify converts any Youtube's URL on it's own line anythiong that can be passed into a valid markdown URL.
// wrapped with angle brackets and prefixed with the optionally passed prefix (default is a TV)
// returns empty String if no arguments are passed.

func Linkify(args ...any) string {

	var in, pre string
	ln := len(args)

	switch {

	case (ln == 0):
		return ""
	case ln >= 2:
		pre = to.String(args[1])
		fallthrough
	case ln == 1:
		pre = "📺 "
		in = to.String(args[0])
	case ln == 2:

	}
	// TODO find and replace
	return LinkExp.ReplaceAllString(in, "$1"+pre+"<$2>")
}
