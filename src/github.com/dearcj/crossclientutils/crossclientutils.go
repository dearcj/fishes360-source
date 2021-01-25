package main

type CrossClientUtils struct{}

func (z *CrossClientUtils) GetSkillProb(baseChance float64, will float64, maxwill float64) float64 {
	return baseChance * (1 - (will/maxwill)*0.75)
}

func main() {
	var y CrossClientUtils
	var z CrossClientUtils = y
	y = z
}
