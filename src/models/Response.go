package models

type Response struct {
	ErrorMessage string `xml:"errormessage" json:"errorMessage"`
	Result       string `xml:"result" json:"result"`
}
