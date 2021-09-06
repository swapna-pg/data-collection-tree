package handler

import (
	"encoding/json"
	"fmt"
	"github.com/swapna-pg/golang/data-collection-tree/models"
	"github.com/swapna-pg/golang/data-collection-tree/tree"
	"net/http"
)

type DataCollector struct {
	tree *tree.Tree
}

func NewDataCollector() *DataCollector {
	return &DataCollector{tree.NewTree()}
}

func (data *DataCollector) Insert(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var reqBody models.InsertRequest
	err := json.NewDecoder(request.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("error while parsing insert request: %+v\n", err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		data.tree.Insert(reqBody.GetInsertDataParams())
		message, _ := json.Marshal(map[string]string{"msg": "Inserted successfully"})
		response.WriteHeader(http.StatusBadRequest)
		response.Write(message)
		response.WriteHeader(http.StatusOK)
	}
}

func (data *DataCollector) Query(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var reqBody models.QueryRequest
	err := json.NewDecoder(request.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("error while parsing query request: %+v\n", err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		country := reqBody.GetCountry()
		result, err := data.tree.GetMetrics(models.Dimension{Key: models.Country, Value: country})
		if err != nil {
			errorMessage, _ := json.Marshal(map[string]string{"error": err.Error()})
			response.WriteHeader(http.StatusBadRequest)
			response.Write(errorMessage)
		} else {
			resBody := models.GetQueryResponse(result)
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(resBody)
		}
	}
}
