package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("error: not same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
