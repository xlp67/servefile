package configs

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/viper"
)
type ConfigsApp struct {
	Directory string 
	BlockDirectories []string 
}
type ConfigHandle struct {
	PathDirectory string `mapstructure:"HANDLER_DIR"`
	PathAddBlock string `mapstructure:"HANDLER_INSERT_BLOCK"`
	PathRemoveBlock string `mapstructure:"HANDLER_REMOVE_BLOCK"`
	AppPort string `mapstructure:"APP_PORT"`
}
type Configs struct {
	ConfigsApp
	ConfigHandle
}
func Config(path string) (*Configs, error) {
	var config Configs
	viper.SetConfigName("app")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {log.Fatal(err)}
	err = viper.Unmarshal(&config.ConfigHandle)
	if err != nil {log.Fatal(err)}
	read, err := os.ReadFile("config.json")
	if err != nil {panic(err)}	
	err = json.Unmarshal(read, &config.ConfigsApp)
	if err != nil {panic(err)}	
	return &config, nil
}