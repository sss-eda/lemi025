package text

import (
	"fmt"
	"regexp"
)

// https://stackoverflow.com/a/44278850/13169340
func groupmap(b []byte, r *regexp.Regexp) (map[string][]byte, error) {
	values := r.FindSubmatch(b)
	if values == nil {
		return nil, fmt.Errorf("string does not match regexp")
	}
	keys := r.SubexpNames()

	// create map
	d := make(map[string][]byte)
	for i := 1; i < len(keys); i++ {
		d[keys[i]] = values[i]
	}

	return d, nil
}
