package controllers

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"strings"

	"github.com/astaxie/beego"
)

var (
	suffix = ".html"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) viewTemplate(value string) (result string) {
	result = value + suffix
	return
}

// JsonResult 响应 json 结果
func (this *BaseController) jsonResult(errCode int, errMsg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["errcode"] = errCode
	jsonData["message"] = errMsg

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//this.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")//解决回退出现json的问题
	//使用gzip原始，json数据会只有原本数据的10分之一左右
	if strings.Contains(strings.ToLower(this.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		this.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		//gzip压缩
		w := gzip.NewWriter(this.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(this.Ctx.ResponseWriter, string(returnJSON))
	}
	this.StopRun()
}
