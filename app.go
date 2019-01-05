//Package goo app.go this file has the global functions structs maps slices or anything else.
package goo

import (
	"bytes"
	"fmt"
)

//Join is like strings.Join, support multiple parameters
func Join(a []string, sep ...string) (string, error) {
	var bu bytes.Buffer
	for _, v := range sep {
		_, err := bu.WriteString(v)
		if err != nil {
			return "", fmt.Errorf("Join error,message: %v", err)
		}
	}
	return bu.String(), nil

}
