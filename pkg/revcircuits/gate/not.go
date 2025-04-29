package gate

import (
	"fmt"

	"github.com/Drofff/go-revcircuits/pkg/revcircuits"
)

type not struct {
	toffoli
}

const TypeNot = "not"

// NewNot accepts a circuit line index for the bit to be flipped (0 to 1 or 1 to 0).
func NewNot(target int) (revcircuits.Gate, error) {
	if target < 0 {
		return nil, fmt.Errorf("invalid index for the target bit: less than zero")
	}

	t, err := NewToffoli(target)
	if err != nil {
		return nil, err
	}
	return &not{toffoli: *t.(*toffoli)}, nil
}

func (n *not) Type() string {
	return TypeNot
}
