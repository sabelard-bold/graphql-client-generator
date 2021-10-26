package flag

import (
	"errors"
	"fmt"
	"strings"
)

type HTTPHeader struct {
	Name  string
	Value string
}

type HTTPHeaders []HTTPHeader

func (f *HTTPHeaders) String() string {
	parts := []string{}

	for _, v := range *f {
		parts = append(parts, fmt.Sprintf("%s:%s", v.Name, v.Value))
	}

	return strings.Join(parts, ", ")
}

func (f *HTTPHeaders) Set(value string) error {
	parts := strings.Split(value, ":")

	if len(parts) != 2 {
		return errors.New("header must be in the format name:value")
	}

	*f = append(*f, HTTPHeader{
		Name:  parts[0],
		Value: parts[1],
	})

	return nil
}
