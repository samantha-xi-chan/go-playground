package play220_fsm

import (
	"fmt"
	"time"
)

type State int

const (
	Idle State = iota
	Processing
	Completed
)

type Event int

const (
	Start Event = iota
	Process
	Finish
	Timeout
	Retry
)

type StateMachine struct {
	currentState State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{currentState: Idle}
}

func (sm *StateMachine) transition(event Event) {
	switch sm.currentState {
	case Idle:
		if event == Start {
			sm.currentState = Processing
			fmt.Println("Transition: Idle -> Processing")
		}
	case Processing:
		if event == Process {
			// Simulate processing
			fmt.Println("Processing...")
			time.Sleep(2 * time.Second)
			sm.currentState = Completed
			fmt.Println("Transition: Processing -> Completed")
		} else if event == Timeout {
			fmt.Println("Transition: Processing -> Idle (Timeout)")
			sm.currentState = Idle
		} else if event == Retry {
			fmt.Println("Transition: Processing -> Idle (Retry)")
			sm.currentState = Idle
		}
	case Completed:
		if event == Finish {
			sm.currentState = Idle
			fmt.Println("Transition: Completed -> Idle")
		}
	}
}

func Play() {
	sm := NewStateMachine()

	// Trigger the state transitions
	sm.transition(Start)
	sm.transition(Process)
	sm.transition(Finish)

	// Simulate a timeout and retry
	fmt.Println("Simulating timeout and retry...")
	go func() {
		time.Sleep(3 * time.Second)
		sm.transition(Timeout) // Transition due to timeout
		time.Sleep(2 * time.Second)
		sm.transition(Retry) // Retry the transition
	}()

	// Wait for the retry
	time.Sleep(6 * time.Second)
}
