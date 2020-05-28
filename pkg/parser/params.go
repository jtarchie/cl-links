package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Range struct {
	Min, Max int
}

func NewRangeFromString(r string) *Range {
	min := 0
	max := math.MaxInt64

	if r[0] == '>' {
		min, _ = strconv.Atoi(r[1:])
	} else if r[0] == '<' {
		max, _ = strconv.Atoi(r[1:])
	} else {
		parts := strings.Split(r, "-")

		min, _ = strconv.Atoi(parts[0])
		if len(parts) == 2 {
			max, _ = strconv.Atoi(parts[1])
		}
	}
	return &Range{
		Min: min,
		Max: max,
	}
}

type Params struct {
	fields map[string]interface{}
}

func (q Params) GetRange(s string) (*Range, error) {
	v, ok := q.fields[s]
	if !ok {
		return nil, nil
	}

	if r, ok := v.(*Range); ok {
		return r, nil
	}

	return nil, fmt.Errorf("could not convert %q to a Range type", s)
}

func (q Params) GetBoolean(s string) (bool, error) {
	v, ok := q.fields[s]
	if !ok {
		return false, nil
	}

	if r, ok := v.(bool); ok {
		return r, nil
	}

	return false, fmt.Errorf("could not convert %q to a bool type", s)
}

func (q Params) GetString(s string) (string, error) {
	v, ok := q.fields[s]
	if !ok {
		return "", nil
	}

	if r, ok := v.(string); ok {
		return r, nil
	}

	return "", fmt.Errorf("could not convert %q to a string type", s)
}

func (q Params) Keys() []string {
	keys := []string{}

	for k, _ := range q.fields {
		keys = append(keys, k)
	}

	return keys
}

func NewParams(fields map[string]interface{}) *Params {
	return &Params{fields: fields}
}
