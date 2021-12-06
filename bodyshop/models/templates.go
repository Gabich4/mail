package models

type TemplateParameter struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

var Templates = map[string]string{ // TODO: move in db
	"testTemplateId1": "<div><h1>{{ .testParamName1}}</h1><p>{{ .testParamName2}}</p></div>",
	"testTemplateId2": "<div><h2>{{ .testParamName2}}</h2><p>Default text</p></div>",
}
