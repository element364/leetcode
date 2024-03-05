package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := map[string]float64{}
	variables := map[string]bool{}
	for i := 0; i < len(equations); i++ {
		a := equations[i][0]
		b := equations[i][1]
		// fmt.Println(va, "/", vb, "=", values[i])
		variables[a] = true
		variables[b] = true
		g[a+"/"+b] = values[i]
		g[b+"/"+a] = 1 / values[i]
	}

	// fmt.Println("G before", g)
	for k := range variables {
		for i := range variables {
			for j := range variables {
				d1, e1 := g[i+"/"+k]
				d2, e2 := g[k+"/"+j]
				if e1 && e2 {
					g[i+"/"+j] = d1 * d2
				}
			}
		}
	}
	// fmt.Println("G after", g)

	result := make([]float64, len(queries))
	for i, query := range queries {
		va := query[0]
		vb := query[1]
		d, e := g[va+"/"+vb]
		if e {
			result[i] = d
		} else {
			result[i] = -1
		}
	}

	return result
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println("Example 1", calcEquation(equations, values, queries))

	equations = [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values = []float64{1.5, 2.5, 5.0}
	queries = [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	fmt.Println("Example 2", calcEquation(equations, values, queries))

	equations = [][]string{{"a", "b"}}
	values = []float64{0.5}
	queries = [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}
	fmt.Println("Example 3", calcEquation(equations, values, queries))

	equations = [][]string{{"a", "c"}, {"b", "e"}, {"c", "d"}, {"e", "d"}}
	values = []float64{2.0, 3.0, 0.5, 5.0}
	queries = [][]string{{"a", "b"}}
	fmt.Println("Example 4", calcEquation(equations, values, queries))
}
