package graph

import (
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {

	n := len(edges) + 1

	tree := make([][]int, n)

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	bob_steps := bobSteps(bob, tree, 0)

	// fmt.Println(bob_steps)

	return aliceProfit(tree, amount, bob_steps)

}

func aliceProfit(adj [][]int, amount []int, bob_steps map[int]int) int {

	max_profit := math.MinInt

	alice_stat_node := 0

	q := [][4]int{{alice_stat_node, 0, 0, 0}} // node - step - profit - longest way

	visited := make([]bool, len(adj))

	visited[alice_stat_node] = true

	for len(q) > 0 {
		current_q_item := q[0]

		node := current_q_item[0]
		step := current_q_item[1]
		profit := current_q_item[2]
		longest_way := current_q_item[3]

		q = q[1:]

		bob_step, ok := bob_steps[node]
		new_profit := profit
		if !ok {
			new_profit += amount[node]
		} else {
			if step == bob_step {
				new_profit += amount[node] / 2
			} else if step > bob_step {
				// bob already passed this node
			} else {
				new_profit += amount[node]
			}
		}

		is_end_leaf := true
		for _, child := range adj[node] {
			if !visited[child] {
				visited[child] = true
				q = append(q, [4]int{child, step + 1, new_profit, longest_way + 1})
				is_end_leaf = false
			}
		}
		if is_end_leaf {
			max_profit = max(max_profit, new_profit)
		}
	}

	return max_profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dfs
// adj[i] relations of i
func bobSteps(bob int, adj [][]int, target int) map[int]int {

	step := 0
	bob_steps := make(map[int]int) // node - step

	visited := make([]bool, len(adj))

	stack := [][2]int{{bob, step}} // node - step

	visited[bob] = true

	found := false

	parents := make([]int, len(adj))

	for len(stack) > 0 && !found {

		node := stack[len(stack)-1][0]
		step = stack[len(stack)-1][1]

		if _, ok := bob_steps[node]; !ok {
			bob_steps[node] = step
		}

		stack = stack[:len(stack)-1]

		if node == target {
			found = true
			break
		}

		for _, child := range adj[node] {

			if !visited[child] {
				visited[child] = true
				stack = append(stack, [2]int{child, step + 1})
				parents[child] = node
			}
		}
	}

	// fmt.Println("parent", parents)

	if !found {
		return bob_steps
	}

	// backtrack to find the path
	current := target
	path := []int{}
	for current != bob {
		path = append([]int{current}, path...)
		current = parents[current]
	}

	filtered_bob_steps := make(map[int]int)
	// create map of steps only for the path
	for _, node := range path {
		if v, ok := bob_steps[node]; ok {
			filtered_bob_steps[node] = v
		}
	}

	filtered_bob_steps[bob] = 0

	return filtered_bob_steps
}
