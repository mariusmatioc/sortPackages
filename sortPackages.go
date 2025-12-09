package pkg

import "fmt"

const (
	REJECTED = "REJECTED"
	SPECIAL  = "SPECIAL"
	STANDARD = "STANDARD"
)

func Sort(width, height, length, mass float64) string {
	bulky, err := isBulky(width, height, length)
	if err != nil {
		return REJECTED
	}
	heavy, err := isHeavy(mass)
	if err != nil {
		return REJECTED
	}
	if heavy && bulky {
		return REJECTED
	}
	if heavy || bulky {
		return SPECIAL
	}
	return STANDARD
}

func isBulky(width, height, length float64) (bool, error) {
	if width <= 0 || height <= 0 || length <= 0 {
		return false, fmt.Errorf("invalid dimensions")
	}
	volume := width * height * length
	if volume >= 1000000 {
		return true, nil
	}
	if width >= 150 || height >= 150 || length >= 150 {
		return true, nil
	}
	return false, nil
}

func isHeavy(mass float64) (bool, error) {
	if mass <= 0 {
		return false, fmt.Errorf("invalid weight")
	}
	return mass >= 20, nil
}
