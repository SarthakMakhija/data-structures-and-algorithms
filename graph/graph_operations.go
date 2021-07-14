package graph

type ShortestPath struct {
	path   []int
	weight int
}

func ShortestPathFrom(source int, destination int, in [3][]int) *ShortestPath {
	type ConnectingVertex struct {
		vertex int
		weight int
	}

	connectingVerticesFrom := func(vertex int) []ConnectingVertex {
		var connectingVertices []ConnectingVertex
		for index := 0; index < len(in[0]); index++ {
			if in[0][index] == vertex {
				connectingVertices = append(connectingVertices,
					ConnectingVertex{
						vertex: in[1][index],
						weight: in[2][index],
					},
				)
			}
		}
		return connectingVertices
	}

	var shortestPath *ShortestPath

	var shortestPathInner func(int, []int, int)
	shortestPathInner = func(currentVertex int, pathParts []int, weight int) {
		if currentVertex == destination {
			if shortestPath == nil {
				shortestPath = &ShortestPath{
					path:   pathParts,
					weight: weight,
				}
			} else {
				if weight < shortestPath.weight {
					shortestPath.path = pathParts
					shortestPath.weight = weight
				}
			}
			return
		} else {
			connectingVertices := connectingVerticesFrom(currentVertex)
			for _, connectingVertex := range connectingVertices {
				shortestPathInner(
					connectingVertex.vertex,
					append(pathParts, connectingVertex.vertex),
					weight+connectingVertex.weight,
				)
			}
		}
	}
	shortestPathInner(source, []int{source}, 0)
	return shortestPath
}
