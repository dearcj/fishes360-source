package main

import (
	"time"
)

const ENDLESS = 99999 * time.Hour

type TransitionFSM struct {
	next         NextStateFunc
	doTransition bool
}

func (fsm *TransitionFSM) Setup(doTransition bool, next NextStateFunc) {
	fsm.doTransition = doTransition
	fsm.next = next
}

type NextRoomFSM struct {
}

type GameOverAllFSM struct {
}
