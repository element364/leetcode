package main

import (
	"fmt"
)

type Mutation struct {
	Gene  string
	Steps int
}

func minMutation(startGene string, endGene string, bank []string) int {
	bankContains := map[string]bool{}
	for _, bankItem := range bank {
		bankContains[bankItem] = true
	}
	genes := []byte{'A', 'C', 'G', 'T'}
	visited := map[string]bool{startGene: true}
	queue := []Mutation{
		{
			Gene:  startGene,
			Steps: 0,
		},
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i := 0; i < len(current.Gene); i++ {
			for _, g := range genes {
				if current.Gene[i] != g {
					next := current.Gene[:i] + string(g) + current.Gene[i+1:]
					if _, contains := bankContains[next]; !contains {
						continue
					}
					if next == endGene {
						return current.Steps + 1
					}

					if _, ok := visited[next]; !ok {
						visited[next] = true
						queue = append(queue, Mutation{
							Gene:  next,
							Steps: current.Steps + 1,
						})
					}
				}
			}
		}
	}
	return -1
}

func main() {
	startGene := "AACCGGTT"
	endGene := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(startGene, endGene, bank))

	startGene = "AACCGGTT"
	endGene = "AAACGGTA"
	bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	fmt.Println(minMutation(startGene, endGene, bank))

	startGene = "AACCGGTT"
	endGene = "AACCGGTA"
	bank = []string{}
	fmt.Println(minMutation(startGene, endGene, bank))

	startGene = "AGCAAAAA"
	endGene = "GACAAAAA"
	bank = []string{"AGTAAAAA", "GGTAAAAA", "GATAAAAA", "GACAAAAA"}
	fmt.Println(minMutation(startGene, endGene, bank))

	startGene = "AAAACCCC"
	endGene = "CCCCCCCC"
	bank = []string{"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC"}
	fmt.Println(minMutation(startGene, endGene, bank))
}
