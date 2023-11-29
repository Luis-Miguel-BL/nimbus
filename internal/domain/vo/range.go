package vo

import "fmt"

type Range struct {
	min int
	max int
}

func NewRange(min int, max int) (r Range, err error) {
	if min < 0 {
		return r, fmt.Errorf("invalid min range value")
	}
	if max < min {
		return r, fmt.Errorf("invalid range")
	}
	return Range{
		min: min,
		max: max,
	}, nil
}

func (v *Range) Min() int {
	return v.min
}
func (v *Range) Max() int {
	return v.max
}
