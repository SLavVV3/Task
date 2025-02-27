package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to   int
	cost int
}

var graph [][]Edge
var n int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	graph = make([][]Edge, n+1)

	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])
		graph[a] = append(graph[a], Edge{b, c})
	}

	for start := 1; start <= n; start++ {
		if found, cycle := findCycle(start); found {
			fmt.Println("YES")
			fmt.Println(strings.Join(cycle, " "))
			return
		}
	}
	fmt.Println("NO")
}

func findCycle(start int) (bool, []string) {
	visited := make([]bool, n+1)
	path := []int{start}
	parents := make(map[int]int)
	dist := make(map[int]int)
	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true

		for _, edge := range graph[u] {
			v, cost := edge.to, edge.cost
			if !visited[v] {
				dist[v] = dist[u] + cost
				parents[v] = u
				visited[v] = true
				queue = append(queue, v)
				path = append(path, v)
			} else if v == start && (dist[u]+cost) < 0 {
				cycle := reconstructPath(u, parents, start)
				return true, cycle
			}
		}
	}
	return false, nil
}

func reconstructPath(u int, parents map[int]int, start int) []string {
	cycle := []int{start}
	current := u
	for current != start {
		cycle = append([]int{current}, cycle...)
		current = parents[current]
	}
	cycle = append([]int{start}, cycle...)

	strCycle := make([]string, len(cycle))
	for i, num := range cycle {
		strCycle[i] = strconv.Itoa(num)
	}
	return strCycle
}
