package fsm

import (
	"errors"
)

type Handler func(*StateMachine) error

type TransitionEntry struct {
	ToState int
	Handler Handler
	Args    interface{}
}

type StateMachine struct {
	TransitionTable map[int][]TransitionEntry
	CurrentState    int
}

func NewFsm() *StateMachine {

	return &StateMachine{
		TransitionTable: make(map[int][]TransitionEntry),
		CurrentState:    -1,
	}
}

func (sm *StateMachine) AddTransition(from_state int, to_state int, handler Handler) {

	// check if from_state exisits in the table
	_, ok := sm.TransitionTable[from_state]
	if !ok {
		sm.TransitionTable[from_state] = make([]TransitionEntry, 0)
	}

	_, ok = sm.TransitionTable[to_state]
	if !ok {
		sm.TransitionTable[to_state] = make([]TransitionEntry, 0)
	}

	sm.TransitionTable[from_state] = append(
		sm.TransitionTable[from_state],
		TransitionEntry{ToState: to_state, Handler: handler},
	)

}

func (sm *StateMachine) SetState(state int) {
	sm.CurrentState = state
}

func (sm *StateMachine) TransitionTo(to_state int) error {

	state_list := sm.TransitionTable[sm.CurrentState]
	//check if to_state is in the Transitiontable

	var handler Handler = nil

	for _, e := range state_list {
		if e.ToState == to_state {
			handler = e.Handler
		}
	}

	if handler == nil {
		return errors.New("Invalid: state transition not found in transition table")
	}

	//call the handler
	err := handler(sm)

	if err != nil {
		return err
	}

	// all OK, move to new state
	sm.CurrentState = to_state
	return nil
}

//does a DFS from source state to dest state and returns the series of
//transitions required to reach the end state
func (sm *StateMachine) FindTransitionPath(from_state int, to_state int) ([]int, error) {

	visited := make(map[int]bool, len(sm.TransitionTable))

	for key := range sm.TransitionTable {
		visited[key] = false
	}

	path := sm.dfs(from_state, to_state, visited)
	if path != nil {

		rev_path := make([]int, len(path))
		for i := len(path) - 1; i >= 0; i-- {
			rev_path[len(path)-i-1] = path[i]
		}

		return rev_path, nil
	}

	return nil, errors.New("No transition path")
}

//actual DFS
func (sm *StateMachine) dfs(cur_state int, dst_state int, visited map[int]bool) []int {

	if cur_state == dst_state {
		return []int{cur_state} //found it
	}

	visited[cur_state] = true
	for _, e := range sm.TransitionTable[cur_state] {
		new_state := e.ToState
		if !visited[new_state] {
			path := sm.dfs(new_state, dst_state, visited)
			if path != nil { //someone found a path below
				return append(path, cur_state)
			}

		}

	}

	return nil
}
