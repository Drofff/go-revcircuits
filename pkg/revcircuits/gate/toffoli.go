package gate

import (
	"fmt"

	"github.com/Drofff/go-revcircuits/pkg/revcircuits"
)

type toffoli struct {
	target   int
	controls []int
}

const TypeToffoli = "toffoli"

// NewToffoli accepts circuit line indexes for the target bit and any number of control bits.
// If all control bits are set to 1, the target bit will be flipped (0 to 1 or 1 to 0).
// If no control line indexes are provided, the resulted gate will always flip the target bit.
// F.e. use "NewToffoli(2, 0, 1)" to create the following gate on a 3-line circuit:
// --o--
// --o--
// --x--
func NewToffoli(target int, controls ...int) (revcircuits.Gate, error) {
	if target < 0 {
		return nil, fmt.Errorf("invalid index for the target bit: less than zero")
	}

	for i, c := range controls {
		if c < 0 {
			return nil, fmt.Errorf("invalid control bit index (c%v): less than zero", i)
		}
	}

	return &toffoli{
		target:   target,
		controls: controls,
	}, nil
}

func (t *toffoli) Type() string {
	return TypeToffoli
}

func (t *toffoli) UsedLines() []int {
	lines := make([]int, 1+len(t.controls))
	lines[0] = t.target
	for i, c := range t.controls {
		lines[i+1] = c
	}
	return lines
}

func flip(b byte) byte {
	if b == 0 {
		return 1
	}
	return 0
}

func (t *toffoli) Evaluate(input []byte) ([]byte, error) {
	gateActivated, err := evalControls(input, t.controls)
	if err != nil {
		return nil, fmt.Errorf("evaluate controls: %w", err)
	}

	if t.target >= len(input) {
		return nil, fmt.Errorf("target bit %v refers to a non-existant line (num of lines is %v)", t.target, len(input))
	}

	if !gateActivated {
		return input, nil
	}

	output := make([]byte, len(input))
	copy(output, input)
	output[t.target] = flip(input[t.target])
	return output, nil
}
