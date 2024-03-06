package models

type RuleTemplateGroup struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Description string `json:"description"`
}

type RuleTemplate struct {
	RuleGroupName     string            `json:"ruleGroupName"`
	RuleName          string            `json:"ruleName"`
	DatasourceType    string            `json:"datasourceType"`
	PrometheusConfig  PrometheusConfig  `json:"prometheusConfig" gorm:"prometheusConfig;serializer:json"`
	AliCloudSLSConfig AliCloudSLSConfig `json:"alicloudSLSConfig" gorm:"alicloudSLSConfig;serializer:json"`
	LokiConfig        LokiConfig        `json:"lokiConfig" gorm:"lokiConfig;serializer:json"`
	EvalInterval      int64             `json:"evalInterval"`
	ForDuration       int64             `json:"forDuration"`
	Annotations       string            `json:"annotations"`
}
