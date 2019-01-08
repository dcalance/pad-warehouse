package models

type Response struct {
	ErrorCode    int
	ErrorMessage string
	Result       []map[string]interface{}
}
