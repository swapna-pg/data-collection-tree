package tree

import (
	"errors"
	"fmt"
	"github.com/swapna-pg/golang/data-collection-tree/models"
	"sync"
)

type Tree struct {
	root  *Node
	mutex sync.Mutex
}

func (tree *Tree) Insert(info models.InsertParams) {
	tree.incrementMetrics(tree.root, info.Metrics)

	dimensions := info.GetSortedDimensions()
	currentNode := tree.root

	for _, dimension := range dimensions {
		childNode, ok := currentNode.children[dimension.Value]
		if !ok {
			childNode = NewNode()
			currentNode.children[dimension.Value] = childNode
		}
		tree.incrementMetrics(childNode, info.Metrics)
		currentNode = childNode
	}
}

func (tree *Tree) incrementMetrics(node *Node, metrics models.Metrics) {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	node.UpdateData(metrics)
}

func (tree *Tree) GetMetrics(dimension models.Dimension) (*models.QueryResult, error) {
	if countryNode, ok := tree.root.children[dimension.Value]; ok {
		metrics := models.Metrics{}
		for key, value := range countryNode.Metrics {
			metrics = append(metrics, models.Metric{key, value})
		}
		return &models.QueryResult{metrics, dimension}, nil
	}
	return nil, errors.New(fmt.Sprintf("dimension %v not found", dimension.Value))
}

func NewTree() *Tree {
	return &Tree{root: NewNode(), mutex: sync.Mutex{}}
}
