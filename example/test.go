package main

import (
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zartbot/gograph"
)

const MAXNode = 4096

func main() {
	g := gograph.NewGraph()
	for idx := 0; idx < MAXNode; idx++ {
		g.AddVertex(strconv.Itoa(idx))
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for idx := 0; idx < MAXNode*MAXNode/50; idx++ {
		u := strconv.Itoa(r1.Intn(MAXNode))
		v := strconv.Itoa(r1.Intn(MAXNode))
		w := int64(r1.Intn(MAXNode))
		g.AddEdge(u, v, w)
	}
	runtime.GOMAXPROCS(8)

	start := time.Now()
	dist, prev := gograph.Dijkstra(g, g.VertexMap["0"])
	elapsed := time.Since(start)

	for idx := 0; idx < 100; idx++ {
		DstVertex := strconv.Itoa(r1.Intn(MAXNode))
		logrus.WithFields(
			logrus.Fields{
				"From":     "0",
				"To":       DstVertex,
				"Distance": dist[g.VertexMap[DstVertex]],
			}).Info("Path", g.Path(g.VertexMap[DstVertex], prev))
	}
	logrus.Warn("Elapsed Time:", elapsed)

	start2 := time.Now()
	gograph.DijkstraParallel(g, g.VertexMap["0"])
	elapsed2 := time.Since(start2)
	logrus.Warn("Elapsed Time in Parallel:", elapsed2)
}
