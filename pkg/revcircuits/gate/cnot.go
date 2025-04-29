package gate

import (
	"fmt"

	"github.com/Drofff/go-revcircuits/pkg/revcircuits"
)

type cnot struct {
	toffoli
}

const TypeCnot = "cnot"

// NewCnot accepts circuit line indexes for the target bit and a single control bit.
// If the control bit is set to 1, the target bit will be flipped (0 to 1 or 1 to 0).
// F.e. use "NewCnot(1, 0)" to create the following gate on a 2-line circuit:
// --o--
// --x--
func NewCnot(target, control int) (revcircuits.Gate, error) {
	if target < 0 {
		return nil, fmt.Errorf("invalid index for the target bit: less than zero")
	}
	if control < 0 {
		return nil, fmt.Errorf("invalid index for the control bit: less than zero")
	}
	t, err := NewToffoli(target, control)
	if err != nil {
		return nil, err
	}
	return &cnot{toffoli: *t.(*toffoli)}, nil
}

func (c *cnot) Type() string {
	return TypeCnot
}
