package vo

import (
	"fmt"
	"math"
)

type Percentage struct {
	percentage float64
}

func NewPercentage(percentage float64) (p Percentage, err error) {
	if percentage < 0 {
		return p, fmt.Errorf("invalid percentage")
	}
	return Percentage{percentage}, nil
}

func (v Percentage) ApplyPercentage(value int) (newValue int) {
	newValue += int(math.Round(float64(value) * v.percentage))

	return newValue
}
