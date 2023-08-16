package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
)

type Case []Step

type Step struct {
	Name   string `json:"name" yaml:"name"`
	Action Action `json:"action" yaml:"action"`

	// env config
	Env       string   `json:"env,omitempty" yaml:"env,omitempty"`
	Tag       []string `json:"tag,omitempty" yaml:"tag,omitempty"`
	Envs      []string `json:"envs,omitempty" yaml:"envs,omitempty"`
	Countries []string `json:"countries,omitempty" yaml:"countries,omitempty"`
	Country   string   `json:"country,omitempty" yaml:"country,omitempty"`

	// request config
	Repeat    interface{} `json:"repeat,omitempty" yaml:"repeat,omitempty"`
	Sleep     int         `json:"sleep,omitempty" yaml:"sleep,omitempty"`
	Variable  *Variable   `json:"variable,omitempty" yaml:"variable,omitempty"`
	Variables interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	Request   interface{} `json:"request,omitempty" yaml:"request,omitempty"`
	Export    Export      `json:"export,omitempty" yaml:"export,omitempty"`
	Check     Check       `json:"check,omitempty" yaml:"check,omitempty"`

	// setup and teardown
	Skip     interface{}   `json:"skip,omitempty" yaml:"skip,omitempty"`
	Reason   string        `json:"reason,omitempty" yaml:"reason,omitempty"`
	Setup    []interface{} `json:"setup,omitempty" yaml:"setup,omitempty"`
	Teardown []interface{} `json:"teardown,omitempty" yaml:"teardown,omitempty"`

	// during loading
	Config           *Step         `json:"config,omitempty" yaml:"config,omitempty"`
	PathsInVariable  []interface{} `json:"paths_in_variable,omitempty" yaml:"paths_in_variable,omitempty"`
	PathsInVariables []interface{} `json:"paths_in_variables,omitempty" yaml:"paths_in_variables,omitempty"`
	PathsInRequest   []interface{} `json:"paths_in_request,omitempty" yaml:"paths_in_request,omitempty"`
	PathsInCheck     []interface{} `json:"paths_in_check,omitempty" yaml:"paths_in_check,omitempty"`
	PathsInExport    []interface{} `json:"paths_in_export,omitempty" yaml:"paths_in_export,omitempty"`
	PathsInSetup     []interface{} `json:"paths_in_setup,omitempty" yaml:"paths_in_setup,omitempty"`
	PathsInTeardown  []interface{} `json:"paths_in_teardown,omitempty" yaml:"paths_in_teardown,omitempty"`
	PathsUnknown     []interface{} `json:"paths_unknown,omitempty" yaml:"paths_unknown,omitempty"`

	// during running
	Model        *TestModel             `json:"model,omitempty" yaml:"model,omitempty"`
	Start        string                 `json:"start,omitempty" yaml:"start,omitempty"`
	End          string                 `json:"end,omitempty" yaml:"end,omitempty"`
	Response     interface{}            `json:"response,omitempty" yaml:"response,omitempty"`
	FullResponse map[string]interface{} `json:"full_response,omitempty" yaml:"full_response,omitempty"`
	FullRR       string                 `json:"full_rr,omitempty" yaml:"full_rr,omitempty"`

	// in locust
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}

type Action string

type Variable interface{}

type Export []interface{}

type Check []interface{}

type TestModel interface{}

func main() {
	var configFile string

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "App description",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("args:", args)
			// Read and parse YAML file
			yamlData, err := os.ReadFile(configFile)
			if err != nil {
				log.Fatal(err)
			}

			var tase Case
			err = yaml.Unmarshal(yamlData, &tase)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("case:", tase)
		},
	}

	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Config file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
