package revcircuits

import (
	"fmt"
	"slices"
)

type Gate interface {
	Type() string
	UsedLines() []int
	Evaluate(input []byte) ([]byte, error)
}

type RevCircuit interface {
	Lines() int
	Gates() []Gate
	PlaceGates(gates ...Gate) error
	RemoveGates(gatePositions ...int)
	// Evaluate
	// input - binary data to put on the circuit, the length must be equal to Lines(), each element must be either 1 or 0.
	Evaluate(input []byte) ([]byte, error)
}

type revCircuit struct {
	lines int
	gates []Gate
}

func NewRevCircuit(lines int, gates ...Gate) (RevCircuit, error) {
	if lines < 1 {
		return nil, fmt.Errorf("a circuit must have, at least, one line")
	}
	return &revCircuit{lines: lines, gates: gates}, nil
}

func (rc *revCircuit) Lines() int {
	return rc.lines
}

func (rc *revCircuit) Gates() []Gate {
	return rc.gates
}

func (rc *revCircuit) PlaceGates(gates ...Gate) error {
	for _, g := range gates {
		for _, l := range g.UsedLines() {
			if l < 0 || l >= rc.lines {
				return fmt.Errorf("gate of type '%v' uses line '%v' which does not exist on the circuit", g.Type(), l)
			}
		}
	}

	rc.gates = append(rc.gates, gates...)
	return nil
}

func (rc *revCircuit) RemoveGates(gatePositions ...int) {
	var gates []Gate
	for i, g := range rc.gates {
		if !slices.Contains(gatePositions, i) {
			gates = append(gates, g)
		}
	}
	rc.gates = gates
}

func (rc *revCircuit) Evaluate(input []byte) ([]byte, error) {
	data := input
	var err error
	for _, g := range rc.gates {
		data, err = g.Evaluate(data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
