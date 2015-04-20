package fsm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	STATE_A = 0
	STATE_B = 1
	STATE_C = 2
	STATE_D = 3
)

func TestFsm(t *testing.T) {

	sm := NewFsm()

	handler_A := func() error {
		fmt.Println("Entering State A")
		return nil
	}

	handler_B := func() error {
		fmt.Println("Entering State B")
		return nil
	}

	handler_C := func() error {
		fmt.Println("Entering State C")
		return nil
	}

	handler_D := func() error {
		fmt.Println("Entering State D")
		return nil
	}

	sm.AddTransition(STATE_A, STATE_B, handler_B)
	sm.AddTransition(STATE_A, STATE_C, handler_C)
	sm.AddTransition(STATE_C, STATE_D, handler_D)
	sm.AddTransition(STATE_B, STATE_A, handler_A)

	sm.SetState(STATE_A)

	err := sm.TransitionTo(STATE_B)
	assert.Nil(t, err, "Valid Transition")

	err = sm.TransitionTo(STATE_D)
	assert.NotNil(t, err, "Invalid Transition")

	state_path, err := sm.FindTransitionPath(STATE_A, STATE_D)
	assert.NotNil(t, state_path)

	for _, state := range state_path {
		fmt.Println(state)
	}

}
