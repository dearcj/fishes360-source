package main

import "time"

const OVERKILL_DAMAGE = 5

type OnKillCallback func(u *BaseCharacter) FList

type BaseCharacter struct {
	Object *Object

	Killed       bool
	HP           float64
	onKill       OnKillCallback
	MaxHP        float64
	CurrentSkill *SkillUsing
}

type SkillUsing struct {
	startIteration uint32
	prop           float64
}

func (a *BaseCharacter) GetTeam() []*BaseCharacter {
	return nil
}

func (a *BaseCharacter) OnDestroy() {
	//a.Object = nil
	//a.sm = nil
}

func (a *BaseCharacter) SetKilled(state bool, dmg *Damage, damageDealer *BaseCharacter) (list FList) {
	if !a.Killed {
		a.Killed = state

		if a.onKill != nil {
			list = list.Add(a.onKill(a))
		}

		a.Object.Remove()
	}

	return
}

func GetSkillProb(baseChance float64, will uint32, maxwill uint32) float64 {
	if maxwill == 0 {
		return 10000
	}

	return baseChance * (1 - (float64(will)/float64(maxwill))*0.75)
}

/*func (a *BaseCharacter) DealDamage(damageDealer *BaseCharacter, damage *Damage) (list FList) {
	damage.Target = a.Object.ID
	if a.Killed || a.Object.doRemove {
		return nil
	}
	//ai := a.Object.FindByComponent(confComponents.Fish)

	var onDeal, onTake FList

	if damage.Value > 0 && damage.Heal {
		damage.Value = -damage.Value
	}

	//DAMAGE => DAMAGE
	if damage.Value < 0 {
		if damage.Value > 0 {
			damage.Value = 0
		}
	}

	if damage.Type != confDamageType.DamageTypeWillOnly.Type {
		deltaInt := math.Round(damage.Value)
		//oldhp := a.HP

		a.HP += deltaInt
		if a.HP > a.MaxHP {
			a.HP = a.MaxHP
		}

		if a.HP < 0 {
			a.HP = 0
		}

		//server.logger.Debug("Changed hp from ", zap.Float64("old", oldhp), zap.Float64("new", a.HP), zap.Float64("delta", deltaInt))

		list = list.AddSingle(a.Object.Effect(confActions.HealthChange).
			V(float32(deltaInt)))
	}

	list = list.Add(onDeal, onTake)

	if a.HP <= 0 && a.Killed == false {

		if damageDealer != nil {
			pl := damageDealer.Object.FindByComponent(confComponents.Player)
			if pl != nil {
				fish := a.Object.FindByComponent(confComponents.Fish)
				if fish != nil {

					_, money := fish.(*Fish).getMoneyPrize(damageDealer.Object.session, pl.(*Player).currentGun.Damage)

					damageDealer.Object.session.NeedToKnow(pl.(*Player).AddMoney(float32(money)))
				}
			}
		}

		list = list.Add(a.Object.factory.RemoveObject(a.Object, true, false))
		//list = list.Add(a.SetKilled(true, damage, damageDealer))
		//server.logger.Debug("Killed")
	}

	return list
}*/

func (a *BaseCharacter) DealProbDamage(damageDealer *BaseCharacter, gun *Gun) (list FList) {
	if a.Killed || a.Object.doRemove {
		return nil
	}

	list = list.AddSingle(a.Object.Effect(confActions.HealthChange).
		V(0))

	prob := gun.CalcProb(a.Object.FindByComponent(confComponents.Fish).(*Fish).FishConfig, a.Object.factory.run.timeline.RTP)

	deltaHp := gun.CalcHpDelta(a.Object.FindByComponent(confComponents.Fish).(*Fish).FishConfig, a.Object.BaseCharacter().HP, a.Object.BaseCharacter().MaxHP, a.Object.factory.run.timeline.RTP)
	oldHp := a.HP
	a.HP = a.HP + deltaHp
	list = list.AddSingle(a.Object.Effect(confActions.HealthChange).
		V(float32(deltaHp)))


	killed := prob > server.Rand()

	if killed {

		if damageDealer != nil {
			pl := damageDealer.Object.FindByComponent(confComponents.Player)
			if pl != nil {
				fish := a.Object.FindByComponent(confComponents.Fish)
				if fish != nil {

					prop := oldHp / a.MaxHP
					if prop > 0.7 && prob < 0.11 {
						damageDealer.Object.session.run.team.NeedToKnow(a.Object.Effect(confActions.Critical) )
					}

					_, money := fish.(*Fish).getMoneyPrize(damageDealer.Object.session, pl.(*Player).currentGun.Damage, a.Object.factory.run.RoomCoef, a.Object.factory.run.timeline.RTP)

					damageDealer.Object.session.NeedToKnow(pl.(*Player).AddMoney(float32(money)))
					r := damageDealer.Object.session.run
					//timeline := r.timeline
					/*if fish.(*Fish).FishConfig.IsBoss {
						if timeline.CurrentScene.Name == "aquaman" {
							if !r.transition {
								r.transition = true
								r.async.DelayedCall(func() {
									r.transition = false
									r.StartScene("deadfish")
								}, time.Second*2)

							}
						}
					}*/
				}
			}
		}
		list = list.Add(a.Object.factory.RemoveObject(a.Object, true, false))

	}

	return
}
