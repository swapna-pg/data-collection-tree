package tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/swapna-pg/golang/data-collection-tree/models"
	"sync"

	//"sync"
	"testing"
)

var (
	mobileDimension    = models.Dimension{Key: models.Device, Value: "Mobile"}
	webDimension       = models.Dimension{Key: models.Device, Value: "Web"}
	inCountryDimension = models.Dimension{Key: models.Country, Value: "IN"}
)

func TestTreeInsertIntoEmptyTree(t *testing.T) {
	tree := NewTree()
	params := models.InsertParams{
		Metrics:    models.Metrics{models.Metric{Key: models.WebReq, Value: 70}, models.Metric{Key: models.TimeSpent, Value: 30}},
		Dimensions: models.Dimensions{inCountryDimension, mobileDimension},
	}

	tree.Insert(params)

	countryNodes := tree.root.children
	assert.Equal(t, len(countryNodes), 1)
	countryNode := countryNodes["IN"]
	deviceDimension := countryNode.children["Mobile"]
	assert.Equal(t, countryNode.Metrics[models.WebReq], 70)
	assert.Equal(t, countryNode.Metrics[models.TimeSpent], 30)
	assert.Equal(t, deviceDimension.Metrics[models.WebReq], 70)
	assert.Equal(t, deviceDimension.Metrics[models.TimeSpent], 30)
}

func TestTreeInsertNewDeviceIntoExistingCountry(t *testing.T) {
	tree := NewTree()
	countryNode := NewNode()
	deviceNode := NewNode()
	mobileMetrics := models.Metrics{models.Metric{Key: models.WebReq, Value: 70}, models.Metric{Key: models.TimeSpent, Value: 30}}
	deviceNode.UpdateData(mobileMetrics)
	countryNode.UpdateData(mobileMetrics)
	tree.root.UpdateData(mobileMetrics)
	countryNode.children["Mobile"] = deviceNode
	tree.root.children["IN"] = countryNode
	webMetrics := models.Metrics{models.Metric{Key: models.WebReq, Value: 50}, models.Metric{Key: models.TimeSpent, Value: 50}}

	tree.Insert(models.InsertParams{webMetrics, models.Dimensions{inCountryDimension, webDimension}})

	assert.Equal(t, len(countryNode.children), 2)
	assert.Equal(t, countryNode.Metrics[models.WebReq], 120)
	assert.Equal(t, countryNode.Metrics[models.TimeSpent], 80)
	assert.Equal(t, countryNode.children["Web"].Metrics[models.WebReq], 50)
	assert.Equal(t, countryNode.children["Web"].Metrics[models.TimeSpent], 50)
}

func TestTreeInsertUpdateExistingDimensions(t *testing.T) {
	tree := NewTree()
	countryNode := NewNode()
	deviceNode := NewNode()
	mobileMetrics := models.Metrics{models.Metric{Key: models.WebReq, Value: 70}, models.Metric{Key: models.TimeSpent, Value: 30}}
	deviceNode.UpdateData(mobileMetrics)
	countryNode.UpdateData(mobileMetrics)
	tree.root.UpdateData(mobileMetrics)
	countryNode.children["Mobile"] = deviceNode
	tree.root.children["IN"] = countryNode

	params := models.InsertParams{
		Metrics:    models.Metrics{models.Metric{Key: models.WebReq, Value: 50}, models.Metric{Key: models.TimeSpent, Value: 50}},
		Dimensions: models.Dimensions{inCountryDimension, mobileDimension},
	}
	tree.Insert(params)

	assert.Equal(t, len(countryNode.children), 1)
	assert.Equal(t, countryNode.Metrics[models.WebReq], 120)
	assert.Equal(t, countryNode.Metrics[models.TimeSpent], 80)
	assert.Equal(t, countryNode.children["Mobile"].Metrics[models.WebReq], 120)
	assert.Equal(t, countryNode.children["Mobile"].Metrics[models.TimeSpent], 80)

}

func TestTreeInsertConcurrently(t *testing.T) {
	tree := NewTree()
	waitGroup := sync.WaitGroup{}
	for i, country := range []string{"IN", "US", "PAK", "CN", "SL", "UK"} {
		for j, device := range []string{"Web", "Mobile", "Tablet"} {
			waitGroup.Add(1)
			dimensions := models.Dimensions{
				models.Dimension{Key: models.Country, Value: country},
				models.Dimension{Key: models.Device, Value: device},
			}
			metrics := models.Metrics{
				models.Metric{Key: models.WebReq, Value: (i + 1) * 10},
				models.Metric{Key: models.TimeSpent, Value: (j + 1) * 20},
			}

			go func(params models.InsertParams) {
				tree.Insert(params)
				waitGroup.Done()
			}(models.InsertParams{metrics, dimensions})
		}
	}
	waitGroup.Wait()
	assert.Equal(t, len(tree.root.children), 6)
	assert.Equal(t, tree.root.Metrics[models.WebReq], 210*3)
	assert.Equal(t, tree.root.Metrics[models.TimeSpent], 120*6)
}

func TestGetMetricsWhenDimensionExists(t *testing.T) {
	tree := NewTree()
	countryNode := NewNode()
	deviceNode := NewNode()
	mobileMetrics := models.Metrics{models.Metric{Key: models.WebReq, Value: 70}, models.Metric{Key: models.TimeSpent, Value: 30}}
	deviceNode.UpdateData(mobileMetrics)
	countryNode.UpdateData(mobileMetrics)
	tree.root.UpdateData(mobileMetrics)
	countryNode.children["Mobile"] = deviceNode
	tree.root.children["IN"] = countryNode

	queryResult, err := tree.GetMetrics(inCountryDimension)

	assert.Nil(t, err)
	assert.Equal(t, queryResult.Dimension, inCountryDimension)
	assert.ElementsMatch(t, queryResult.Metrics, models.Metrics{models.Metric{Key: models.TimeSpent, Value: 30}, models.Metric{Key: models.WebReq, Value: 70}})
}

func TestGetMetricsWhenDimensionNotExists(t *testing.T) {
	tree := NewTree()

	queryResult, err := tree.GetMetrics(inCountryDimension)

	assert.NotNil(t, err)
	assert.Nil(t, queryResult)
}
