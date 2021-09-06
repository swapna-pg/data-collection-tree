package tree

import "github.com/swapna-pg/golang/data-collection-tree/models"

type Node struct {
	Metrics  map[string]int
	children map[string]*Node
}

func NewNode() *Node {
	node := Node{
		make(map[string]int),
		make(map[string]*Node),
	}
	return &node
}

func (node *Node) UpdateData(metrics models.Metrics) {
	for _, metric := range metrics {
		node.Metrics[metric.Key] += metric.Value
	}
}
