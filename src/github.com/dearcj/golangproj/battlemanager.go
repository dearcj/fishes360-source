package main

import (
	as "github.com/dearcj/golangproj/asyncutils"
	"time"
)

type BattleResult uint32
type BattleMode uint32

const (
	BR_PLAYERS_WIN BattleResult = iota
	BR_ENEMIES_WIN
)

type FListFunc func() FList
type BattleFSM struct {
	run             *Run
	cbAfterTurn     *as.AsyncUtil
	cbAfterBattle   *as.AsyncUtil
	afterSkillFuncs []FListFunc
}

type NextBattleFSM struct {
	*BattleFSM
}

func CreateBattleManager() *BattleFSM {
	farFuture := time.Now().Add(time.Hour * 999999)

	return &BattleFSM{
		cbAfterBattle: as.CreateAsycnUtil(&farFuture),
		cbAfterTurn:   as.CreateAsycnUtil(&farFuture),
	}
}

func (bm *BattleFSM) CallAfterTurn(f as.AsyncCallbackFunction) {
	bm.cbAfterTurn.DelayedCall(f, 0)
}

func (bm *BattleFSM) CallAfterbattle(f as.AsyncCallbackFunction) {
	bm.cbAfterBattle.DelayedCall(f, 0)
}

func (bm *BattleFSM) FinishTurn(s *Session) {
}

func (bm *BattleFSM) StartBattle(r *Run) (TurnStateType, time.Duration, CurStateFunc, NextStateFunc) {
	return 0, config.Player.PrepareTime, func() {
		//bm.BuffsOnBattleStart()
	}, bm.SetPlayerTurn
}

func (bm *BattleFSM) EffectsAfterMove(f FListFunc) {
	bm.afterSkillFuncs = append(bm.afterSkillFuncs, f)
}

func (bm *BattleFSM) ExecBuffsHarm() FList {
	var se FList
	for _, a := range bm.run.factory.actors {
		unit := a.FindByComponent(config.Components.Unit)
		if unit != nil {
			/*barray := unit.(*Unit).Buffs
			for _, b := range barray {
				if b.DoEveryTurn != nil {
					se = append(se, b.DoEveryTurn()...)
				}
			}*/
		}
	}

	return se
}

func (bm *BattleFSM) SetPlayerTurn(r *Run) (TurnStateType, time.Duration, CurStateFunc, NextStateFunc) {

	if bm.run.turnNumber > 0 {
		for _, a := range r.factory.actors {
			player := a.FindByComponent(config.Components.Player)
			if player != nil && a.session != nil && r.stateFinished[a.session.id] == nil {
				bm.FinishTurn(a.session)
			}
		}
	}

	bm.run.turnNumber++

	bm.run.playerTurnStarted = time.Now()

	return 0, config.Player.TurnTime, nil, bm.SetExecPlayerTurn
}

func (bm *BattleFSM) BotsFinishScene() {
	bm.run.async.DelayedCall(func() {
		for _, session := range bm.run.team.sessions {
			if session.IsBot() {
				bm.run.FinishSessionState(session)
			}
		}
	}, 2*time.Second)
}

func (bm *BattleFSM) SetExecPlayerTurn(r *Run) (TurnStateType, time.Duration, CurStateFunc, NextStateFunc) {
	/*var earlyTurnEffects FList

	//	r.SetState(config.TurnState.BATTLE_BEFORE_TURN, 0, nil)

	bm.run.team.NeedToKnow(earlyTurnEffects)

	//TODO: ADD CAN USE CHECKING AND MOVES CHECKING
	//TODO: ADD 1 SKILL CHECKING AND 1 ITEM CHECKING
	var br BattleResult
	var turnEffects FList
	br, turnEffects, _ = bm.run.BlindGame.ExecTurnActions(bm.run.turnActions.actions)
	buffHarmsEffects := bm.ExecBuffsHarm()

	bm.run.team.NeedToKnow(turnEffects, buffHarmsEffects)
	bm.cbAfterTurn.Exec()
	bm.run.BlindGame.ClearActions(nil, 0, 0)

	if br != BR_NONE {
		bm.cbAfterBattle.Exec()
		battleEndEffect, _ := bm.BattleEnd()
		r.team.NeedToKnow(battleEndEffect)

		for _, s := range r.team.sessions {
			s.NeedToKnow(s.player.Effect(confActions.AnimationEnded))
		}

		//		return config.TurnState.BATTLE_AFTER_TURN, 0, r.HoldUntilAnimation(func(r *Run) (TurnStateType, time.Duration, NextStateFunc) {
		//			return bm.BattleFinish(br, killed)
		//		})
	} else {
		//	r.team.NeedToKnow(cf(confActions.StartEnemyTurn))

		//		return bm.SetEnemyTurn(r)
	}
	*/
	return 0, 0, nil, nil
}
