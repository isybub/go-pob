package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Vilsol/go-pob/cache"
	"github.com/andybalholm/brotli"
	"github.com/dominikbraun/graph"
)

type TreeVersion string

const (
	TreeVersion3_10 = TreeVersion("3_10")
	TreeVersion3_11 = TreeVersion("3_11")
	TreeVersion3_12 = TreeVersion("3_12")
	TreeVersion3_13 = TreeVersion("3_13")
	TreeVersion3_14 = TreeVersion("3_14")
	TreeVersion3_15 = TreeVersion("3_15")
	TreeVersion3_16 = TreeVersion("3_16")
	TreeVersion3_17 = TreeVersion("3_17")
	TreeVersion3_18 = TreeVersion("3_18")
)

const LatestTreeVersion = TreeVersion3_18
const DefaultTreeVersion = TreeVersion3_10

type TreeVersionData struct {
	Display      string
	Num          float64
	URL          string
	cachedTree   *Tree
	rawTree      []byte
	graph        graph.Graph[int64, int64]
	adjacencyMap map[int64]map[int64]graph.Edge[int64]
}

const cdnTreeBase = "https://go-pob-data.pages.dev/data/%s/tree/data.json.br"

func (v *TreeVersionData) Tree() *Tree {
	if v.cachedTree != nil {
		return v.cachedTree
	}

	var outTree Tree
	if err := json.Unmarshal(v.RawTree(), &outTree); err != nil {
		panic(fmt.Errorf("failed to decode file: %w", err))
	}
	v.cachedTree = &outTree

	return v.cachedTree
}

func (v *TreeVersionData) RawTree() []byte {
	if v.rawTree != nil {
		return v.rawTree
	}

	treeURL := fmt.Sprintf(cdnTreeBase, v.Display)
	var compressedTree []byte
	if cache.Disk().Exists(treeURL) {
		var err error
		compressedTree, err = cache.Disk().Get(treeURL)
		if err != nil {
			panic(err)
		}
	} else {
		slog.Debug("fetching", slog.String("url", treeURL))
		req, _ := http.NewRequest(http.MethodGet, treeURL, nil)
		req = req.WithContext(context.Background())
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(fmt.Errorf("failed to fetch url: %s: %w", treeURL, err))
		}
		defer response.Body.Close()

		compressedTree, err = io.ReadAll(response.Body)
		if err != nil {
			panic(fmt.Errorf("failed to read response body: %w", err))
		}

		defer func() {
			_ = cache.Disk().Set(treeURL, compressedTree)
		}()
	}

	unzipStream := brotli.NewReader(bytes.NewReader(compressedTree))

	var err error
	v.rawTree, err = io.ReadAll(unzipStream)
	if err != nil {
		panic(fmt.Errorf("failed to read unzipped data: %w", err))
	}

	return v.rawTree
}

func (v *TreeVersionData) getGraph() (graph.Graph[int64, int64], map[int64]map[int64]graph.Edge[int64]) {
	if v.graph != nil {
		return v.graph, v.adjacencyMap
	}

	g := graph.New(func(v int64) int64 {
		return v
	}, graph.Directed())

	for _, node := range v.Tree().Nodes {
		if node.Skill == nil {
			continue
		}

		_ = g.AddVertex(*node.Skill)
	}

	for _, node := range v.Tree().Nodes {
		if node.Skill == nil {
			continue
		}

		for _, target := range node.Out {
			targetID, err := strconv.ParseInt(target, 10, 64)
			if err != nil {
				continue
			}

			targetNode := v.Tree().Nodes[target]
			if targetNode.ClassStartIndex != nil {
				continue
			}

			if (targetNode.AscendancyName != nil && node.AscendancyName != nil && *targetNode.AscendancyName != *node.AscendancyName) ||
				(targetNode.AscendancyName == nil && node.AscendancyName != nil) ||
				(node.AscendancyName == nil && targetNode.AscendancyName != nil) {
				continue
			}

			_ = g.AddEdge(*node.Skill, targetID)

			// Most edges are bidirectional, but masteries are an exception;
			// you can't use a mastery to travel between 2 notables in a cluster
			if targetNode.IsMastery == nil || !*targetNode.IsMastery {
				_ = g.AddEdge(targetID, *node.Skill)
			}
		}
	}

	v.graph = g

	// We can pre-calculate the adjacency map, as the graph won't change
	// (at least, until we add support for thread of hope/impossible escape/etc)
	v.adjacencyMap, _ = v.graph.AdjacencyMap()

	return v.graph, v.adjacencyMap
}

var TreeVersions = make(map[TreeVersion]*TreeVersionData)
