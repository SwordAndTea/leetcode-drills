package graph

import (
	"algorithm/other"
	"errors"
	"math"
	"sort"
)

type Edge struct {
	Start  int
	End    int
	Weight int
}

type Graph struct {
	AdjacencyList map[int] /*start vertex*/ []*Edge /*end vertices*/
}

func NewGraph(adjacencyList map[int][]*Edge) *Graph {
	return &Graph{
		AdjacencyList: adjacencyList,
	}
}

// DFSTravel performs depth-first search traversal starting from the specified vertex.
// The start parameter is required because Go map iteration order is non-deterministic,
// which would cause inconsistent traversal results if we iterate over AdjacencyList to pick a starting vertex.
//
// The time complexity is O(V+E) and the space complexity is O(H), H is the max depth of the graph,
// in worst case, O(V) a long chain, best case O(log V)
func (g *Graph) DFSTravel(start int) []int {
	visitInfo := make(map[int]bool)
	result := make([]int, 0, len(g.AdjacencyList))

	var dfsRecursive func(vertex int)
	dfsRecursive = func(vertex int) {
		visitInfo[vertex] = true
		result = append(result, vertex)
		for _, v := range g.AdjacencyList[vertex] {
			if !visitInfo[v.End] {
				dfsRecursive(v.End)
			}
		}
	}

	// Start from the specified vertex
	dfsRecursive(start)

	// Visit remaining unvisited vertices (for disconnected graphs)
	for k := range g.AdjacencyList {
		if !visitInfo[k] {
			dfsRecursive(k)
		}
	}

	return result
}

// BFSTravel performs breadth-first search traversal starting from the specified vertex.
// The start parameter is required because Go map iteration order is non-deterministic,
// which would cause inconsistent traversal results if we iterate over AdjacencyList to pick a starting vertex.
//
// The time complexity of BFS is O(V+E) and the space complexity is O(V)
// (worst case: the queue holds almost all vertices, e.g., a wide graph)
func (g *Graph) BFSTravel(start int) []int {
	inQueued := make(map[int]bool, len(g.AdjacencyList)) // whether the node has been put into visit queue before or not
	result := make([]int, 0, len(g.AdjacencyList))

	var bfs func(vertex int)
	bfs = func(vertex int) {
		queue := []int{vertex}
		inQueued[vertex] = true
		i := 0
		for i < len(queue) {
			curVertex := queue[i]
			result = append(result, curVertex)
			i++
			for _, v := range g.AdjacencyList[curVertex] {
				if !inQueued[v.End] {
					queue = append(queue, v.End)
					inQueued[v.End] = true
				}
			}
		}
	}

	// Start from the specified vertex
	bfs(start)

	// Visit remaining unvisited vertices (for disconnected graphs)
	for k := range g.AdjacencyList {
		if !inQueued[k] {
			bfs(k)
		}
	}

	return result
}

// ShortestPathDijkstra can only handle situation when edge weights are non-negative
// the time complexity of Dijkstra is O(V^2 + E) ≈ O(V^2), V is the number of Vertices and E is the number of Edges.
// if we use the heap to help find the vertex with minimal distance, the time complexity can be reduced to O(VlogV + E)
func (g *Graph) ShortestPathDijkstra(start int, end int) (int, []int) {
	visit := make(map[int]bool)
	visit[start] = true

	distance := make(map[int]int) // distance that store distance starting from start point
	distance[start] = 0

	predecessor := make(map[int]int)
	predecessor[start] = start

	curNode := start
	for {
		visit[curNode] = true
		for _, edge := range g.AdjacencyList[curNode] {
			if !visit[edge.End] {
				if dis, ok := distance[edge.End]; !ok || dis > distance[curNode]+edge.Weight {
					distance[edge.End] = distance[curNode] + edge.Weight
					predecessor[edge.End] = curNode
				} // if there are other criteria, we simply add another else if to compare that criteria and of course need to store that info
			}
		}

		// TODO: maybe use a heap to manage this, but we may need to create new struct that contains both vertex and distance
		minV := math.MaxInt
		hasNewNode := false
		for vertex, dis := range distance {
			if !visit[vertex] && dis < minV {
				minV = dis
				curNode = vertex
				hasNewNode = true
			}
		}

		if !hasNewNode {
			break
		}
	}

	// construct path
	if dis, ok := distance[end]; ok {
		path := make([]int, 0, 4)
		path = append(path, end)
		curNode = end
		for predecessor[curNode] != curNode {
			path = append(path, predecessor[curNode])
			curNode = predecessor[curNode]
		}
		// TODO: reverse the path
		return dis, path
	}

	return -1, nil // can not reach to end point
}

// ShortestPathBellmanFord can handle graph with negative edges,
// but can not handle graph with negative circles (sum of the total weight in the circle is negative)
// the time complexity fo Bellman-Ford is O(VE)
func (g *Graph) ShortestPathBellmanFord(start int, end int) (int, []int, error) {
	distance := make(map[int]int)
	for k, _ := range g.AdjacencyList {
		distance[k] = math.MaxInt
	}
	distance[start] = 0

	predecessor := make(map[int]int)

	for i := 0; i < len(g.AdjacencyList)-1; i++ { // i from 0 to number of vertices - 1
		hasUpdate := false
		for begin, endNodes := range g.AdjacencyList {
			for _, v := range endNodes {
				if distance[begin] != math.MaxInt && distance[v.End] > distance[begin]+v.Weight {
					distance[v.End] = distance[begin] + v.Weight
					hasUpdate = true
					predecessor[v.End] = begin
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
			if dis, ok := distance[begin]; ok && distance[v.End] > dis+v.Weight {
				return -1, nil, errors.New("exists negative circle")
			}
		}
	}

	// construct path
	if dis, ok := distance[end]; ok {
		path := make([]int, 0, 4)
		path = append(path, end)
		curNode := end
		for predecessor[curNode] != curNode {
			path = append(path, predecessor[curNode])
			curNode = predecessor[curNode]
		}
		// TODO: reverse the path
		return dis, path, nil
	}

	return -1, nil, errors.New("can not reach end node")
}

// ShortestPathSPFA is an optimized version of Bellman-Ford algorithm, instead iterate all edges in each loop,
// it used a queue to store the node that distance is updated, and we only iterate over that queue
// the time complexity of SPFA is O(E)
func (g *Graph) ShortestPathSPFA(start int, end int) (int, []int, error) {
	distance := make(map[int]int)
	for k, _ := range g.AdjacencyList {
		distance[k] = math.MaxInt
	}
	distance[start] = 0

	predecessor := make(map[int]int)

	nodeInQueueInfo := make(map[int]bool)
	nodeEnqueueCount := make(map[int]int) // the number of time that a node is putted into the queue
	nodeQueue := make([]int, 1, len(g.AdjacencyList))
	nodeQueue[0] = start
	nodeInQueueInfo[start] = true
	nodeEnqueueCount[start] = 1
	nodeQueueStart := 0

	for nodeQueueStart != len(nodeQueue) {
		begin := nodeQueue[nodeQueueStart] // pop the first node in the queue
		endNodes := g.AdjacencyList[begin]
		nodeQueueStart += 1
		nodeInQueueInfo[begin] = false // mark node is not in queue anymore
		for _, v := range endNodes {
			if distance[begin] != math.MaxInt && distance[v.End] > distance[begin]+v.Weight {
				distance[v.End] = distance[begin] + v.Weight
				predecessor[v.End] = begin
				if !nodeInQueueInfo[v.End] {
					nodeQueue = append(nodeQueue, v.End)
					nodeInQueueInfo[v.End] = true // mark node entered the queue
					nodeEnqueueCount[v.End] += 1  // increase the node en-queue count
					if nodeEnqueueCount[v.End] == len(g.AdjacencyList) {
						return -1, nil, errors.New("exists negative circle")
					}
				}
			}
		}
	}

	// construct path
	if dis, ok := distance[end]; ok {
		path := make([]int, 0, 4)
		path = append(path, end)
		curNode := end
		for predecessor[curNode] != curNode {
			path = append(path, predecessor[curNode])
			curNode = predecessor[curNode]
		}
		// TODO: reverse the path
		return dis, path, nil
	}

	return -1, nil, errors.New("can not reach end node")
}

// ShortestPathFloyd can solve all-pairs shortest path
// it can handle edge with negative weight, and it can detect negative circles
func (g *Graph) ShortestPathFloyd() (map[int]map[int]int, error) {
	dis := make(map[int] /*start*/ map[int] /*end*/ int /*distance*/)
	n := len(g.AdjacencyList)

	for begin, edges := range g.AdjacencyList {
		for _, edge := range edges {
			if dis[begin] == nil {
				dis[begin] = map[int]int{}
			}
			dis[begin][edge.End] = edge.Weight
		}
	}

	// Note: we must put the k in the outer-most loop
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][k] != 0 && dis[k][j] != 0 && (dis[i][j] == 0 || dis[i][k]+dis[k][j] < dis[i][j]) {
					dis[i][j] = dis[i][k] + dis[k][j]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if dis[i][i] < 0 {
			return nil, errors.New("exists negative circle")
		}
	}

	return dis, nil
}

// MinimumSpinningTreePrime is very much like Dijkstra, it operates by choose the vertices,
// the time complexity of prime algorithm is O(V^2 + E) ≈ O(V^2)
func (g *Graph) MinimumSpinningTreePrime() (int /*total weight of the spinning tree*/, map[int]int) {
	var start int
	for k, _ := range g.AdjacencyList {
		start = k
		break
	}

	visit := make(map[int]bool) // indicate whether node is selected

	distance := make(map[int]int) // distance from node to the current spinning tree
	distance[start] = 0
	predecessor := make(map[int]int)
	predecessor[start] = start

	totalDistance := 0
	curNode := start
	for {
		visit[curNode] = true
		for _, edge := range g.AdjacencyList[curNode] {
			if !visit[edge.End] {
				// use curNode as bridge, update the distance of node to the current spinning tree
				if dis, ok := distance[edge.End]; !ok || dis > edge.Weight {
					distance[edge.End] = edge.Weight
					predecessor[edge.End] = curNode
				}
			}
		}

		// find the node that is not in the spinning tree and has the minimum distance to the spinning tree
		minV := math.MaxInt
		hasNewNode := false
		for vertex, dis := range distance {
			if !visit[vertex] && dis < minV {
				minV = dis
				curNode = vertex
				hasNewNode = true
			}
		}
		if !hasNewNode {
			break
		}

		totalDistance += minV
	}

	return totalDistance, predecessor
}

// MinimumSpinningTreeKruskal it operate by choose edges, it needs to use union find set to help detect circles
// the time complexity of Kruskal is O(V+E)
func (g *Graph) MinimumSpinningTreeKruskal() (int /*root*/, int /*total weight*/, error) {
	edges := make([]*Edge, 0, len(g.AdjacencyList))  // all edges
	vertices := make([]int, 0, len(g.AdjacencyList)) // all vertices
	for start, nodes := range g.AdjacencyList {
		vertices = append(vertices, start)
		for _, v := range nodes {
			edges = append(edges, &Edge{
				Start:  start,
				End:    v.End,
				Weight: v.Weight,
			})
		}
	}

	totalWeight := 0

	// sort edge in ascending order
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// initialize union find set, very vertex will belong to a set
	ufs := other.NewUnionFindSet(vertices)

	selectedEdgeCount := 0
	for _, e := range edges {
		if selectedEdgeCount == len(vertices)-1 {
			break
		}
		hasUnited := ufs.Union(e.Start, e.End)
		if hasUnited { // has united means the two vertices of current edge is not in a set,
			// furthermore, they will not introduce a circle
			selectedEdgeCount += 1
			totalWeight += e.Weight
		}
	}
	if selectedEdgeCount != len(vertices)-1 {
		return -1, -1, errors.New("not connected graph")
	}

	return ufs.FindTop(vertices[0]), totalWeight, nil
}

// TopologicalSort time complexity of this algorithm is O(V+E)
func (g *Graph) TopologicalSort() ([]int, error) {
	inDegree := make(map[int]int)

	// initialize
	for _, edges := range g.AdjacencyList {
		for _, e := range edges {
			inDegree[e.End] += 1
		}
	}

	result := make([]int, 0, len(g.AdjacencyList))

	for {
		hasFind := false
		for vertex, d := range inDegree {
			if d == 0 {
				hasFind = true
				result = append(result, vertex)
				for _, v := range g.AdjacencyList[vertex] {
					inDegree[v.End] -= 1
				}
				delete(inDegree, d)
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

// CriticalPath solve the longest time need to complex task on AOE (Activity on Edge) graph,
// in an AOE graph, vertex stands for event and weigh on edge stands for the time to complete the activity
// the node in the critical path means the earliest time to start activity on that node
// is equal to the latest time to start activity on that node
//
// For the AOV graph, if we now the time to complete the activity on that node, we can transform the graph
// into a AOE graph by duplicate the vertices, and for each pair of duplicated vertices, assign the weight on that
// edge with the weight of the original vertex. While for the original edges, assign weight with zero.
func (g *Graph) CriticalPath() (criticalPath []int, pathLength int, err error) {
	// Step 1: Get topological order
	topoOrder, err := g.TopologicalSort()
	if err != nil {
		return nil, -1, err
	}

	// Step 2: Calculate earliest start time (ES) for each vertex - forward pass
	// ES[v] = max(ES[u] + weight(u,v)) for all edges (u,v) entering v
	es := make(map[int]int)
	for _, v := range topoOrder {
		es[v] = 0
	}

	for _, u := range topoOrder {
		for _, edge := range g.AdjacencyList[u] {
			if es[u]+edge.Weight > es[edge.End] {
				es[edge.End] = es[u] + edge.Weight
			}
		}
	}

	// Find the maximum ES value (this is the critical path length)
	maxES := 0
	for _, e := range es {
		if e > maxES {
			maxES = e
		}
	}

	// Step 3: Calculate latest start time (LS) for each vertex - backward pass
	// LS[v] = min(LS[u] - weight(v,u)) for all edges (v,u) leaving v
	ls := make(map[int]int)
	for _, v := range topoOrder {
		ls[v] = maxES
	}

	// Traverse in reverse topological order
	for i := len(topoOrder) - 1; i >= 0; i-- {
		u := topoOrder[i]
		for _, edge := range g.AdjacencyList[u] {
			if ls[edge.End]-edge.Weight < ls[u] {
				ls[u] = ls[edge.End] - edge.Weight
			}
		}
	}

	// Step 4: Find vertices where ES == LS (critical vertices)
	criticalPath = make([]int, 0)
	for _, v := range topoOrder { // still, we find the critical path in topological order
		if es[v] == ls[v] {
			criticalPath = append(criticalPath, v)
		}
	}

	return criticalPath, maxES, nil
}
