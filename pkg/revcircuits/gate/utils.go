package gate

import "fmt"

func evalControls(input []byte, controls []int) (bool, error) {
	if len(controls) == 0 {
		return true, nil
	}

	res := true
	for i, c := range controls {
		if len(input) <= c {
			return false, fmt.Errorf("control c%v refers to a non-existant line %v (num of lines is %v)", i, c, len(input))
		}

		if input[c] == 0 {
			res = false
			break
		}
	}
	return res, nil
}
