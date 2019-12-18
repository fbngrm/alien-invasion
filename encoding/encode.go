package encoding

import (
	"fmt"

	"github.com/fgrimme/alien-invasion/world"
)

// Encode encodes m into a string representation.
func Encode(m world.Map) string {
	var s string
	for cityName, links := range m {
		if len(links) == 0 {
			continue
		}
		s = fmt.Sprintf("%s%s", s, cityName)
		for link, direction := range links {
			s = fmt.Sprintf("%s %s=%s", s, direction, link)
		}
		s = fmt.Sprintf("%s\n", s)
	}
	return s
}
