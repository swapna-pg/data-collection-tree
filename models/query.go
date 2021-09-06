package models

type QueryRequest struct {
	Dim Dimensions
}

func (q QueryRequest) GetCountry() string {
	for _, dim := range q.Dim {
		if dim.Key == "country" {
			return dim.Value
		}
	}
	return ""
}

type QueryAPIResponse struct {
	Dim     Dimensions `json:"dim"`
	Metrics Metrics    `json:"metrics"`
}

func GetQueryResponse(result *QueryResult) QueryAPIResponse {
	return QueryAPIResponse{Dimensions{result.Dimension}, result.Metrics}
}
