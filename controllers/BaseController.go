package controllers

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
)

type BaseController struct {
	beego.Controller
}

func viewTemplate(value string) string {
	view := value + utils.Suffix
	return view
}

// Prepare 预处理.
func (this *BaseController) Prepare() {
	// form
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	// ajax
	this.Data["xsrf_token"] = this.XSRFToken()

	// session
	this.Data["account"] = this.GetAccountSession()
}

func (this *BaseController) GetAccountSession() models.Account {
	data, ok := this.GetSession(utils.SESSION_USER).(models.Account)
	logs.Debug(">> session:", data)
	if !ok {
		errors.New("session nil")
	}
	return data

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
