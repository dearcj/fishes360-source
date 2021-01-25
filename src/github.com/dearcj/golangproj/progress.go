package main

import (
	"time"
)

type ProfileStats struct {
	TotalTurn            int
	TotalTimeDuringMoves time.Duration
	AVGTurnTime          time.Duration
}

const XP_RECALC = 5
const GLOBAL_START_ENERGY = 100

type Progress struct {
	earlyTurn    bool
	earlyTurns   uint32 //amount of early turns in a row
	currentLevel string
	currentWorld string
	MaxHP        float64
}

func (p *Progress) resetProgress() {
}
