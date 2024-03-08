package pkg

type Request struct {
	Host   string
	Path   string
	Method string
	JSON   map[string]any
}

type Step struct {
	Name     string              `json:"name" yaml:"name"`
	Action   ActionType          `json:"action" yaml:"action"`
	Request  Request             `json:"request" yaml:"request"`
	Variable []map[string]any    `json:"variable,omitempty" yaml:"variable,omitempty"`
	Export   []map[string]string `json:"export,omitempty" yaml:"export,omitempty"`
	Check    []map[string][]any  `json:"check,omitempty" yaml:"check,omitempty"`
}
