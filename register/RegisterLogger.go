package register

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// RegisterLogger 注册日志
func RegisterLogger(log string) {

	logs.SetLogFuncCall(true)
	logs.SetLogger("console")
	// logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()

	logPath := filepath.Join(log, "log.log")

	if _, err := os.Stat(logPath); os.IsNotExist(err) {

		os.MkdirAll(log, 0777)

		if f, err := os.Create(logPath); err == nil {
			f.Close()
			config := make(map[string]interface{}, 1)

			config["filename"] = logPath

			b, _ := json.Marshal(config)

			beego.SetLogger("file", string(b))
		}
	}

	beego.SetLogFuncCall(true)
	beego.BeeLogger.Async()
}
