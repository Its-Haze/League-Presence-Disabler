package config

import (
	"encoding/json"
	"log"
	"os"
)

func ModifyJSON(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading JSON file from %s: %v\n", filePath, err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("Error unmarshalling JSON file: %v\n", err)
		return
	}

	modified := false
	plugins, ok := data["plugins"].([]interface{})
	if !ok {
		log.Printf("Error asserting JSON data for plugins at %s\n", filePath)
		return
	}

	for i, pluginInterface := range plugins {
		plugin, ok := pluginInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if plugin["name"] == "rcp-be-lol-discord-rp" {
			plugins = append(plugins[:i], plugins[i+1:]...)
			modified = true
		}
	}

	if modified {
		data["plugins"] = plugins
		newData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			log.Printf("Error marshalling JSON: %v\n", err)
			return
		}
		err = os.WriteFile(filePath, newData, 0644)
		if err != nil {
			log.Printf("Error writing modified JSON to %s: %v\n", filePath, err)
			return
		}
		log.Printf("Successfully disabled League Native Rich Presence in %s\n", filePath)
	}
}
