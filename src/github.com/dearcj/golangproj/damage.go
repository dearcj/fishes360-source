package main

type Damage struct {
	Value        float64
	Type         DamageType
	DoWillDamage bool
	IsCritical   bool
	Target       uint32
	Depth        uint32
	Heal         bool
	Skill        *BaseSkill
}

func (d *Damage) SetValue(v float64) *Damage {
	d.Value = v
	return d
}

func (d *Damage) RangeValue(min float64, max float64) *Damage {
	d.Value = -min - float64(server.Rand())*(max-min)
	return d
}

func (d *Damage) SetType(t DamageType) *Damage {
	d.Type = t
	return d
}

func (d *Damage) Inherit() *Damage {
	var dd = &Damage{}

	if d != nil {
		*dd = *d
	}
	dd.Depth++
	return dd
}

func (d *Damage) Clone() *Damage {
	newDamage := dmg(0, 0, nil)
	*newDamage = *d
	return newDamage
}

func dmg(value float64, dtype DamageType, skill *BaseSkill) *Damage {
	return &Damage{
		DoWillDamage: true,
		Value:        -value,
		Type:         dtype,
		Skill:        skill,
	}
}

func dmgRange(valueMin float64, valueMax float64, dtype DamageType, skill *BaseSkill) *Damage {
	return &Damage{
		DoWillDamage: true,
		Value:        -valueMin - float64(server.Rand())*(valueMax-valueMin),
		Type:         dtype,
		Skill:        skill,
	}
}

func dmgcrit(value float64, dtype DamageType, skill *BaseSkill) *Damage {
	return &Damage{
		DoWillDamage: true,
		Value:        -value,
		Type:         dtype,
		IsCritical:   true,
		Skill:        skill,
	}
}
