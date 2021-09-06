package models

import "sort"

func levelOfDimension(dim string) int {
	switch dim {
	case Country:
		return 1
	case Device:
		return 2
	}
	return 0
}

type Metric struct {
	Key   string `json:"key"`
	Value int `json:"val"`
}

type Metrics []Metric

type Dimension struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

type Dimensions []Dimension

type InsertParams struct {
	Metrics    Metrics `json:"dim"`
	Dimensions Dimensions `json:"metrics"`
}

func (p InsertParams) GetSortedDimensions() Dimensions {
	sort.Slice(p.Dimensions, func(i, j int) bool {
		return levelOfDimension(p.Dimensions[i].Key) < levelOfDimension(p.Dimensions[j].Key)
	})
	return p.Dimensions
}

type QueryResult struct {
	Metrics    Metrics
	Dimension Dimension
}
