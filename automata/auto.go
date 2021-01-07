package automata

import (
	"fmt"
	"strings"
)

type state string

type stateSet map[state]bool

func (set stateSet) discription() string {
	var sb strings.Builder
	sb.WriteString("{")
	for k, v := range set {
		if v {
			fmt.Fprintf(&sb, "%s,", k)
		}
	}
	return strings.TrimSuffix(sb.String(), ",") + "}"
}

func (set stateSet) add(s state, ss ...state) {
	set[s] = true
	for _, v := range ss {
		set[v] = true
	}
}

type symbol rune

type ttab map[state]tst

func (t ttab) get(s state, sym symbol) stateSet {
	if _, ok := t[s]; !ok {
		t[s] = make(tst)
	}
	tt := t[s]
	if _, ok := tt[sym]; !ok {
		tt[sym] = make(stateSet)
	}
	return tt[sym]
}

func (t ttab) getAllStates() stateSet {
	set := make(stateSet)
	for s, t := range t {
		set.add(s)
		for _, tt := range t {
			for ss, v := range tt {
				if v {
					set.add(ss)
				}
			}
		}
	}
	return set
}

type tst map[symbol]stateSet

// NFA  without epsilon transation
type NFA struct {
	q0    state
	delta ttab
	final stateSet
}

// NewNFA ...
func NewNFA(q0 state, final ...state) *NFA {
	nfa := new(NFA)
	nfa.q0 = q0
	nfa.final = make(stateSet)
	for _, v := range final {
		nfa.final.add(v)
	}
	nfa.delta = make(ttab)
	return nfa
}

// AddTrasition ...
func (nfa *NFA) AddTrasition(q state, s symbol, next state, nexts ...state) {
	nfa.delta.get(q, s).add(next, nexts...)
}

// Discription ...
func (nfa *NFA) Discription() string {
	// Q
	// gether all the states from delta
	var sb strings.Builder
	fmt.Fprintf(&sb, "Q = %s", nfa.delta.getAllStates().discription())
	// Σ

}

// DFA ...
type DFA struct {
	nfa *NFA
}
