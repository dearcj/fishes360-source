package main

import (
	"github.com/dearcj/golangproj/network"
	"strings"
	"time"
)

type DamageType uint8
type SkillTimingType uint8
type SkillTarget uint32
type SkillTargetGroup uint32

type DamageTypes struct {
	DamageTypeHoly     DamageTypeInfo
	DamageTypePhysical DamageTypeInfo
	DamageTypeFire     DamageTypeInfo
	DamageTypeVoid     DamageTypeInfo
	DamageTypePoison   DamageTypeInfo
	DamageTypeWillOnly DamageTypeInfo
}

const (
	ULT_ENERGY_COST = 75
	AVG_ANIM_DUR    = 1500 * time.Millisecond
)

type DamageTypeInfo struct {
	ID          uint32
	Type        DamageType
	Description string
}

type SkillTimings struct {
	Immediate SkillTimingType
	OnCast    SkillTimingType
	NextTurn  SkillTimingType
}

var confDamageType = DamageTypes{
	DamageTypePhysical: DamageTypeInfo{
		ID:          0,
		Type:        1 << 0,
		Description: "physical damage",
	},
	DamageTypeHoly: DamageTypeInfo{
		ID:          1,
		Type:        1 << 1,
		Description: "holy damage",
	},
	DamageTypeFire: DamageTypeInfo{
		ID:          2,
		Type:        1 << 2,
		Description: "fire damage",
	},
	DamageTypePoison: DamageTypeInfo{
		ID:          3,
		Type:        1 << 3,
		Description: "poison damage",
	},
	DamageTypeVoid: DamageTypeInfo{
		ID:          4,
		Type:        1 << 4,
		Description: "void damage",
	},
	DamageTypeWillOnly: DamageTypeInfo{
		ID:          4,
		Type:        1 << 5,
		Description: "will damage",
	},
}

type SkillSlot = uint32

const (
	SkillSlotOther SkillSlot = iota
	SkillSlotMain  SkillSlot = iota
)

type BaseSkill struct {
	SkillSlot           SkillSlot
	SkipEnvAnim         bool
	EnergyCost          int
	AnimationTime       time.Duration
	NumTargets          int32 //Default 1, amount of targets from current target
	ID                  uint32
	UseRange            uint32
	Title               string
	TotalFrames         uint32
	DescriptionCompiled string
	Ico                 string
	SkillTiming         SkillTimingType
	Cooldown            int
	AOERadius           uint32
	EffectId            uint32
	Repeat              int
	AITag               int
	Target              SkillTarget
	TargetGroup         SkillTargetGroup
	RecastChance        int
}

func ColorTagString(s string, tag string) string {
	res := ""
	words := strings.Split(s, " ")
	for _, x := range words {
		res = res + " " + tag + x
	}

	return res
}

type Actions struct {
	Appear       data.ActionType
	FinishState  data.ActionType
	HealthChange data.ActionType
	ChangeState  data.ActionType
	StartScene   data.ActionType
	Remove       data.ActionType
	AngleChange  data.ActionType
	Shoot        data.ActionType
	MoneyChange  data.ActionType
	MakeBet      data.ActionType
	Critical     data.ActionType
	NoMoneyMode  data.ActionType
}

const (
	ANIM_ABSTRACT_CAST = iota
)

type StatusList struct {
	Stun   uint32
	Sleep  uint32
	Blind  uint32
	Warmth uint32
}

var confActions = (UNSAFE_INCREMENT_INT32_STRUCT(&Actions{})).(*Actions)

func InterfaceMapToLowerStr(x map[string]interface{}) map[string]uint32 {
	var mp = make(map[string]uint32)
	for k, v := range x {
		mp[strings.ToLower(k)] = v.(uint32)
	}

	return mp
}
