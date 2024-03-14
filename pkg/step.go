package pkg

import (
	"fmt"
)

type Step struct {
	Case     Case
	Name     string              `json:"name" yaml:"name"`
	Action   ActionType          `json:"action" yaml:"action"`
	Request  Request             `json:"request" yaml:"request"`
	Variable []map[string]any    `json:"variable,omitempty" yaml:"variable,omitempty"`
	Export   []map[string]string `json:"export,omitempty" yaml:"export,omitempty"`
	Check    []map[string][]any  `json:"check,omitempty" yaml:"check,omitempty"`
	Setup    []func()
	Teardown []func()
}

func (s *Step) Before() {
	s.Case.Cache.Set("key", "value")
	for _, function := range s.Setup {
		function()
	}
}

func (s *Step) After() {
	Logger.Info(fmt.Sprintf("%s", s.Case.Cache.Get("key")))
	for _, function := range s.Teardown {
		function()
	}
}

func (s *Step) Run(tase Case) {
	s.Case = tase
	s.Before()
	s.Execute()
	s.After()
}

func (s *Step) Execute() {
	switch s.Action {
	case HTTP:
		HTTPExecute(s)
	default:
		HTTPExecute(s)
	}
}
