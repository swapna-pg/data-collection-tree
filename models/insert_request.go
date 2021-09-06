package models

type InsertRequest struct {
	Dim     Dimensions `json:"dim"`
	Metrics Metrics    `json:"metrics"`
}

func (r InsertRequest) GetInsertDataParams() InsertParams {
	params := InsertParams{}
	for _, dim := range r.Dim {
		params.Dimensions = append(params.Dimensions, Dimension{dim.Key, dim.Value})
	}

	for _, metric := range r.Metrics {
		params.Metrics = append(params.Metrics, Metric{metric.Key, metric.Value})
	}

	return params
}
