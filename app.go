//Package goo app.go this file has the global functions structs maps slices or anything else.
package goo

import (
	"bytes"
	"fmt"
	"strings"
)

//add the value to the slice separated by sep,and return a string of the value.
func Join(a []string, sep string, value ...string) (string, error) {
	var bu bytes.Buffer
	for _, v := range value {
		_, err := bu.WriteString(v)
		if err != nil {
			return "", fmt.Errorf("Join error,message: %v", err)
		}
		_, err = bu.WriteString(sep)
		if err != nil {
			return "", fmt.Errorf("Join error,message: %v", err)
		}
	}
	if len(a) == 0 {
		return bu.String(), nil
	}
	return strings.Join(a, sep) + bu.String(), nil

}
