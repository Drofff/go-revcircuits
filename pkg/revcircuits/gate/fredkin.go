package gate

import (
	"fmt"

	"github.com/Drofff/go-revcircuits/pkg/revcircuits"
)

type fredkin struct {
	target1  int
	target2  int
	controls []int
}

const TypeFredkin = "fredkin"

// NewFredkin accepts circuit line indexes for the two target bits to be swapped and any number of control bits.
// If no control bits are set, the swap will always happen i.e. the control will always be assumed equal 1.
// F.e. use "NewFredkin(0, 1, 3)" to create the following gate on the 4 line circuit:
// --x--
// --x--
// -----
// --o--
func NewFredkin(target1, target2 int, controls ...int) (revcircuits.Gate, error) {
	if target1 < 0 {
		return nil, fmt.Errorf("invalid index for the first target bit: less than zero")
	}
	if target2 < 0 {
		return nil, fmt.Errorf("invalid index for the second target bit: less than zero")
	}

	for i, c := range controls {
		if c < 0 {
			return nil, fmt.Errorf("invalid control bit index (c%v): less than zero", i)
		}
	}

	return &fredkin{
		target1:  target1,
		target2:  target2,
		controls: controls,
	}, nil
}

func (f *fredkin) Type() string {
	return TypeFredkin
}

func (f *fredkin) UsedLines() []int {
	lines := make([]int, 2+len(f.controls))
	lines[0] = f.target1
	lines[1] = f.target2
	for i, c := range f.controls {
		lines[i+2] = c
	}
	return lines
}

func (f *fredkin) swap(input []byte) []byte {
	output := make([]byte, len(input))
	copy(output, input)

	output[f.target1] = input[f.target2]
	output[f.target2] = input[f.target1]
	return output
}

func (f *fredkin) Evaluate(input []byte) ([]byte, error) {
	gateActivated, err := evalControls(input, f.controls)
	if err != nil {
		return nil, fmt.Errorf("evaluate controls: %w", err)
	}

	if f.target1 >= len(input) || f.target2 >= len(input) {
		return nil, fmt.Errorf("target bits (%v, %v) refrer to non-existant line/s (num of lines is %v)", f.target1, f.target2, len(input))
	}

	if !gateActivated {
		return input, nil
	}

	return f.swap(input), nil
}
