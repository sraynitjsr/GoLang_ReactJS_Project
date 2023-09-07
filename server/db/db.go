package db

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/sraynitjsr/server/model"
)

func GetAllData() ([]model.Configuration, error) {
	dataPath := "DATA"
	var allData []model.Configuration

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", path, err)
				return err
			}
			var config model.Configuration
			if err := json.Unmarshal(fileContent, &config); err != nil {
				log.Printf("Error parsing JSON file %s: %v\n", path, err)
				return err
			}
			allData = append(allData, config)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return allData, nil
}

func GetNEByType(neType string) ([]model.Configuration, error) {
	dataPath := "DATA"
	var filteredData []model.Configuration

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", path, err)
				return err
			}
			var config model.Configuration
			if err := json.Unmarshal(fileContent, &config); err != nil {
				log.Printf("Error parsing JSON file %s: %v\n", path, err)
				return err
			}
			if config.NEType == neType {
				filteredData = append(filteredData, config)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filteredData, nil
}

func GetByBuildVersion(buildVersion string) ([]model.Configuration, error) {
	dataPath := "DATA"
	var filteredData []model.Configuration

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", path, err)
				return err
			}
			var config model.Configuration
			if err := json.Unmarshal(fileContent, &config); err != nil {
				log.Printf("Error parsing JSON file %s: %v\n", path, err)
				return err
			}
			if config.BuildVersion == buildVersion {
				filteredData = append(filteredData, config)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filteredData, nil
}

func GetBySupportedConfig(supportedConfig string) ([]model.Configuration, error) {
	dataPath := "DATA"
	var filteredData []model.Configuration

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", path, err)
				return err
			}
			var config model.Configuration
			if err := json.Unmarshal(fileContent, &config); err != nil {
				log.Printf("Error parsing JSON file %s: %v\n", path, err)
				return err
			}
			for key := range config.SupportedConfigurations {
				if key == supportedConfig {
					filteredData = append(filteredData, config)
					break
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filteredData, nil
}

func GetByNETypeAndBuildVersion(neType, buildVersion string) ([]model.Configuration, error) {
	dataPath := "DATA"
	var filteredData []model.Configuration

	err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			fileContent, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", path, err)
				return err
			}
			var config model.Configuration
			if err := json.Unmarshal(fileContent, &config); err != nil {
				log.Printf("Error parsing JSON file %s: %v\n", path, err)
				return err
			}
			if config.NEType == neType && config.BuildVersion == buildVersion {
				filteredData = append(filteredData, config)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filteredData, nil
}
