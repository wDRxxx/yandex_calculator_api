package models

type Input struct {
	Expression string `json:"expression"`
}

type SuccessOutput struct {
	Result float64 `json:"result"`
}

type ErrorOutput struct {
	Error string `json:"error"`
}
