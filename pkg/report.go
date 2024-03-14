package pkg

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

func GenerateReportFromJson(tmplPath string, jsonPath string) {
	_, err := os.Stat(jsonPath)
	if err != nil {
		jsonPath = "template/data.json"
	}

	// 从文件读取 JSON 数据源
	jsonFile, err := os.Open(jsonPath)
	defer jsonFile.Close()
	if err != nil {
		log.Fatal("Failed to open JSON file:", err)
	}

	data := map[string]any{}
	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		log.Fatal("Failed to decode JSON:", err)
	}

	GenerateReport(tmplPath, data)
}

func GenerateReport(tmplPath string, data map[string]any) {
	_, err := os.Stat(tmplPath)
	if err != nil {
		tmplPath = "template/report.html"
	}

	// 从文件读取 HTML 模板
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}

	// 创建新文件用于写入HTML报告
	reportFile, err := os.Create("logs/report.html")
	if err != nil {
		log.Fatal("Failed to create report file:", err)
	}
	defer reportFile.Close()

	_ = tmpl.Execute(reportFile, data)
}
