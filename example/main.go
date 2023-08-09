package example

import (
	"fmt"
	"log"

	"github.com/0mlml/cfgparser"
)

var defaultConfig = &cfgparser.Config{
	BoolOptions: map[string]bool{
		"debug":        false,
		"auto_restart": true,
	},
	StringOptions: map[string]string{
		"app_name":             "MyApp",
		"db_connection_string": "host=localhost;user=user;password=pass;db=mydb",
	},
	IntOptions: map[string]int{
		"max_retries": 3,
		"port":        8080,
	},
	FloatOptions: map[string]float64{
		"version": 1.0,
	},
}

func main() {
	// Set the default configuration
	cfgparser.SetDefaultConfig(defaultConfig)

	// Parse the configuration file
	configPath := "example.cfg"
	config, err := cfgparser.ParseConfig(configPath)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Print the current configuration
	fmt.Println("Current Configuration:")
	printConfig(config)

	// Modify a value
	config.BoolOptions["debug"] = true
	config.StringOptions["app_name"] = "UpdatedApp"

	// Write the modified configuration to a new file
	newConfigPath := "updated_example.cfg"
	err = cfgparser.WriteConfig(newConfigPath, config)
	if err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}

	fmt.Printf("Updated configuration written to %s\n", newConfigPath)
}

func printConfig(config *cfgparser.Config) {
	fmt.Println("[bool]")
	for key, value := range config.BoolOptions {
		fmt.Printf("%s=%v\n", key, value)
	}

	fmt.Println("[string]")
	for key, value := range config.StringOptions {
		fmt.Printf("%s=%s\n", key, value)
	}

	fmt.Println("[int]")
	for key, value := range config.IntOptions {
		fmt.Printf("%s=%d\n", key, value)
	}

	fmt.Println("[float]")
	for key, value := range config.FloatOptions {
		fmt.Printf("%s=%f\n", key, value)
	}
}
