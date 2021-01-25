package main

import "sort"

type ProbDoFunction func()

type ProbCaseArray []ProbCase

type ProbCase struct {
	probCoef float64
	do       ProbDoFunction
	data     interface{}
}

func ProbAbsolete(probs *[]ProbCase) int {
	p := *probs
	sort.Slice(p, func(i int, j int) bool {
		return p[i].probCoef > p[j].probCoef
	})

	var bestInx = -1
	for inx, x := range p {
		if server.Rand() < x.probCoef {
			bestInx = inx
			///res = x
			break
		}
	}

	if bestInx >= 0 {
		if (*probs)[bestInx].do != nil {
			(*probs)[bestInx].do()
		}
	}

	return bestInx
}

func ProbArray(probs *[]ProbCase) int {
	summ := 0.0
	for _, x := range *probs {
		summ += x.probCoef
	}

	seek := server.Rand() * summ

	left := 0.0
	right := 0.0
	//	var res *ProbCase
	var bestInx = -1
	for inx, x := range *probs {
		right += x.probCoef
		if seek > left && seek <= right {
			bestInx = inx
			///res = x
			break
		}
		left += x.probCoef
	}
	if bestInx >= 0 {
		if (*probs)[bestInx].do != nil {
			(*probs)[bestInx].do()
		}
	}

	return bestInx
}

func Prob(probArray []*ProbCase) {
	summ := 0.0
	for _, x := range probArray {
		summ += x.probCoef
	}

	seek := server.Rand() * summ

	left := 0.0
	right := 0.0
	var res *ProbCase
	for _, x := range probArray {
		right += x.probCoef
		if seek > left && seek <= right {
			res = x
			break
		}
		left += x.probCoef
	}

	res.do()
}
