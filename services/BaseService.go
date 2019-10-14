package services

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
)

type BaseService struct{}

func getOrm() orm.Ormer {
	o := orm.NewOrm()
	return o
}

func TableName(model interface{}) string {
	val := reflect.ValueOf(model)
	table := snakeString(reflect.Indirect(val).Type().Name())
	return utils.DatabasePrefix + table
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func getPageSQLString(sql string, page *models.Page) (querySQL string, conutSQL string) {
	conutSQL = "SELECT count(1) FROM (" + sql + ") t_count"

	// if utils.Driver == "mysql" {
	pageSize := page.CurrentPage
	offset := (page.CurrentPage - 1) * page.ShowCount
	querySQL = sql + " LIMIT " + fmt.Sprintf("%d,%d", offset, pageSize)
	// }

	return
}
