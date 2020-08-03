package mysql

import (
	"database/sql"
	"strings"

	"github.com/renjingneng/a_simple_go_project/core/container"
)

type Base struct {
	Tablename string
	Dbname    string
	DbptrW    *sql.DB
	DbptrR    *sql.DB
}

func NewBase(Dbname string) *Base {
	return &Base{
		Dbname: Dbname,
		DbptrW: container.GetEntityFromMysqlContainer(Dbname, "W"),
		DbptrR: container.GetEntityFromMysqlContainer(Dbname, "R"),
	}
}

func (b *Base) SetTablename(tablename string) {
	b.Tablename = tablename
}

func (b *Base) FetchRow(fields string, condition map[string]string) map[string]interface{} {
	querySQL, values := b.BuildQuerySQL(condition, map[string]string{"limit": "0,1", "fields": fields})
	stmt, err := b.DbptrR.Prepare(querySQL)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(values...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	result := b.FetchResult(rows)
	if len(result) > 0 {
		return result[0]
	} else {
		return nil
	}

}

func (b *Base) BuildQuerySQL(condition map[string]string, other map[string]string) (string, []interface{}) {
	if _, ok := other["fields"]; !ok {
		other["fields"] = "*"
	}
	if _, ok := other["order"]; !ok {
		other["order"] = " ORDER BY id desc"
	} else {
		other["order"] = " ORDER BY " + other["order"]
	}
	if _, ok := other["group"]; !ok {
		other["group"] = ""
	} else {
		other["group"] = " GROUP BY " + other["group"]
	}
	if _, ok := other["limit"]; !ok {
		other["limit"] = ""
	} else {
		other["limit"] = " LIMIT " + other["limit"]
	}
	var where, values = b.BuildCondition(condition)

	querySQL := "SELECT " + other["fields"] + " FROM " + b.Tablename + " WHERE " + where + other["group"] + other["order"] + other["limit"]
	return querySQL, values
}

func (b *Base) BuildCondition(condition map[string]string) (string, []interface{}) {
	var where string = " 1"
	var values []interface{}
	for k, v := range condition {
		str := strings.Split(k, " ")
		if len(str) == 1 {
			where += " AND " + k + "=" + "?"
		} else if str[1] == "IN" {
			where += " AND " + str[0] + " " + str[1] + "(?)"
		} else {
			where += " AND " + str[0] + " " + str[1] + "?"
		}
		values = append(values, v)
	}
	return where, values
}

func (b *Base) FetchResult(rows *sql.Rows) []map[string]interface{} {

	var result []map[string]interface{}
	//获取记录列
	if columns, err := rows.Columns(); err != nil {
		return nil
	} else {
		//拼接记录Map
		values := make([]sql.RawBytes, len(columns))
		scans := make([]interface{}, len(columns))
		for i := range values {
			scans[i] = &values[i]
		}
		for rows.Next() {
			_ = rows.Scan(scans...)
			each := map[string]interface{}{}
			for i, col := range values {
				each[columns[i]] = string(col)
			}
			result = append(result, each)
		}
		if err := rows.Err(); err != nil {
			return nil
		}

	}
	return result
}
