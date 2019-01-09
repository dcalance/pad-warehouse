package utils

import (
	"strconv"
	"strings"

	"../models"
)

func CreateCrudQueryString(data models.ReqCRUD) string {
	resString := data.Operation
	switch data.Operation {
	case "select":
		resString += " json " + strings.Join(data.Columns, ", ") + " FROM " + data.Table
		if data.Join != "" {
			resString += data.Join + " ON " + data.JoinCondition
		}
		if data.Conditions != "" {
			resString += " WHERE " + data.Conditions
		}
		if data.Orderby != "" {
			resString += " ORDER BY " + data.Orderby
		}
		if data.Limit != 0 {
			resString += " LIMIT " + strconv.Itoa(data.Limit)
		}
		if data.AllowFiltering {
			resString += " ALLOW FILTERING"
		}

	case "update":
		resString += data.Table
		if data.TTL > 0 {
			resString += " USING TTL" + strconv.Itoa(data.TTL)
			if data.Timestamp > 0 {
				resString += " AND "
			}
		}
		if data.Timestamp > 0 {
			resString += " USING TIMESTAMP " + strconv.Itoa(data.Timestamp)
		}
		resString += " SET " + data.Set + " WHERE " + data.Conditions
		if data.ExistsCondition != "" {
			resString += "IF" + data.ExistsCondition
		}

	case "insert":
		resString += " INTO " + data.Table + " (" + strings.Join(data.ColumnsList, ", ") + ") " + " VALUES "
		for _, element := range data.InsertValues {
			resString += " (" + strings.Join(element, ", ") + "),"
		}
		resString = resString[:len(resString)-1]
		if data.IfNotExists {
			resString += " IF NOT EXISTS "
		}
		if data.TTL > 0 {
			resString += "USING TTL " + strconv.Itoa(data.TTL)
			if data.Timestamp > 0 {
				resString += " AND TIMESTAMP" + strconv.Itoa(data.Timestamp)
			}
		}
	case "delete":
		resString += strings.Join(data.ColumnsList, ", ")
		resString += " FROM " + data.Table
		if data.Timestamp > 0 {
			resString += " USING TIMESTAMP " + strconv.Itoa(data.Timestamp)
		}
		if data.Conditions != "" {
			resString += " WHERE " + data.Conditions
		}
		if data.ExistsCondition != "" {
			resString += "IF" + data.ExistsCondition
		}
	}
	resString += ";"
	return resString
}
