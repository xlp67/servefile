package internal

import (
	"fileserve/configs"
	"log"
	"encoding/json"
	"os"
	"strings"
)

type Block interface {
	Insert(name string)
	Remove(name string)
}

func Insert(name string) {
	if name == "" {
		log.Println("name required")
} else {
	var config *configs.ConfigsApp
	var check string
	read, err := os.ReadFile("config.json")
	if err != nil {panic(err)}	
	err = json.Unmarshal(read, &config)
	if err != nil {panic(err)}
	for _, block := range config.BlockDirectories {
		if strings.Contains(block, name) {
			check = block
			log.Printf("the %s already exist\n", name)
			break
		}
	}
	if check == "" {
	config.BlockDirectories = append(config.BlockDirectories, name)
	dataJson, err := json.Marshal(&config)
	if err != nil {panic(err)}
	err = os.WriteFile("config.json", dataJson, 0700)
	if err != nil {panic(err)}
	log.Printf("the %s registred\n", name)
	}
	}
}


func Remove(name string) {
	var config *configs.Configs
	read, err := os.ReadFile("config.json")
	if err != nil {panic(err)}	
	err = json.Unmarshal(read, &config)
	if err != nil {panic(err)}
	var check string
	for i, block := range config.BlockDirectories {
		if name == block {
			check = block
			left := config.BlockDirectories[:i]
			right := config.BlockDirectories[i+1:]
			var result []string
			result = append(result, left...)
			result = append(result, right...)
			config.BlockDirectories = config.BlockDirectories[:0]
			config.BlockDirectories = append(config.BlockDirectories, result...)
			dataJson, err := json.Marshal(&config)
			if err != nil {panic(err)}
			err = os.WriteFile("config.json", dataJson, 0700)
			if err != nil {panic(err)}
			log.Printf("%s removed\n", name)
			break
		}
	}
	if check == "" {
		log.Printf("%s not found\n", name)
	}
}