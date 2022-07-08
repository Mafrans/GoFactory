package gameobject

import (
	"mafrans/gorogue/util"
	"math"

	"golang.org/x/exp/slices"
)

func PathFind(origin, destination [2]int) [][2]int {
	openQueue := util.PriorityQueue[[2]int]{}
	parents := make(map[[2]int][2]int)
	stepScores := make(map[[2]int]float64)
	totalScores := make(map[[2]int]float64)

	openQueue.Push(origin, 0)
	stepScores[origin] = 0
	totalScores[origin] = 0

	for len(openQueue) > 0 {
		current, i := openQueue.Lowest()
		if current.Value == destination {
			return retrace(parents, current.Value)
		}

		openQueue = slices.Delete(openQueue, i, i+1)
		for _, neighbor := range getNeighbors(current.Value) {
			if stepScore, ok := stepScores[current.Value]; ok {
				stepScore += neighbor.weight
				neighborScore, ok := stepScores[neighbor.position]

				if !ok || stepScore < neighborScore {
					totalScore := stepScore + distance(current.Value, neighbor.position)

					parents[neighbor.position] = current.Value
					stepScores[neighbor.position] = stepScore
					totalScores[neighbor.position] = totalScore

					notInOpenQueue := true
					for _, item := range openQueue {
						if item.Value == neighbor.position {
							notInOpenQueue = false
						}
					}

					if notInOpenQueue {
						openQueue.Push(neighbor.position, totalScore)
					}
				}
			}
		}
	}

	return nil
}

func retrace(parents map[[2]int][2]int, node [2]int) [][2]int {
	result := [][2]int{node}
	for {
		if parent, ok := parents[node]; ok {
			node = parent
			result = append(result, parent)
		} else {
			return result
		}
	}
}

func distance(from, to [2]int) float64 {
	return math.Sqrt(
		math.Pow(float64(to[0]-from[0]), 2) +
			math.Pow(float64(to[0]-from[0]), 2),
	)
}

type neighbor struct {
	position [2]int
	weight   float64
}

func getNeighbors(pos [2]int) []neighbor {
	return []neighbor{
		{position: [2]int{pos[0], pos[1] - 1}, weight: 1},
		{position: [2]int{pos[0] + 1, pos[1] - 1}, weight: math.Sqrt2},
		{position: [2]int{pos[0] + 1, pos[1]}, weight: 1},
		{position: [2]int{pos[0] + 1, pos[1] + 1}, weight: math.Sqrt2},
		{position: [2]int{pos[0], pos[1] + 1}, weight: 1},
		{position: [2]int{pos[0] - 1, pos[1] + 1}, weight: math.Sqrt2},
		{position: [2]int{pos[0] - 1, pos[1]}, weight: 1},
		{position: [2]int{pos[0] - 1, pos[1] - 1}, weight: math.Sqrt2},
	}
}
