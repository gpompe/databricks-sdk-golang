package httpmodels

type SparkVersionsRespItem struct {
	Key  string `json:"key,omitempty" url:"key,omitempty"`
	Name string `json:"name,omitempty" url:"name,omitempty"`
}

type SparkVersionsRespItemList struct {
	Versions SparkVersionsRespItem `json:"versions,omitempty" url:"versions,omitempty"`
}
