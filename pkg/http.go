package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func HTTPExecute(step *Step) {
	Logger.Info(step.Request.String())

	var payload, _ = json.Marshal(step.Request.JSON)
	var resp, err = http.Post(step.Request.Host+step.Request.Path, "application/json", bytes.NewReader(payload))
	if err != nil {
		Logger.Error(err.Error())
	}
	// 读取响应的内容
	body, _ := io.ReadAll(resp.Body)
	Logger.Info(string(body))
}
