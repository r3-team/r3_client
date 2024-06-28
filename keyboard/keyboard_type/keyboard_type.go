package keyboard_type

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-vgo/robotgo"
)

var (
	keysSupported     = []string{"ALT-TAB", "ESCAPE", "ENTER", "TAB"}
	regexCharsWithKey = regexp.MustCompile(fmt.Sprintf(`(?m:^(.*?)(\{(%s)\})?$)`, strings.Join(keysSupported, "|"))) // finds 0+ chars before one or zero special keys
	regexKeysFind     = regexp.MustCompile(fmt.Sprintf(`(\{(%s)\})`, strings.Join(keysSupported, "|")))              // finds special keys: {TAB}, {ENTER}, etc.
)

// executes keyboard inputs
// supports unicode characters to type as well some special keys, defined by their placeholders: {TAB}, {ENTER}, ...
func Do(input string) {

	// add newlines after each supported special key
	// each line then contains 0+ chars to type and one or zero special keys to execute
	inputNewLines := regexKeysFind.ReplaceAllString(input, "$1\n")

	// find chars + special key for each line
	matches := regexCharsWithKey.FindAllStringSubmatch(inputNewLines, -1)
	for _, matchesSub := range matches {

		// characters
		if len(matchesSub) > 1 && matchesSub[1] != "" {
			for _, c := range strings.Split(matchesSub[1], "") {
				robotgo.UnicodeType(uint32(robotgo.CharCodeAt(c, 0)))
				// some latency after each character
				robotgo.MilliSleep(40)
			}
		}

		// special key
		if len(matchesSub) > 3 {
			// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys

			switch matchesSub[3] {
			case "ALT-TAB":
				robotgo.KeyDown("alt")
				robotgo.KeyTap("tab")
				robotgo.KeyUp("alt")
			case "ESCAPE":
				robotgo.KeyTap("escape")
			case "ENTER":
				robotgo.KeyTap("enter")
			case "TAB":
				robotgo.KeyTap("tab")
			}
			// some latency after an action occurred
			robotgo.MilliSleep(100)
		}
	}
}
