package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fyllekanin/gnsagent/internal/schema"
	"github.com/fyllekanin/gnsagent/internal/util"

	"github.com/jasonlvhit/gocron"
)

func runTask(config schema.ConfigSchema) {
	result, err := util.GetIpFromEndPoints(config.EndPoints)
	if err != nil {
		fmt.Println("Error: could not get any public IP address")
		return
	}

	for _, domain := range config.Domains {
		util.UpdateDnsService(result, domain)
	}
}

func main() {
	configFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error: config.json file could not be located")
		return
	}

	var config schema.ConfigSchema
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error: config.json is not valid JSON format")
		return
	}

	runTask(config)
	gocron.Every(5).Minutes().Do(runTask, config)
}
