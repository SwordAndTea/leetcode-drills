package _401_410

import "container/heap"

type Cell struct {
	height, x, y int
}

type MinHeap []*Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[:n-1]
	return x
}

// leetcode problem No. 407

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])
	visited := make([][]bool, m)

	var minHeap MinHeap
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		heap.Push(&minHeap, &Cell{heightMap[i][0], i, 0})
		visited[i][0] = true
		heap.Push(&minHeap, &Cell{heightMap[i][n-1], i, n - 1})
		visited[i][n-1] = true
	}

	for j := 1; j < n-1; j++ {
		heap.Push(&minHeap, &Cell{heightMap[0][j], 0, j})
		heap.Push(&minHeap, &Cell{heightMap[m-1][j], m - 1, j})
		visited[0][j], visited[m-1][j] = true, true
	}

	ans := 0
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for minHeap.Len() > 0 {
		cell := heap.Pop(&minHeap).(*Cell)
		for _, dir := range directions {
			nx, ny := cell.x+dir[0], cell.y+dir[1]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny] {
				ans += max(0, cell.height-heightMap[nx][ny])
				heap.Push(&minHeap, &Cell{max(cell.height, heightMap[nx][ny]), nx, ny})
				visited[nx][ny] = true
			}
		}
	}

	return ans
}
