package graph

import (
	"algorithm/other"
	"errors"
	"math"
	"sort"
)

type Node struct {
	End    int
	Weight int
}

type Graph struct {
	AdjacencyList map[int][]*Node
}

func NewGraph(adjacencyList map[int][]*Node) *Graph {
	return &Graph{
		AdjacencyList: adjacencyList,
	}
}

func dfsTravelImpl(g *Graph, start int, visitInfo map[int]bool, result *[]int) {
	endNodes := g.AdjacencyList[start]
	for _, v := range endNodes {
		if !visitInfo[v.End] {
			*result = append(*result, v.End)
			visitInfo[v.End] = true
			dfsTravelImpl(g, v.End, visitInfo, result)
		}
	}
}

func (g *Graph) DFSTravel() []int {
	visitInfo := make(map[int]bool, len(g.AdjacencyList))
	result := make([]int, 0, len(g.AdjacencyList))
	for k, _ := range g.AdjacencyList {
		if !visitInfo[k] {
			result = append(result, k)
			visitInfo[k] = true
			dfsTravelImpl(g, k, visitInfo, &result)
		}
	}
	return result
}

func isNodeInVisitQueue(visitQueue []int, n int) bool {
	for _, v := range visitQueue {
		if v == n {
			return true
		}
	}

	return false
}

func (g *Graph) BFSTravel() []int {
	visitQueue := make([]int, 0, len(g.AdjacencyList))
	result := make([]int, 0, len(g.AdjacencyList))
	queueStart := 0

	for k, _ := range g.AdjacencyList {
		if !isNodeInVisitQueue(visitQueue, k) {
			visitQueue = append(visitQueue, k)
			for queueStart != len(visitQueue) {
				cur := visitQueue[queueStart]
				result = append(result, cur)
				queueStart += 1
				endNodes := g.AdjacencyList[cur]
				for _, v := range endNodes {
					if !isNodeInVisitQueue(visitQueue, v.End) {
						visitQueue = append(visitQueue, v.End)
					}
				}
			}
		}
	}

	return result
}

// Dijkstra can only handle situation when edge weights are non-negative
func shortestPathDijkstra(g *Graph, start int) (map[int]int, map[int][]int) {
	visit := make(map[int]bool)
	visit[start] = true

	dis := make(map[int]int) // distance that store distance starting from start point

	pre := make(map[int][]int)

	curNode := start
	nextNodes := g.AdjacencyList[curNode]
	for {
		for _, v := range nextNodes {
			if !visit[v.End] {
				if dis[v.End] == 0 || dis[v.End] > dis[curNode]+v.Weight {
					dis[v.End] = dis[curNode] + v.Weight
					length := len(pre[curNode]) + 1
					preNodeList := make([]int, length)
					copy(preNodeList, pre[curNode])
					preNodeList[length-1] = curNode
					pre[v.End] = preNodeList
				}
			}
		}

		minV := 0
		hasNewNode := false
		for k, _ := range g.AdjacencyList {
			if !visit[k] && dis[k] != 0 {
				if minV == 0 || dis[k] < minV {
					minV = dis[k]
					curNode = k
					hasNewNode = true
				}
			}
		}
		if !hasNewNode {
			break
		}
		nextNodes = g.AdjacencyList[curNode]
		visit[curNode] = true
	}

	return dis, pre
}

func (g *Graph) ShortestPathDijkstra(start int, end int) (int, []int) {
	dis, pre := shortestPathDijkstra(g, start)
	return dis[end], pre[end]
}

func shortestPathBellmanFordImpl(g *Graph, start int) (map[int]int, map[int][]int, error) {
	dis := make(map[int]int)

	for k, _ := range g.AdjacencyList {
		dis[k] = math.MaxInt
	}
	dis[start] = 0

	pre := make(map[int][]int)

	for i := 0; i < len(g.AdjacencyList)-1; i++ { // i from 0 to number of vertices - 1
		hasUpdate := false
		for begin, endNodes := range g.AdjacencyList {
			for _, v := range endNodes {
				if dis[begin] != math.MaxInt && dis[v.End] > dis[begin]+v.Weight {
					dis[v.End] = dis[begin] + v.Weight
					hasUpdate = true
					length := len(pre[begin]) + 1
					preNodeList := make([]int, length)
					copy(preNodeList, pre[begin])
					preNodeList[length-1] = begin
					pre[v.End] = preNodeList
				}
			}
		}
		if !hasUpdate {
			break
		}
	}

	// do one more round iteration, if still has updates, then means there exist negative circle
	for begin, endNodes := range g.AdjacencyList {
		for _, v := range endNodes {
			if dis[begin] != math.MaxInt && dis[v.End] > dis[begin]+v.Weight {
				return nil, nil, errors.New("exists negative circle")
			}
		}
	}

	return dis, pre, nil
}

func (g *Graph) ShortestPathBellmanFord(start int, end int) (int, []int, error) {
	dis, pre, err := shortestPathBellmanFordImpl(g, start)
	if err != nil {
		return 0, nil, err
	}
	return dis[end], pre[end], nil
}

func shortestPathSPFAImpl(g *Graph, start int) (map[int]int, map[int][]int, error) {
	dis := make(map[int]int)

	for k, _ := range g.AdjacencyList {
		dis[k] = math.MaxInt
	}
	dis[start] = 0

	pre := make(map[int][]int)

	nodeInQueueInfo := make(map[int]bool)
	nodeEnqueueCount := make(map[int]int)
	nodeQueue := make([]int, 1, len(g.AdjacencyList))
	nodeQueue[0] = start
	nodeInQueueInfo[start] = true
	nodeEnqueueCount[start] = 1
	nodeQueueStart := 0

	for nodeQueueStart != len(nodeQueue) {
		begin := nodeQueue[nodeQueueStart]
		endNodes := g.AdjacencyList[begin]
		nodeQueueStart += 1
		nodeInQueueInfo[begin] = false
		for _, v := range endNodes {
			if dis[begin] != math.MaxInt && dis[v.End] > dis[begin]+v.Weight {
				dis[v.End] = dis[begin] + v.Weight
				length := len(pre[begin]) + 1
				preNodeList := make([]int, length)
				copy(preNodeList, pre[begin])
				preNodeList[length-1] = begin
				pre[v.End] = preNodeList
				if !nodeInQueueInfo[v.End] {
					nodeQueue = append(nodeQueue, v.End)
					nodeEnqueueCount[v.End] += 1
					if nodeEnqueueCount[v.End] == len(g.AdjacencyList) {
						return nil, nil, errors.New("exists negative circle")
					}
				}
			}
		}
	}

	return dis, pre, nil
}

func (g *Graph) ShortestPathSPFA(start int, end int) (int, []int, error) {
	dis, pre, err := shortestPathSPFAImpl(g, start)
	if err != nil {
		return 0, nil, err
	}
	return dis[end], pre[end], nil
}

func (g *Graph) ShortestPathFloyd() map[int]map[int]int {
	dis := make(map[int]map[int]int)
	n := len(g.AdjacencyList)

	for begin, endNodes := range g.AdjacencyList {
		for _, v := range endNodes {
			if dis[begin] == nil {
				dis[begin] = map[int]int{}
			}
			dis[begin][v.End] = v.Weight
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][k] != 0 && dis[k][j] != 0 && (dis[i][j] == 0 || dis[i][k]+dis[k][j] < dis[i][j]) {
					dis[i][j] = dis[i][k] + dis[k][j]
				}
			}
		}
	}

	return dis
}

func (g *Graph) MinimumSpinningTreePrime() map[int]int {
	start := 0
	visit := make(map[int]bool)
	visit[start] = true

	dis := make(map[int]int) // distance that store distance from set tree

	pre := make(map[int]int)

	curNode := start
	nextNodes := g.AdjacencyList[curNode]
	for {
		for _, v := range nextNodes {
			if !visit[v.End] {
				if dis[v.End] == 0 || dis[v.End] > v.Weight {
					dis[v.End] = v.Weight
					pre[v.End] = curNode
				}
			}
		}

		minV := 0
		hasNewNode := false
		for k, _ := range g.AdjacencyList {
			if !visit[k] && dis[k] != 0 {
				if minV == 0 || dis[k] < minV {
					minV = dis[k]
					curNode = k
					hasNewNode = true
				}
			}
		}
		if !hasNewNode {
			break
		}
		nextNodes = g.AdjacencyList[curNode]
		visit[curNode] = true
	}

	return pre
}

type startEndEdges struct {
	Start  int
	End    int
	Weight int
}

func (g *Graph) MinimumSpinningTreeKruskal() (int, error) {
	edges := make([]*startEndEdges, 0, len(g.AdjacencyList))
	vertices := make([]int, 0, len(g.AdjacencyList))
	for start, nodes := range g.AdjacencyList {
		vertices = append(vertices, start)
		for _, v := range nodes {
			edges = append(edges, &startEndEdges{
				Start:  start,
				End:    v.End,
				Weight: v.Weight,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	ufs := other.NewUnionFindSet(vertices)

	selectedEdgeCount := 0
	for _, e := range edges {
		if selectedEdgeCount == len(vertices)-1 {
			break
		}
		hasUnited := ufs.Union(e.Start, e.End)
		if hasUnited {
			selectedEdgeCount += 1
		}
	}
	if selectedEdgeCount != len(vertices)-1 {
		return -1, errors.New("not connected graph")
	}

	return ufs.FindTop(vertices[0]), nil
}

func (g *Graph) TopologicalSort() ([]int, error) {
	inDegree := make(map[int]int)

	for _, node := range g.AdjacencyList {
		for _, v := range node {
			inDegree[v.End] += 1
		}
	}

	result := make([]int, 0, len(g.AdjacencyList))

	for {
		hasFind := false
		for _, d := range inDegree {
			if d == 0 {
				hasFind = true
				result = append(result, d)
				for _, v := range g.AdjacencyList[d] {
					inDegree[v.End] -= 1
				}
			}
		}
		if !hasFind {
			for _, d := range inDegree {
				if d != 0 {
					return nil, errors.New("not a DAG")
				}
			}
			break
		}
	}

	return result, nil
}
