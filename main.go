package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	from, to, cost int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		cost, _ := strconv.Atoi(parts[2])
		edges[i] = Edge{from, to, cost}
	}

	dist := make([]int, n+1)
	prev := make([]int, n+1)
	for i := range dist {
		dist[i] = 1 << 30
	}
	dist[1] = 0

	var lastUpdated int
	for i := 0; i < n; i++ {
		lastUpdated = -1
		for _, edge := range edges {
			if dist[edge.from] < 1<<30 && dist[edge.to] > dist[edge.from]+edge.cost {
				dist[edge.to] = dist[edge.from] + edge.cost
				prev[edge.to] = edge.from
				lastUpdated = edge.to
			}
		}
	}

	if lastUpdated == -1 {
		fmt.Println("NO")
		return
	}

	// Находим цикл
	cycle := make([]int, 0)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		lastUpdated = prev[lastUpdated]
	}
	current := lastUpdated
	for {
		cycle = append(cycle, current)
		if visited[current] {
			break
		}
		visited[current] = true
		current = prev[current]
	}

	// Реверсируем цикл, чтобы начать с минимального элемента
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		cycle[i], cycle[j] = cycle[j], cycle[i]
	}

	fmt.Println("YES")
	for _, node := range cycle {
		fmt.Print(node, " ")
	}
	fmt.Println()
}
