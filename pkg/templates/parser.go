package templates

import (
	"bytes"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"text/template"
	"time"
	"watchAlert/internal/global"
	"watchAlert/internal/models"
	"watchAlert/pkg/tools"
)

var tmpl *template.Template

// ParserTemplate 处理告警推送的消息模版
func ParserTemplate(defineName string, alert models.AlertCurEvent, templateStr string) string {

	firstTriggerTime := time.Unix(alert.FirstTriggerTime, 0).Format(global.Layout)
	recoverTime := time.Unix(alert.RecoverTime, 0).Format(global.Layout)
	alert.FirstTriggerTimeFormat = firstTriggerTime
	alert.RecoverTimeFormat = recoverTime

	tmpl = template.Must(template.New("tmpl").Parse(templateStr))

	var (
		buf bytes.Buffer
		err error
	)

	if defineName == "Card" {
		err = tmpl.Execute(&buf, alert)
		// 当前告警的 json 反序列化成 map 对象, 用于解析报警事件详情中的 ${xx} 变量
		data := tools.ConvertEventToMap(alert)
		return tools.ParserVariables(buf.String(), data)
	}

	err = tmpl.ExecuteTemplate(&buf, defineName, alert)
	if err != nil {
		logc.Error(context.Background(), "告警模版执行失败 ->", err.Error())
		return ""
	}

	// 前面只会渲染出模版框架, 下面来渲染告警数据内容
	if defineName == "Event" {
		data := tools.ConvertEventToMap(alert)
		return tools.ParserVariables(buf.String(), data)
	}

	return buf.String()

}
