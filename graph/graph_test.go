package graph

import (
	"slices"
	"testing"
)

func TestGraph_DFSTravel(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End: 1,
		},
		{
			End: 2,
		},
	}

	al[1] = []*Edge{
		{
			End: 3,
		},
		{
			End: 4,
		},
	}

	al[2] = []*Edge{
		{
			End: 1,
		},
		{
			End: 4,
		},
	}

	al[4] = []*Edge{
		{
			End: 3,
		},
		{
			End: 5,
		},
	}

	al[5] = []*Edge{
		{
			End: 3,
		},
	}

	g := NewGraph(al)
	r := g.DFSTravel()
	if !slices.Equal(r, []int{0, 1, 3, 4, 5, 2}) {
		t.Errorf("dfs travel not correct, result: %+v", r)
	}
}

func TestGraph_BFSTravel(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End: 1,
		},
		{
			End: 2,
		},
	}

	al[1] = []*Edge{
		{
			End: 3,
		},
		{
			End: 4,
		},
	}

	al[2] = []*Edge{
		{
			End: 1,
		},
		{
			End: 4,
		},
	}

	al[4] = []*Edge{
		{
			End: 3,
		},
		{
			End: 5,
		},
	}

	al[5] = []*Edge{
		{
			End: 3,
		},
	}

	g := NewGraph(al)
	r := g.BFSTravel()

	if !slices.Equal(r, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("bfs travel not correct, result: %+v", r)
	}
}

func TestGraph_ShortestPathDijkstra(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End:    1,
			Weight: 1,
		},
		{
			End:    3,
			Weight: 4,
		},
		{
			End:    4,
			Weight: 4,
		},
	}

	al[1] = []*Edge{
		{
			End:    3,
			Weight: 2,
		},
	}

	al[2] = []*Edge{
		{
			End:    5,
			Weight: 1,
		},
	}

	al[3] = []*Edge{
		{
			End:    2,
			Weight: 2,
		},
		{
			End:    4,
			Weight: 3,
		},
	}

	al[4] = []*Edge{
		{
			End:    5,
			Weight: 3,
		},
	}

	g := NewGraph(al)

	dis, pre := g.ShortestPathDijkstra(0, 5)
	t.Logf("dis is %d and pre is %+v", dis, pre)
}

func TestGraph_ShortestPathBellmanFord(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End:    1,
			Weight: 1,
		},
		{
			End:    3,
			Weight: 4,
		},
		{
			End:    4,
			Weight: 4,
		},
	}

	al[1] = []*Edge{
		{
			End:    3,
			Weight: 2,
		},
	}

	al[2] = []*Edge{
		{
			End:    5,
			Weight: 1,
		},
	}

	al[3] = []*Edge{
		{
			End:    2,
			Weight: 2,
		},
		{
			End:    4,
			Weight: 3,
		},
	}

	al[4] = []*Edge{
		{
			End:    5,
			Weight: 3,
		},
	}

	al[5] = nil

	g := NewGraph(al)

	dis, pre, err := g.ShortestPathBellmanFord(0, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("dis is %d and pre is %+v", dis, pre)
}

func TestGraph_ShortestPathSPFA(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End:    1,
			Weight: 1,
		},
		{
			End:    3,
			Weight: 4,
		},
		{
			End:    4,
			Weight: 4,
		},
	}

	al[1] = []*Edge{
		{
			End:    3,
			Weight: 2,
		},
	}

	al[2] = []*Edge{
		{
			End:    5,
			Weight: 1,
		},
	}

	al[3] = []*Edge{
		{
			End:    2,
			Weight: 2,
		},
		{
			End:    4,
			Weight: 3,
		},
	}

	al[4] = []*Edge{
		{
			End:    5,
			Weight: 3,
		},
	}

	al[5] = nil

	g := NewGraph(al)

	dis, pre, err := g.ShortestPathSPFA(0, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("dis is %d and pre is %+v", dis, pre)
}

func TestGraph_ShortestPathFloyd(t *testing.T) {
	al := make(map[int][]*Edge)
	al[0] = []*Edge{
		{
			End:    1,
			Weight: 1,
		},
		{
			End:    3,
			Weight: 4,
		},
		{
			End:    4,
			Weight: 4,
		},
	}

	al[1] = []*Edge{
		{
			End:    3,
			Weight: 2,
		},
	}

	al[2] = []*Edge{
		{
			End:    5,
			Weight: 1,
		},
	}

	al[3] = []*Edge{
		{
			End:    2,
			Weight: 2,
		},
		{
			End:    4,
			Weight: 3,
		},
	}

	al[4] = []*Edge{
		{
			End:    5,
			Weight: 3,
		},
	}

	al[5] = nil

	g := NewGraph(al)

	dis, _ := g.ShortestPathFloyd()
	t.Logf("dis for 0 to 5 is %d", dis[0][5])
}
