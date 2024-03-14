package pkg

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func GetAbsPath(path string) string {
	workDir, _ := os.Getwd()
	absPath, _ := filepath.Abs(filepath.Join(workDir, filepath.ToSlash(path)))
	return absPath
}

func PathToCases(path string) ([]Case, error) {
	fileList, err := PathToList(GetAbsPath(path))
	if err != nil {
		return nil, err
	}

	var caseList []Case
	for _, filePath := range fileList {
		tase, err := FileToCase(filePath)
		if err != nil {
			return nil, err
		}
		caseList = append(caseList, tase)
	}
	return caseList, nil
}

func PathToList(path string) ([]string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// 如果路径是一个文件，则直接返回该文件路径
	if !fileInfo.IsDir() {
		return []string{path}, nil
	}

	return DirToFiles(path), nil
}

func DirToFiles(dirPath string) []string {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil
	}

	var fileList []string
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			fileList = append(fileList, DirToFiles(filePath)...)
		} else {
			fileList = append(fileList, filePath)
		}
	}

	return fileList
}

func FileToCase(filePath string) (Case, error) {
	Logger.Info(filePath)
	// Read and parse YAML file
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		return Case{}, err
	}
	Logger.Info(string(yamlData))

	// Unmarshal steps
	var steps []Step
	err = yaml.Unmarshal(yamlData, &steps)
	if err != nil {
		return Case{}, err
	}

	return Case{Steps: steps, Cache: map[string]any{}}, nil
}
