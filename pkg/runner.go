package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func RunCase(casePath string) {
	_, err := os.Stat(casePath)
	if err != nil {
		casePath = "examples/post_json.yaml"
	}

	// Read and parse YAML file
	yamlData, err := os.ReadFile(casePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yamlData))

	var tase []Step
	err = yaml.Unmarshal(yamlData, &tase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("case:", tase[0].Request)

	for _, step := range tase {
		var payload, _ = json.Marshal(step.Request.JSON)
		var resp, err = http.Post(step.Request.Host+step.Request.Path, "application/json", bytes.NewReader(payload))
		fmt.Println(resp, err)
	}

}
