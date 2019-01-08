package models

type ReqQuery struct {
	Operation      string
	Table          string
	Join           string
	JoinCondition  string
	Columns        []string
	Conditions     string
	Orderby        string
	Limit          int
	AllowFiltering bool
}
