package revcircuits_test

import (
	"slices"
	"testing"

	"github.com/Drofff/go-revcircuits/pkg/revcircuits"
	"github.com/Drofff/go-revcircuits/pkg/revcircuits/gate"
)

func TestRevCircuit_AllGates(t *testing.T) {
	toffoli, err := gate.NewToffoli(2, 0, 1)
	if err != nil {
		t.Fatalf("failed to create a toffoli gate: %e", err)
	}

	fredkin, err := gate.NewFredkin(3, 4, 5)
	if err != nil {
		t.Fatalf("failed to create a frednkin gate: %e", err)
	}

	cnot, err := gate.NewCnot(6, 7)
	if err != nil {
		t.Fatalf("failed to create a cnot gate: %e", err)
	}

	not, err := gate.NewNot(8)
	if err != nil {
		t.Fatalf("failed to create a not gate: %e", err)
	}

	rc, err := revcircuits.NewRevCircuit(10, toffoli, fredkin, cnot, not)
	if err != nil {
		t.Fatalf("failed to create the circuit: %e", err)
	}

	input1 := []byte{
		// toffoli input
		1, 1, 0,
		// fredkin input
		0, 1, 1,
		// cnot input
		0, 1,
		// not input
		0,
		// spare bit
		1,
	}
	wantOutput1 := []byte{
		// toffoli output
		1, 1, 1,
		// fredkin output
		1, 0, 1,
		// cnot output
		1, 1,
		// not output
		1,
		// spare bit out
		1,
	}

	output1, err := rc.Evaluate(input1)
	if err != nil {
		t.Fatalf("failed to evaluate the circuit (1): %e", err)
	}

	if output1[0] != wantOutput1[0] || output1[1] != wantOutput1[1] || output1[2] != wantOutput1[2] {
		t.Errorf("invalid toffoli output; expected %v but got %v", wantOutput1[:3], output1[:3])
	}

	if output1[3] != wantOutput1[3] || output1[4] != wantOutput1[4] || output1[5] != wantOutput1[5] {
		t.Errorf("invalid fredkin output; expected %v but got %v", wantOutput1[3:6], output1[3:6])
	}

	if output1[6] != wantOutput1[6] || output1[7] != wantOutput1[7] {
		t.Errorf("invalid cnot output; expected %v but got %v", wantOutput1[6:8], output1[6:8])
	}

	if output1[8] != wantOutput1[8] {
		t.Errorf("invalid not output; expected %v but got %v", wantOutput1[8], output1[8])
	}

	if output1[9] != wantOutput1[9] {
		t.Errorf("unexpected mutation of the spare bit; expected %v but got %v", wantOutput1[9], output1[9])
	}

	input2 := []byte{
		// toffoli input
		1, 0, 0,
		// fredkin input
		0, 1, 0,
		// cnot input
		1, 0,
		// not input
		0,
		// spare bit
		0,
	}
	wantOutput2 := []byte{
		// toffoli output
		1, 0, 0,
		// fredkin output
		0, 1, 0,
		// cnot output
		1, 0,
		// not output
		1,
		// spare bit out
		0,
	}

	output2, err := rc.Evaluate(input2)
	if err != nil {
		t.Fatalf("failed to evaluate the circuit (2): %e", err)
	}

	if output2[0] != wantOutput2[0] || output2[1] != wantOutput2[1] || output2[2] != wantOutput2[2] {
		t.Errorf("invalid toffoli output; expected %v but got %v", wantOutput2[:3], output2[:3])
	}

	if output2[3] != wantOutput2[3] || output2[4] != wantOutput2[4] || output2[5] != wantOutput2[5] {
		t.Errorf("invalid fredkin output; expected %v but got %v", wantOutput2[3:6], output2[3:6])
	}

	if output2[6] != wantOutput2[6] || output2[7] != wantOutput2[7] {
		t.Errorf("invalid cnot output; expected %v but got %v", wantOutput2[6:8], output2[6:8])
	}

	if output2[8] != wantOutput2[8] {
		t.Errorf("invalid not output; expected %v but got %v", wantOutput2[8], output2[8])
	}

	if output2[9] != wantOutput2[9] {
		t.Errorf("unexpected mutation of the spare bit; expected %v but got %v", wantOutput2[9], output2[9])
	}
}

func TestRevCircuit_SequenceOfGates(t *testing.T) {
	// --x--x--x--
	// -----x-----
	// --o--o-----
	// --o--------
	//   T  F  CN

	toffoli, err := gate.NewToffoli(0, 2, 3)
	if err != nil {
		t.Fatalf("new toffoli: %e", err)
	}
	fredkin, err := gate.NewFredkin(0, 1, 2)
	if err != nil {
		t.Fatalf("new fredkin: %e", err)
	}
	not, err := gate.NewNot(0)
	if err != nil {
		t.Fatalf("new not: %e", err)
	}
	rc, err := revcircuits.NewRevCircuit(3, toffoli, fredkin, not)
	if err != nil {
		t.Fatalf("new circuit: %e", err)
	}

	input := []byte{0, 0, 1, 1}
	output, err := rc.Evaluate(input)
	if err != nil {
		t.Fatalf("evaluate the circuit: %e", err)
	}

	wantOutput := []byte{1, 1, 1, 1}
	if !slices.Equal(wantOutput, output) {
		t.Fatalf("unexpected output; wanted %v but got %v", wantOutput, output)
	}
}
